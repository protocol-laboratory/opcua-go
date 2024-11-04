package enc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"

	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
	"github.com/protocol-laboratory/opcua-go/opcua/util"
)

type Decoder interface {
	ReadMsg() (*uamsg.Message, error)
}

// FastDecoder Performance-sensitive structures need to implement fast decoder
type FastDecoder interface {
	FastDecode(r io.Reader, v reflect.Value) error
}

type bufferedDecoder struct {
	r             *superReader
	maxBufferSize int64
}

func NewDefaultDecoder(r io.Reader, maxBufferSize int64) Decoder {
	b := &bufferedDecoder{
		r: &superReader{
			r:    bufio.NewReader(r),
			buff: bytes.NewBuffer(nil),
		},
		maxBufferSize: maxBufferSize,
	}
	return b
}

func (d *bufferedDecoder) ReadMsg() (*uamsg.Message, error) {
	msg := &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{},
	}

	var (
		messageHeaderLen  int
		securityHeaderLen int
		sequenceHeaderLen int
		messageSize       uint32
		totalBodySize     int64
	)

	for {
		d.r.modifyState(false)
		p, err := d.r.readN(3)
		if err != nil {
			return nil, err
		}
		msg.MessageType = uamsg.MessageTypeEnum(p)

		chunkType, err := d.r.readByte()
		if err != nil {
			return nil, err
		}
		msg.ChunkType = uamsg.ChunkTypeEnum(chunkType)

		b, err := d.r.readN(4)
		if err != nil {
			return nil, err
		}
		messageSize = binary.LittleEndian.Uint32(b)

		switch msg.MessageType {
		case uamsg.HelloMessageType, uamsg.AcknowledgeMessageType:
			messageHeaderLen = 3 + 1 + 4
			securityHeaderLen = 0
			sequenceHeaderLen = 0
		case uamsg.OpenSecureChannelMessageType, uamsg.MsgMessageType, uamsg.CloseSecureChannelMessageType:
			b, err = d.r.readN(4)
			if err != nil {
				return nil, err
			}
			msg.SecureChannelId = util.GetPtr(binary.LittleEndian.Uint32(b))
			messageHeaderLen = 3 + 1 + 4 + 4

			if msg.MessageType == uamsg.OpenSecureChannelMessageType {
				securityHeader := &uamsg.AsymmetricSecurityHeader{}
				err = d.readTo(reflect.ValueOf(securityHeader).Elem())
				if err != nil {
					return nil, err
				}
				msg.SecurityHeader = securityHeader
				securityHeaderLen = 4 + len(securityHeader.SecurityPolicyUri) + 4 + len(securityHeader.SenderCertificate) + 4 + len(securityHeader.ReceiverCertificateThumbprint)
			} else {
				securityHeader := &uamsg.SymmetricSecurityHeader{}
				err = d.readTo(reflect.ValueOf(securityHeader).Elem())
				if err != nil {
					return nil, err
				}
				msg.SecurityHeader = securityHeader
				securityHeaderLen = 4
			}

			sequenceHeader := &uamsg.SequenceHeader{}
			err = d.readTo(reflect.ValueOf(sequenceHeader).Elem())
			if err != nil {
				return nil, err
			}
			msg.SequenceHeader = sequenceHeader
			sequenceHeaderLen = 8
		}

		switch msg.ChunkType {
		case uamsg.IntermediateChunkType, uamsg.FinalChunkType:
			bodySize := messageSize - uint32(messageHeaderLen) - uint32(securityHeaderLen) - uint32(sequenceHeaderLen)
			err := d.r.readN2Buffer(int32(bodySize))
			if err != nil {
				return nil, err
			}
			totalBodySize += int64(bodySize)
			if totalBodySize+int64(messageHeaderLen)+int64(securityHeaderLen)+int64(sequenceHeaderLen) > d.maxBufferSize {
				return nil, errors.New("too big message")
			}
			if msg.ChunkType == uamsg.IntermediateChunkType {
				continue
			}
			msg.MessageSize = uint32(totalBodySize + int64(messageHeaderLen) + int64(securityHeaderLen) + int64(sequenceHeaderLen))
		case uamsg.AbortChunkType:
			fallthrough
		default:
			return nil, errors.New("not support chunk type")
		}

		err = d.fillMessageBody(msg)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}
}

func (d *bufferedDecoder) fillMessageBody(msg *uamsg.Message) error {
	d.r.modifyState(true) // read bytes from buffer
	switch msg.MessageType {
	case uamsg.HelloMessageType:
		msg.MessageBody = &uamsg.HelloMessageExtras{}
		err := d.readTo(reflect.ValueOf(msg.MessageBody).Elem())
		if err != nil {
			return err
		}
	case uamsg.AcknowledgeMessageType:
		msg.MessageBody = &uamsg.AcknowledgeMessageExtras{}
		err := d.readTo(reflect.ValueOf(msg.MessageBody).Elem())
		if err != nil {
			return err
		}
	case uamsg.OpenSecureChannelMessageType, uamsg.MsgMessageType, uamsg.CloseSecureChannelMessageType:
		messageBody := &uamsg.GenericBody{}

		messageBody.TypeId = &uamsg.ExpandedNodeId{}
		err := d.readTo(reflect.ValueOf(messageBody.TypeId).Elem())
		if err != nil {
			return err
		}

		// TODO process more Identifier type
		switch messageBody.TypeId.Identifier {
		case uamsg.ObjectOpenSecureChannelRequest_Encoding_DefaultBinary.Identifier:
			service := &uamsg.OpenSecureChannelRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectOpenSecureChannelResponse_Encoding_DefaultBinary.Identifier:
			service := &uamsg.OpenSecureChannelResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectCreateSessionRequest_Encoding_DefaultBinary.Identifier:
			service := &uamsg.CreateSessionRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectCreateSessionResponse_Encoding_DefaultBinary.Identifier:
			service := &uamsg.CreateSessionResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectActivateSessionRequest_Encoding_DefaultBinary.Identifier:
			service := &uamsg.ActivateSessionRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectActivateSessionResponse_Encoding_DefaultBinary.Identifier:
			service := &uamsg.ActivateSessionResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectGetEndpointsRequest_Encoding_DefaultBinary.Identifier:
			service := &uamsg.GetEndpointsRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectGetEndpointsResponse_Encoding_DefaultBinary.Identifier:
			service := &uamsg.GetEndpointsResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectReadRequest_Encoding_DefaultBinary.Identifier:
			service := &uamsg.ReadRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectReadResponse_Encoding_DefaultBinary.Identifier:
			service := &uamsg.ReadResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectCloseSecureChannelRequest_Encoding_DefaultBinary.Identifier:
			service := &uamsg.CloseSecureChannelRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectCloseSecureChannelResponse_Encoding_DefaultBinary.Identifier:
			service := &uamsg.CloseSecureChannelResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectCloseSessionRequest_Encoding_DefaultBinary.Identifier:
			service := &uamsg.CloseSessionRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case uamsg.ObjectCloseSessionResponse_Encoding_DefaultBinary.Identifier:
			service := &uamsg.CloseSessionResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		default:
			return fmt.Errorf("unsupported service type: %v", messageBody.TypeId.Identifier)
		}
		msg.MessageBody = messageBody
	default:
	}
	return nil
}

// read data from buffer into struct
func (d *bufferedDecoder) readTo(value reflect.Value) error {
	valueType := value.Type()
	switch value.Kind() {
	// set new value by value kind
	case reflect.Bool:
		dataByte, err := d.r.readByte()
		if err != nil {
			return err
		}
		value.SetBool(dataByte == 1)
	case reflect.Uint8:
		dataByte, err := d.r.readByte()
		if err != nil {
			return err
		}
		value.SetUint(uint64(dataByte))
	case reflect.Uint16:
		b, err := d.r.readN(2)
		if err != nil {
			return err
		}
		value.SetUint(uint64(binary.LittleEndian.Uint16(b)))
	case reflect.Uint32:
		b, err := d.r.readN(4)
		if err != nil {
			return err
		}
		value.SetUint(uint64(binary.LittleEndian.Uint32(b)))
	case reflect.Uint64:
		b, err := d.r.readN(8)
		if err != nil {
			return err
		}
		value.SetUint(binary.LittleEndian.Uint64(b))
	case reflect.Int32:
		b, err := d.r.readN(4)
		if err != nil {
			return err
		}
		value.SetInt(int64(int32(binary.LittleEndian.Uint32(b))))
	case reflect.Int64:
		b, err := d.r.readN(8)
		if err != nil {
			return err
		}
		value.SetInt(int64(binary.LittleEndian.Uint64(b)))
	case reflect.Float32:
		b, err := d.r.readN(4)
		if err != nil {
			return err
		}
		value.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(b))))
	case reflect.Float64:
		b, err := d.r.readN(8)
		if err != nil {
			return err
		}
		value.SetFloat(math.Float64frombits(binary.LittleEndian.Uint64(b)))
	case reflect.String:
		b, err := d.r.readN(4)
		if err != nil {
			return err
		}
		length := int32(binary.LittleEndian.Uint32(b))
		if length == -1 {
			return nil
		}
		b, err = d.r.readN(length)
		if err != nil {
			return err
		}
		value.SetString(string(b))
	case reflect.Slice:
		err, done := d.readToSlice(value, valueType)
		if done {
			return err
		}
	case reflect.Ptr, reflect.Interface:
		if value.IsNil() {
			return errors.New("ptr or interface{} variant can not be nil")
		}
		err := d.readTo(value.Elem())
		if err != nil {
			return err
		}
	case reflect.Struct:
		err := readToStruct(value, valueType, d)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported type: %v", valueType.Kind())
	}

	return nil
}

func (d *bufferedDecoder) readToSlice(value reflect.Value, valueType reflect.Type) (error, bool) {
	b, err := d.r.readN(4)
	if err != nil {
		return err, true
	}
	sliceLen := int32(binary.LittleEndian.Uint32(b))

	if sliceLen == -1 {
		return nil, true
	}

	sliceValue := reflect.MakeSlice(valueType, int(sliceLen), int(sliceLen))
	for i := 0; i < sliceValue.Len(); i++ {
		if valueType.Elem().Kind() == reflect.Ptr {
			structPtr := reflect.New(valueType.Elem().Elem())
			err = d.readTo(structPtr.Elem())
			if err != nil {
				return err, true
			}
			sliceValue.Index(i).Set(structPtr)
		} else {
			// simple slice, like []uint32
			err = d.readTo(sliceValue.Index(i))
			if err != nil {
				return err, true
			}
		}
	}
	value.Set(sliceValue)
	return nil, false
}

func readToStruct(value reflect.Value, valueType reflect.Type, d *bufferedDecoder) error {
	decoder, ok := SpecialStructDecoderMap[value.Type().Name()]
	if !ok {
		// recursively construct all members
		for i := 0; i < value.NumField(); i++ {
			if value.Field(i).Kind() == reflect.Ptr && value.Field(i).IsNil() {
				structPtr := reflect.New(valueType.Field(i).Type.Elem())
				err := d.readTo(structPtr.Elem())
				if err != nil {
					return err
				}
				value.Field(i).Set(structPtr)
			} else {
				err := d.readTo(value.Field(i))
				if err != nil {
					return err
				}
			}
		}
	} else {
		err := decoder(d.r, value)
		if err != nil {
			return err
		}
	}
	return nil
}

type superReader struct {
	state bool // read bytes from buff while true
	r     *bufio.Reader
	buff  *bytes.Buffer
}

func (s *superReader) readN(n int32) ([]byte, error) {
	if s.state {
		return s.readNFromBuffer(n)
	}
	if n < 0 {
		return nil, errors.New("byte`num can`t be less than 0")
	}
	// todo optimize byte slice read
	p := make([]byte, n)
	readNum, err := s.r.Read(p)
	if err != nil {
		return nil, err
	}
	if readNum != int(n) {
		return nil, errors.New("no enough bytes")
	}
	return p, nil
}

func (s *superReader) readByte() (byte, error) {
	if s.state {
		return s.readByteFromBuffer()
	}
	return s.r.ReadByte()
}

func (s *superReader) readN2Buffer(n int32) error {
	s.modifyState(false)
	byteData, err := s.readN(n)
	if err != nil {
		return err
	}
	num, err := s.buff.Write(byteData)
	if err != nil || num != int(n) {
		return err
	}
	return nil
}

func (s *superReader) readByteFromBuffer() (byte, error) {
	return s.buff.ReadByte()
}

func (s *superReader) readNFromBuffer(n int32) ([]byte, error) {
	if n < 0 {
		return nil, errors.New("byte`num can`t be less than 0")
	}
	// todo optimize byte slice read
	p := make([]byte, n)
	readNum, err := s.buff.Read(p)
	if err != nil {
		return nil, err
	}
	if readNum != int(n) {
		return nil, errors.New("no enough bytes")
	}
	return p, nil
}

func (s *superReader) modifyState(state bool) {
	s.state = state
}

package enc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"math"

	"errors"
	"io"
	"opcua-go/opcua/uamsg"
	"opcua-go/opcua/util"
	"reflect"
)

type Decoder interface {
	ReadMsg() (*uamsg.Message, error)
}

// 当前兜底逻辑用反射是按照协议实现的decoder
// 后续可以根据测试数据，在性能瓶颈处的结构体上做加速
// 需要结构体主动实现FastDecoder，按照自己的结构，快速完成解码，避免反射逻辑
type FastDecoder interface{
	FastDecode(r io.Reader, v reflect.Value) error
}

type bufferdDecoder struct {
	r             *superReader
	maxBufferSize int64
}

func NewDefaultDecoder(r io.Reader, maxBufferSize int64) Decoder {
	b := &bufferdDecoder{
		r: &superReader{
			r:    bufio.NewReader(r),
			buff: bytes.NewBuffer(nil),
		},
		maxBufferSize: maxBufferSize,
	}
	return b
}

func (d *bufferdDecoder) ReadMsg() (*uamsg.Message, error) {
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
		case uamsg.OpenSecureChannelMessageType, uamsg.MsgMessageType:
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
		case uamsg.IntermediateChunckType, uamsg.FinalChunckType:
			bodySize := messageSize - uint32(messageHeaderLen) - uint32(securityHeaderLen) - uint32(sequenceHeaderLen)
			err := d.r.readN2Buffer(int32(bodySize))
			if err != nil {
				return nil, err
			}
			totalBodySize += int64(bodySize)
			if totalBodySize+int64(messageHeaderLen)+int64(securityHeaderLen)+int64(sequenceHeaderLen) > d.maxBufferSize {
				return nil, errors.New("too big message")
			}
			if msg.ChunkType == uamsg.IntermediateChunckType {
				continue
			}
			msg.MessageSize = uint32(totalBodySize + int64(messageHeaderLen) + int64(securityHeaderLen) + int64(sequenceHeaderLen))
		case uamsg.AbortChunckType:
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

func (d *bufferdDecoder) fillMessageBody(msg *uamsg.Message) error {
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
		serviceType, ok := messageBody.TypeId.Identifier.(uint16)
		if !ok {
			return errors.New("know type service")
		}

		switch serviceType {
		case 446:
			// open secure channel request
			service := &uamsg.OpenSecureChannelServiceRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case 449:
			// open secure channel response
			service := &uamsg.OpenSecureChannelServiceResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case 461:
			// create session request
			service := &uamsg.CreateSessionRequest{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		case 464:
			// create session response
			service := &uamsg.CreateSessionResponse{}
			err = d.readTo(reflect.ValueOf(service).Elem())
			if err != nil {
				return err
			}
			messageBody.Service = service
		default:
			return errors.New("know type service")
		}
		msg.MessageBody = messageBody
	default:
	}
	return nil
}

// 把字节流读取到指定结构体中，这里的value只能是struct类型的指针
func (d *bufferdDecoder) readTo(value reflect.Value) error {
	vtyp := value.Type()
	switch value.Kind() {
	// 根据字段类型设置新的值
	case reflect.Uint8:
		dataByte, err := d.r.readByte()
		if err != nil {
			return err
		}
		value.SetUint(uint64(dataByte))
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
		value.SetUint(uint64(binary.LittleEndian.Uint64(b)))
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
		value.SetFloat(float64(math.Float64frombits(binary.LittleEndian.Uint64(b))))
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
		b, err := d.r.readN(4)
		if err != nil {
			return err
		}
		sliceLen := int32(binary.LittleEndian.Uint32(b))

		if sliceLen == -1 {
			return nil
		}

		sliceValue := reflect.MakeSlice(vtyp, int(sliceLen), int(sliceLen))
		for i := 0; i < sliceValue.Len(); i++ {
			if vtyp.Elem().Kind() == reflect.Ptr {
				// []*MyStruct
				// vtyp = slice, .Elem() = *MyStruct, .Elem() = MyStruct,
				structPtr := reflect.New(vtyp.Elem().Elem())
				err = d.readTo(structPtr.Elem())
				if err != nil {
					return err
				}
				sliceValue.Index(i).Set(structPtr)
			} else {
				// []uint32
				err = d.readTo(sliceValue.Index(i))
				if err != nil {
					return err
				}
			}
		}
		value.Set(sliceValue)
	case reflect.Ptr, reflect.Interface:
		if value.IsNil() {
			return errors.New("ptr or interface{} variant can not be nil")
		}
		err := d.readTo(value.Elem())
		if err != nil {
			return err
		}
	case reflect.Struct:
		decoder, ok := SpecialStructDecoderMap[value.Type().Name()]
		if !ok {
			// 递归构造所有成员
			for i := 0; i < value.NumField(); i++ {
				if value.Field(i).Kind() == reflect.Ptr && value.Field(i).IsNil() {
					structPtr := reflect.New(vtyp.Field(i).Type.Elem())
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

	default:
		return errors.New("unsupport type")
	}

	return nil
}

type superReader struct {
	state bool
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
	// todo 优化byte切片的获取
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
	// todo 优化byte切片的获取
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

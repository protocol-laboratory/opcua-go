package enc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
	"strings"

	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
)

var ErrValueIsNil = errors.New("value is nil")

type Encoder interface {
	Encode(v *uamsg.Message, chunkSize int) ([][]byte, error)
	SetSequenceNumberGenerator(func() uint32)
}

// FastEncoder Performance-sensitive structures need to implement fast encoder
type FastEncoder interface {
	FastEncode() ([]byte, error)
}

type DefaultEncoder struct {
	sequenceNumberGenerator func() uint32
}

func NewDefaultEncoder() *DefaultEncoder {
	return &DefaultEncoder{}
}

func (e *DefaultEncoder) Encode(v *uamsg.Message, chunkSize int) ([][]byte, error) {
	chunks := make([][]byte, 0)

	var (
		messageHeaderLength  int
		securityHeaderLength int
		sequenceHeaderLength int

		securityHeaderBytes []byte
		sequenceHeaderBytes []byte
	)

	switch v.MessageType {
	case uamsg.HelloMessageType, uamsg.AcknowledgeMessageType:
		messageHeaderLength = 3 + 1 + 4
		securityHeaderLength = 0
		sequenceHeaderLength = 0
	case uamsg.OpenSecureChannelMessageType, uamsg.MsgMessageType, uamsg.CloseSecureChannelMessageType:
		var err error
		securityHeaderBytes, err = genericEncoder(v.SecurityHeader)
		if err != nil {
			return nil, err
		}
		messageHeaderLength = 3 + 1 + 4 + 4
		securityHeaderLength = len(securityHeaderBytes)
		sequenceHeaderLength = 8
	default:
		return nil, errors.New("not support message type")
	}

	dataBytes, err := genericEncoder(v.MessageBody)
	if err != nil {
		return nil, err
	}

	leftBodySize := len(dataBytes)
	headerLength := messageHeaderLength + securityHeaderLength + sequenceHeaderLength

	for leftBodySize > 0 {
		tempBuff := bytes.NewBuffer(nil)
		leftMsgSize := leftBodySize + headerLength
		if leftMsgSize > chunkSize {
			v.MessageHeader.ChunkType = uamsg.IntermediateChunkType
			v.MessageHeader.MessageSize = uint32(chunkSize)
		} else {
			v.MessageHeader.ChunkType = uamsg.FinalChunkType
			v.MessageHeader.MessageSize = uint32(leftMsgSize)
		}
		headerBytes, err := genericEncoder(v.MessageHeader)
		if err != nil {
			return nil, err
		}
		err = binary.Write(tempBuff, binary.LittleEndian, headerBytes)
		if err != nil {
			return nil, err
		}
		if v.MessageHeader.MessageType != uamsg.HelloMessageType && v.MessageHeader.MessageType != uamsg.AcknowledgeMessageType {
			err = binary.Write(tempBuff, binary.LittleEndian, securityHeaderBytes)
			if err != nil {
				return nil, err
			}

			if e.sequenceNumberGenerator != nil {
				v.SequenceHeader.SequenceNumber = e.sequenceNumberGenerator()
			}
			sequenceHeaderBytes, err = genericEncoder(v.SequenceHeader)
			if err != nil {
				return nil, err
			}
			err = binary.Write(tempBuff, binary.LittleEndian, sequenceHeaderBytes)
			if err != nil {
				return nil, err
			}
		}

		if v.MessageHeader.ChunkType == uamsg.IntermediateChunkType {
			writeBody := dataBytes[:chunkSize-headerLength]
			dataBytes = dataBytes[chunkSize-headerLength:]
			leftBodySize = len(dataBytes)
			err = binary.Write(tempBuff, binary.LittleEndian, writeBody)
			if err != nil {
				return nil, err
			}
		} else {
			leftBodySize = 0
			err = binary.Write(tempBuff, binary.LittleEndian, dataBytes)
			if err != nil {
				return nil, err
			}
		}
		chunks = append(chunks, tempBuff.Bytes())
	}
	return chunks, nil
}

func genericEncoder(v interface{}) ([]byte, error) {
	buff := bytes.NewBuffer(nil)

	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, ErrValueIsNil
		}
		value = reflect.Indirect(value)
	}

	switch value.Kind() {
	case reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		err := binary.Write(buff, binary.LittleEndian, value.Interface())
		if err != nil {
			return nil, err
		}
		return buff.Bytes(), nil
	case reflect.String:
		strBytes, err := StringEncoder(value.Interface())
		if err != nil {
			return nil, err
		}
		return strBytes, nil
	case reflect.Slice:
		if value.Type().Elem().String() == reflect.Uint8.String() {
			// byte slice fast path
			valueBytes, err := ByteStringEncoder(value.Interface())
			if err != nil {
				return nil, err
			}
			return valueBytes, nil
		}

		length := int32(value.Len())
		if length == 0 && value.IsNil() {
			length = -1
		}
		err := binary.Write(buff, binary.LittleEndian, length)
		if err != nil {
			return nil, err
		}

		for i := 0; i < value.Len(); i++ {
			elemBytes, err := genericEncoder(value.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			buff.Write(elemBytes)
		}

	case reflect.Struct:
		if value.Type().Implements(reflect.TypeOf((*FastEncoder)(nil)).Elem()) {
			// fast path for all struct which implements FastEncoder
			v, _ := value.Interface().(FastEncoder)
			dataBytes, err := v.FastEncode()
			if err != nil {
				return nil, err
			}
			buff.Write(dataBytes)
			return buff.Bytes(), nil
		}
		encoder, ok := SpecialStructEncoderMap[value.Type().Name()]
		if !ok {
			for i := 0; i < value.NumField(); i++ {
				elemBytes, err := genericEncoder(value.Field(i).Interface())
				if err != nil {
					if !errors.Is(err, ErrValueIsNil) {
						return nil, err
					}
					// enc: omitempty标记的值解不出来可以不加入编码，不算错误
					tagValue, ok := value.Type().Field(i).Tag.Lookup("enc")
					if !ok {
						return nil, err
					}
					if !strings.Contains(tagValue, "omitempty") {
						return nil, err
					}
					continue
				}
				buff.Write(elemBytes)
			}
		} else {
			valueBytes, err := encoder(value.Interface())
			if err != nil {
				return nil, err
			}
			return valueBytes, nil
		}
	default:
		err := binary.Write(buff, binary.LittleEndian, value.Interface())
		if err != nil {
			return nil, err
		}
	}

	return buff.Bytes(), nil
}

func (e *DefaultEncoder) SetSequenceNumberGenerator(f func() uint32) {
	e.sequenceNumberGenerator = f
}

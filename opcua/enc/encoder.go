/*
encoder共计三种：generic encoder、special encoder、basic encoder。
*/
package enc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"opcua-go/opcua/uamsg"
	"reflect"
	"strings"
)

var ErrValueIsNil = errors.New("value is nill")

type Encoder interface {
	Encode(v *uamsg.Message, chunksize int) ([][]byte, error)
	SetSequanceNumberGenerator(func() uint32)
}

type DefaultEndocer struct {
	sequanceNumberGenerator func() uint32
}

func (e *DefaultEndocer) Encode(v *uamsg.Message, chunksize int) ([][]byte, error) {
	chunks := make([][]byte, 0)

	var (
		messageHeaderLength  = 0
		securityHeaderLength = 0
		sequenceHeaderLength = 0

		securityHeaderBytes = make([]byte, 0)
		sequenceHeaderBytes = make([]byte, 0)
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

	leftBodysize := len(dataBytes)
	headerLength := messageHeaderLength + securityHeaderLength + sequenceHeaderLength

	for leftBodysize > 0 {
		tempbuff := bytes.NewBuffer(nil)
		leftMsgSize := leftBodysize + headerLength
		if leftMsgSize > chunksize {
			v.MessageHeader.ChunkType = uamsg.IntermediateChunckType
			v.MessageHeader.MessageSize = uint32(chunksize)
		} else {
			v.MessageHeader.ChunkType = uamsg.FinalChunckType
			v.MessageHeader.MessageSize = uint32(leftMsgSize)
		}
		headerBytes, err := genericEncoder(v.MessageHeader)
		if err != nil {
			return nil, err
		}
		err = binary.Write(tempbuff, binary.LittleEndian, headerBytes)
		if err != nil {
			return nil, err
		}
		if v.MessageHeader.MessageType != uamsg.HelloMessageType && v.MessageHeader.MessageType != uamsg.AcknowledgeMessageType {
			// 只有这两种消息会有这个消息头
			err = binary.Write(tempbuff, binary.LittleEndian, securityHeaderBytes)
			if err != nil {
				return nil, err
			}

			if e.sequanceNumberGenerator != nil {
				v.SequenceHeader.SequenceNumber = e.sequanceNumberGenerator()
			}
			sequenceHeaderBytes, err = genericEncoder(v.SequenceHeader)
			if err != nil {
				return nil, err
			}
			err = binary.Write(tempbuff, binary.LittleEndian, sequenceHeaderBytes)
			if err != nil {
				return nil, err
			}
		}

		if v.MessageHeader.ChunkType == uamsg.IntermediateChunckType {
			writeBody := dataBytes[:chunksize-headerLength]
			dataBytes = dataBytes[chunksize-headerLength:]
			leftBodysize = len(dataBytes)
			err = binary.Write(tempbuff, binary.LittleEndian, writeBody)
			if err != nil {
				return nil, err
			}
		} else {
			leftBodysize = 0
			err = binary.Write(tempbuff, binary.LittleEndian, dataBytes)
			if err != nil {
				return nil, err
			}
		}
		chunks = append(chunks, tempbuff.Bytes())
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
		binary.Write(buff, binary.LittleEndian, value.Interface())
		return buff.Bytes(), nil
	case reflect.String:
		strBytes, err := StringEncoder(value.Interface())
		if err != nil {
			return nil, err
		}
		return strBytes, nil
	case reflect.Slice:
		if value.Type().Elem().String() == reflect.Uint8.String() {
			// byte 数组加速分支
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
		binary.Write(buff, binary.LittleEndian, length)

		for i := 0; i < value.Len(); i++ {
			elemBytes, err := genericEncoder(value.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			buff.Write(elemBytes)
		}

	case reflect.Struct:
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
		binary.Write(buff, binary.LittleEndian, value.Interface())
	}

	return buff.Bytes(), nil
}

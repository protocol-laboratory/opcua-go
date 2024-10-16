package enc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
)

func StringEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("error type of string, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	length := int32(len(value))
	if length == 0 {
		length = -1
	}
	err := binary.Write(buff, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	buff.WriteString(value)
	return buff.Bytes(), nil
}

func ByteStringEncoder(v interface{}) ([]byte, error) {
	value, ok := v.([]byte)
	if !ok {
		return nil, fmt.Errorf("error type of bytestring, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	// ByteString's empty and null encodings are inconsistent
	length := int32(len(value))
	if length == 0 && value == nil {
		length = -1
	}
	err := binary.Write(buff, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	buff.Write(value)
	return buff.Bytes(), nil
}

func GuidEncoder(v interface{}) ([]byte, error) {
	var guid uamsg.Guid
	temp, ok := v.(*uamsg.Guid)
	if !ok {
		guid, ok = v.(uamsg.Guid)
		if !ok {
			return nil, fmt.Errorf("value has wrong type, should be guid")
		}
		temp = &guid
	}
	guid = *temp
	buff := bytes.NewBuffer(nil)
	err := binary.Write(buff, binary.LittleEndian, guid.Data1)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.LittleEndian, guid.Data2)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.LittleEndian, guid.Data3)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.LittleEndian, guid.Data4)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func QualifiedNameEncoder(v interface{}) ([]byte, error) {
	var qualifiedName uamsg.QualifiedName
	temp, ok := v.(*uamsg.QualifiedName)
	if !ok {
		qualifiedName, ok = v.(uamsg.QualifiedName)
		if !ok {
			return nil, fmt.Errorf("value has wrong type, should be qualified name")
		}
		temp = &qualifiedName
	}
	qualifiedName = *temp
	buff := bytes.NewBuffer(nil)
	err := binary.Write(buff, binary.LittleEndian, qualifiedName.NamespaceIndex)
	if err != nil {
		return nil, err
	}
	tempBytes, err := StringEncoder(qualifiedName.Name)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buff, binary.LittleEndian, tempBytes)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

type SpecialEncoder func(v interface{}) ([]byte, error)

// SpecialStructEncoderMap Such types belong to the built-in types of the OPC UA protocol, but require special adaptation of the coding language.
var SpecialStructEncoderMap = map[string]SpecialEncoder{
	"NodeId":          NodeIdEncoder,
	"ExpandedNodeId":  ExpandedNodeIdEncoder,
	"ExtensionObject": ExtensionObjectEncoder,
	"DiagnosticInfo":  DiagnosticInfoEncoder,
	"LocalizedText":   LocalizedTextEncoder,
	"DataValue":       DataValueEncoder,
	"Variant":         VariantEncoder,
}

func NodeIdEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.NodeId)
	if !ok {
		return nil, fmt.Errorf("error type of NodeId, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	err := binary.Write(buff, binary.LittleEndian, value.EncodingType)
	if err != nil {
		return nil, err
	}

	// node id use low four bytes
	switch value.EncodingType & 0x0f {
	case uamsg.TwoByte:
		identifier, ok := value.Identifier.(byte)
		if !ok {
			return nil, fmt.Errorf("two byte nodeId`s identifier has wrong type")
		}
		err := binary.Write(buff, binary.LittleEndian, identifier)
		if err != nil {
			return nil, err
		}
	case uamsg.FourByte:
		identifier, ok := value.Identifier.(uint16)
		if !ok {
			return nil, fmt.Errorf("four byte nodeId`s identifier has wrong type")
		}
		err := binary.Write(buff, binary.LittleEndian, byte(value.Namespace))
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, identifier)
		if err != nil {
			return nil, err
		}
	case uamsg.Numeric:
		identifier, ok := value.Identifier.(uint32)
		if !ok {
			return nil, fmt.Errorf("numeric nodeId has wrong type identifier")
		}
		err := binary.Write(buff, binary.LittleEndian, value.Namespace)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, identifier)
		if err != nil {
			return nil, err
		}
	case uamsg.String:
		identifier, ok := value.Identifier.(string)
		if !ok {
			return nil, fmt.Errorf("identifier has wrong type, should be string")
		}
		err := binary.Write(buff, binary.LittleEndian, value.Namespace)
		if err != nil {
			return nil, err
		}
		idBytes, err := StringEncoder(identifier)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, idBytes)
		if err != nil {
			return nil, err
		}
	case uamsg.GuidType:
		idBytes, err := GuidEncoder(value.Identifier)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, value.Namespace)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, idBytes)
		if err != nil {
			return nil, err
		}
	case uamsg.ByteString:
		identifier, ok := value.Identifier.([]byte)
		if !ok {
			return nil, fmt.Errorf("bytestring nodeId has wrong type identifier")
		}
		err := binary.Write(buff, binary.LittleEndian, value.Namespace)
		if err != nil {
			return nil, err
		}
		idBytes, err := ByteStringEncoder(identifier)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, idBytes)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("nodeId has unknown type identifier")
	}
	return buff.Bytes(), nil
}

func ExpandedNodeIdEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.ExpandedNodeId)
	if !ok {
		return nil, fmt.Errorf("error type of ExpandedNodeId, %v", reflect.TypeOf(v).Name())
	}

	nodeIdBytes, err := NodeIdEncoder(*value.NodeId)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(nil)
	buff.Write(nodeIdBytes)
	if (value.EncodingType & uamsg.NamespaceUriFlag) != 0 {
		nsUriBytes, err := StringEncoder(value.NamespaceUri)
		if err != nil {
			return nodeIdBytes, err
		}
		err = binary.Write(buff, binary.LittleEndian, nsUriBytes)
		if err != nil {
			return nil, err
		}
	}
	if (value.EncodingType & uamsg.ServerIndexFlag) != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.ServerIndex)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

func ExtensionObjectEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.ExtensionObject)
	if !ok {
		return nil, fmt.Errorf("error type of ExtensionObject, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	typeIdBytes, err := NodeIdEncoder(*value.TypeId)
	if err != nil {
		return nil, err
	}
	buff.Write(typeIdBytes)
	buff.WriteByte(value.Encoding)
	switch value.Encoding {
	case 0x00:
	case 0x01:
		fallthrough
	case 0x02:
		bodyBytes, err := StringEncoder(value.Body)
		if err != nil {
			return nil, err
		}
		bodyLength := int32(len(bodyBytes))
		err = binary.Write(buff, binary.LittleEndian, bodyLength)
		if err != nil {
			return nil, err
		}
		buff.Write(bodyBytes)
	}
	return buff.Bytes(), nil
}

func DiagnosticInfoEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.DiagnosticInfo)
	if !ok {
		return nil, fmt.Errorf("error type of DiagnosticInfo, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	buff.WriteByte(value.EncodingMask)

	// todo optimization, chain mode
	if value.EncodingMask&uamsg.SymbolicIdFlag != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.SymbolicId)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&uamsg.NamespaceFlag != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.NamespaceUri)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&uamsg.LocaleFlag != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.Locale)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&uamsg.LocalizedTextFlag != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.LocalizedText)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&uamsg.AdditionalInfoFlag != 0 {
		tempBytes, err := StringEncoder(value.AdditionalInfo)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, tempBytes)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&uamsg.InnerStatusCodeFlag != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.InnerStatusCode)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&uamsg.InnerDiagnosticInfoFlag != 0 {
		dataBytes, err := DiagnosticInfoEncoder(*value.InnerDiagnosticInfo)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, dataBytes)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

func LocalizedTextEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.LocalizedText)
	if !ok {
		return nil, fmt.Errorf("error type of LocalizedText, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	buff.WriteByte(value.EncodingMask)

	if value.EncodingMask&0x01 != 0 {
		localeBytes, err := StringEncoder(value.Locale)
		if err != nil {
			return nil, err
		}
		buff.Write(localeBytes)
	}
	if value.EncodingMask&0x02 != 0 {
		textBytes, err := StringEncoder(value.Text)
		if err != nil {
			return nil, err
		}
		buff.Write(textBytes)
	}

	return buff.Bytes(), nil
}

func DataValueEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.DataValue)
	if !ok {
		return nil, fmt.Errorf("error type of DataValue, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	buff.WriteByte(value.EncodingMask)

	// todo optimization, chain mode
	if value.EncodingMask&0x01 != 0 {
		tempBytes, err := VariantEncoder(*value.Value)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buff, binary.LittleEndian, tempBytes)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&0x02 != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.ResultStatusCode)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&0x04 != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.SourceTimestamp)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&0x08 != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.ServerTimestamp)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&0x10 != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.SourcePicoSeconds)
		if err != nil {
			return nil, err
		}
	}
	if value.EncodingMask&0x20 != 0 {
		err := binary.Write(buff, binary.LittleEndian, value.ServerPicoSeconds)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

func VariantEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.Variant)
	if !ok {
		return nil, fmt.Errorf("error type of DataValue, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	buff.WriteByte(value.EncodingMask)

	// type-encoding implemented according to the table specified by the OPC UA protocol
	encodingFn := func(encodingMask byte, v interface{}) ([]byte, error) {
		tempBuff := bytes.NewBuffer(nil)
		switch encodingMask & 0x3f {
		case 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0d, 0x13:
			err := binary.Write(tempBuff, binary.LittleEndian, v)
			if err != nil {
				return nil, err
			}
		case 0x0c:
			tempBytes, err := StringEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x0e:
			idBytes, err := GuidEncoder(v)
			if err != nil {
				return nil, err
			}
			err = binary.Write(buff, binary.LittleEndian, idBytes)
			if err != nil {
				return nil, err
			}
		case 0x0f:
			tempBytes, err := ByteStringEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x10:
			tempBytes, err := StringEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x11:
			tempBytes, err := NodeIdEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x12:
			tempBytes, err := ExpandedNodeIdEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x14:
			tempBytes, err := QualifiedNameEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x15:
			tempBytes, err := LocalizedTextEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x16:
			tempBytes, err := ExtensionObjectEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x17:
			tempBytes, err := DataValueEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x18:
			tempBytes, err := VariantEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x19:
			tempBytes, err := DiagnosticInfoEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		default:
			return nil, errors.New("not support variant mask type")
		}
		return tempBuff.Bytes(), nil
	}

	switch {
	case value.EncodingMask&0xff == 0:
		// null variant, nothing to do here
	case value.EncodingMask&0x40 != 0:
		// multidimensional matrix
		// when Value stores a multidimensional matrix
		// ArrayLength stores the number of all elements of the matrix.
		err := binary.Write(buff, binary.LittleEndian, value.ArrayLength)
		if err != nil {
			return nil, err
		}
		vv := reflect.ValueOf(value.Value)
		for i := 0; i < vv.Len(); i++ {
			for j := 0; j < vv.Index(i).Len(); j++ {
				tempBytes, err := encodingFn(value.EncodingMask, vv.Index(i).Index(j).Interface())
				if err != nil {
					return nil, err
				}
				buff.Write(tempBytes)
			}
		}

		// ArrayDimensionsLength is the first dimension of the multidimensional matrix
		err = binary.Write(buff, binary.LittleEndian, value.ArrayDimensionsLength)
		if err != nil {
			return nil, err
		}

		// ArrayDimensions record the second dimension of the multidimensional matrix
		length := int32(len(value.ArrayDimensions))
		err = binary.Write(buff, binary.LittleEndian, length)
		if err != nil {
			return nil, err
		}
		for i := 0; i < int(length); i++ {
			err := binary.Write(buff, binary.LittleEndian, value.ArrayDimensions[i])
			if err != nil {
				return nil, err
			}
		}
	case value.EncodingMask&0x80 != 0:
		// one dimensional array
		err := binary.Write(buff, binary.LittleEndian, value.ArrayLength)
		if err != nil {
			return nil, err
		}
		vv := reflect.ValueOf(value.Value)
		for i := 0; i < vv.Len(); i++ {
			tempBytes, err := encodingFn(value.EncodingMask, vv.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			buff.Write(tempBytes)
		}
	default:
		// basic type
		tempBytes, err := encodingFn(value.EncodingMask, value.Value)
		if err != nil {
			return nil, err
		}
		buff.Write(tempBytes)
	}

	return buff.Bytes(), nil
}

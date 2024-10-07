package enc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"opcua-go/opcua/uamsg"
	"reflect"
)

// basic encoder
func StringEncoder(v interface{}) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	value, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("error type of string, %v", reflect.TypeOf(v).Name())
	}
	length := int32(len(value))
	if length == 0 {
		length = -1
	}
	binary.Write(buff, binary.LittleEndian, length)
	buff.WriteString(value)
	return buff.Bytes(), nil
}

func ByteStringEncoder(v interface{}) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	value, ok := v.([]byte)
	if !ok {
		return nil, fmt.Errorf("error type of bytestring, %v", reflect.TypeOf(v).Name())
	}
	// bytestring的empty和null编码不一致

	length := int32(len(value))
	if length == 0 && value == nil {
		length = -1
	}
	binary.Write(buff, binary.LittleEndian, length)
	buff.Write(value)
	return buff.Bytes(), nil
}

// special encoder
type SpecialEncoder func(v interface{}) ([]byte, error)

var SpecialStructEncoderMap = map[string]SpecialEncoder{
	"NodeId":          NodeIdEncoder,
	"ExpandedNodeId":  ExpandedNodeIdEncoder,
	"ExtensionObject": ExtensionObjectEncoder,
	"DiagnosticInfo":  DiagnosticInfoEncoder,
	"LocalizedText":   LocalizedTextEncoder,
	"DataValue":       DataValueEncoder,
	"Variant":         VariantEncoder,
}

// NodeId 类型的数据编码比较抽象，需要特殊处理
func NodeIdEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.NodeId)
	if !ok {
		return nil, fmt.Errorf("error type of NodeId, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	binary.Write(buff, binary.LittleEndian, value.EncodingType)

	// node id 使用低4位
	switch value.EncodingType & 0x0f {
	case uamsg.TwoByte:
		identifier, ok := value.Identifier.(byte)
		if !ok {
			return nil, fmt.Errorf("two byte nodeId`s identifier has wrong type")
		}
		binary.Write(buff, binary.LittleEndian, identifier)
	case uamsg.FourByte:
		identifier, ok := value.Identifier.(uint16)
		if !ok {
			return nil, fmt.Errorf("four byte nodeId`s identifier has wrong type")
		}
		binary.Write(buff, binary.LittleEndian, byte(value.Namespace))
		binary.Write(buff, binary.LittleEndian, identifier)
	case uamsg.Numeric:
		identifier, ok := value.Identifier.(uint32)
		if !ok {
			return nil, fmt.Errorf("numeric nodeId has wrong type identifier")
		}
		binary.Write(buff, binary.LittleEndian, value.Namespace)
		binary.Write(buff, binary.LittleEndian, identifier)
	case uamsg.String:
		identifier, ok := value.Identifier.(string)
		if !ok {
			return nil, fmt.Errorf("string nodeId has wrong type identifier")
		}
		binary.Write(buff, binary.LittleEndian, value.Namespace)
		idBytes, err := StringEncoder(identifier)
		if err != nil {
			return nil, err
		}
		binary.Write(buff, binary.LittleEndian, idBytes)
	case uamsg.GuidType:
		var identifier uamsg.Guid
		temp, ok := value.Identifier.(*uamsg.Guid)
		if !ok {
			identifier, ok = value.Identifier.(uamsg.Guid)
			if !ok {
				return nil, fmt.Errorf("string nodeId has wrong type identifier")
			}
			temp = &identifier
		}
		identifier = *temp

		binary.Write(buff, binary.LittleEndian, value.Namespace)
		binary.Write(buff, binary.LittleEndian, identifier.Data1)
		binary.Write(buff, binary.LittleEndian, identifier.Data2)
		binary.Write(buff, binary.LittleEndian, identifier.Data3)
		binary.Write(buff, binary.LittleEndian, identifier.Data4)
	case uamsg.ByteString:
		identifier, ok := value.Identifier.([]byte)
		if !ok {
			return nil, fmt.Errorf("bytestring nodeId has wrong type identifier")
		}
		binary.Write(buff, binary.LittleEndian, value.Namespace)
		idBytes, err := ByteStringEncoder(identifier)
		if err != nil {
			return nil, err
		}
		binary.Write(buff, binary.LittleEndian, idBytes)
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

	nodeIdBytes, err := NodeIdEncoder(value.NodeId)
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
		binary.Write(buff, binary.LittleEndian, nsUriBytes)
	}
	if (value.EncodingType & uamsg.ServerIndexFlag) != 0 {
		binary.Write(buff, binary.LittleEndian, value.ServerIndex)
	}
	return buff.Bytes(), nil
}

func ExtensionObjectEncoder(v interface{}) ([]byte, error) {
	value, ok := v.(uamsg.ExtensionObject)
	if !ok {
		return nil, fmt.Errorf("error type of ExtensionObject, %v", reflect.TypeOf(v).Name())
	}
	buff := bytes.NewBuffer(nil)
	typeIdBytes, err := NodeIdEncoder(value.TypeId)
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
		binary.Write(buff, binary.LittleEndian, bodyLength)
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
	switch {
	case value.EncodingMask&uamsg.SymbolicIdFlag != 0:
		binary.Write(buff, binary.LittleEndian, value.SymbolicId)
		fallthrough
	case value.EncodingMask&uamsg.NamespaceFlag != 0:
		binary.Write(buff, binary.LittleEndian, value.NamespaceUri)
		fallthrough
	case value.EncodingMask&uamsg.LocaleFlag != 0:
		binary.Write(buff, binary.LittleEndian, value.Locale)
		fallthrough
	case value.EncodingMask&uamsg.LocalizedTextFlag != 0:
		binary.Write(buff, binary.LittleEndian, value.LocalizedText)
		fallthrough
	case value.EncodingMask&uamsg.AdditionalInfoFlag != 0:
		tempBytes, err := StringEncoder(value.AdditionalInfo)
		if err != nil {
			return nil, err
		}
		binary.Write(buff, binary.LittleEndian, tempBytes)
		fallthrough
	case value.EncodingMask&uamsg.InnerStatusCodeFlag != 0:
		binary.Write(buff, binary.LittleEndian, value.InnerStatusCode)
		fallthrough
	case value.EncodingMask&uamsg.InnerDiagnosticInfoFlag != 0:
		dataBytes, err := DiagnosticInfoEncoder(*value.InnerDiagnosticInfo)
		if err != nil {
			return nil, err
		}
		binary.Write(buff, binary.LittleEndian, dataBytes)
	default:
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
	switch {
	case value.EncodingMask&0x01 != 0:
		localeBytes, err := StringEncoder(value.Locale)
		if err != nil {
			return nil, err
		}
		buff.Write(localeBytes)
		fallthrough
	case value.EncodingMask&0x02 != 0:
		textBytes, err := StringEncoder(value.Text)
		if err != nil {
			return nil, err
		}
		buff.Write(textBytes)
	default:
		return nil, errors.New("not support localized text mask type")
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

	// todo 优化写法
	if (value.EncodingMask&0x01 != 0) {
		tempBytes, err := VariantEncoder(*value.Value)
		if err != nil {
			return nil, err
		}
		binary.Write(buff, binary.LittleEndian, tempBytes)
	}
	if value.EncodingMask&0x02 != 0 {
		binary.Write(buff, binary.LittleEndian, value.ResultStatusCode)
	}
	if value.EncodingMask&0x04 != 0 {
		binary.Write(buff, binary.LittleEndian, value.SourceTimestamp)
	}
	if value.EncodingMask&0x08 != 0 {
		binary.Write(buff, binary.LittleEndian, value.ServerTimestamp)
	}
	if value.EncodingMask&0x10 != 0 {
		binary.Write(buff, binary.LittleEndian, value.SourcePicoSeconds)
	}
	if value.EncodingMask&0x20 != 0 {
		binary.Write(buff, binary.LittleEndian, value.ServerPicoSeconds)
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

	encodeFn := func(encodingMask byte, v interface{}) ([]byte, error) {
		tempBuff := bytes.NewBuffer(nil)
		switch encodingMask & 0x3f {
		case 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0d, 0x13:
			// bool
			binary.Write(tempBuff, binary.LittleEndian, v)
		case 0x0c:
			tempBytes, err := StringEncoder(v)
			if err != nil {
				return nil, err
			}
			tempBuff.Write(tempBytes)
		case 0x0e:
			var guid uamsg.Guid
			temp, ok := v.(*uamsg.Guid)
			if !ok {
				guid, ok = v.(uamsg.Guid)
				if !ok {
					return nil, fmt.Errorf("string nodeId has wrong type identifier")
				}
				temp = &guid
			}
			guid = *temp
			binary.Write(tempBuff, binary.LittleEndian, guid.Data1)
			binary.Write(tempBuff, binary.LittleEndian, guid.Data2)
			binary.Write(tempBuff, binary.LittleEndian, guid.Data3)
			binary.Write(tempBuff, binary.LittleEndian, guid.Data4)
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
			var qualifiedName uamsg.QualifiedName
			temp, ok := v.(*uamsg.QualifiedName)
			if !ok {
				qualifiedName, ok = v.(uamsg.QualifiedName)
				if !ok {
					return nil, fmt.Errorf("string nodeId has wrong type identifier")
				}
				temp = &qualifiedName
			}
			qualifiedName = *temp
			binary.Write(tempBuff, binary.LittleEndian, qualifiedName.NamespaceIndex)
			tempBytes, err := StringEncoder(qualifiedName.Name)
			if err != nil {
				return nil, err
			}
			binary.Write(tempBuff, binary.LittleEndian, tempBytes)
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

	// 分析数据类型
	switch {
	case value.EncodingMask&0xff == 0:
		// null variant, nothing to do
	case value.EncodingMask&0x40 != 0:
		// 多维矩阵
		binary.Write(buff, binary.LittleEndian, value.ArrayLength)
		vv := reflect.ValueOf(value.Value)
		for i := 0; i < vv.Len(); i++ {
			for j := 0; j < vv.Index(i).Len(); j++ {
				tempBytes, err := encodeFn(value.EncodingMask, vv.Index(i).Index(j).Interface())
				if err != nil {
					return nil, err
				}
				buff.Write(tempBytes)
			}
		}

		binary.Write(buff, binary.LittleEndian, value.ArrayDimensionsLength)

		length := int32(len(value.ArrayDimensions))
		binary.Write(buff, binary.LittleEndian, length)
		for i := 0; i < int(length); i++ {
			binary.Write(buff, binary.LittleEndian, value.ArrayDimensions[i])
		}
	case value.EncodingMask&0x80 != 0:
		// 一维数组
		binary.Write(buff, binary.LittleEndian, value.ArrayLength)
		vv := reflect.ValueOf(value.Value)
		for i := 0; i < vv.Len(); i++ {
			tempBytes, err := encodeFn(value.EncodingMask, vv.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			buff.Write(tempBytes)
		}
	default:
		// 单值走这里
		tempBytes, err := encodeFn(value.EncodingMask, value.Value)
		if err != nil {
			return nil, err
		}
		buff.Write(tempBytes)
	}

	return buff.Bytes(), nil
}

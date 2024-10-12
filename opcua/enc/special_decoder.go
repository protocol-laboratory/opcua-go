package enc

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"opcua-go/opcua/uamsg"
)

func StringDecoder(r *superReader, v reflect.Value) error {
	b, err := r.readN(4)
	if err != nil {
		return err
	}
	length := int32(binary.LittleEndian.Uint32(b))
	if length == -1 {
		return nil
	}
	b, err = r.readN(length)
	if err != nil {
		return err
	}
	v.SetString(string(b))
	return nil
}

func GuidDecoder(r *superReader, v reflect.Value) error {
	guid := uamsg.Guid{}
	dataBytes, err := r.readN(4)
	if err != nil {
		return err
	}
	guid.Data1 = binary.LittleEndian.Uint32(dataBytes)

	dataBytes, err = r.readN(2)
	if err != nil {
		return err
	}
	guid.Data2 = binary.LittleEndian.Uint16(dataBytes)

	dataBytes, err = r.readN(2)
	if err != nil {
		return err
	}
	guid.Data3 = binary.LittleEndian.Uint16(dataBytes)

	dataBytes, err = r.readN(8)
	if err != nil {
		return err
	}
	guid.Data4 = binary.LittleEndian.Uint64(dataBytes)
	v.Set(reflect.ValueOf(guid))
	return nil
}

func ByteStringDecoder(r *superReader, v reflect.Value) error {
	b, err := r.readN(4)
	if err != nil {
		return err
	}
	length := int32(binary.LittleEndian.Uint32(b))
	if length == -1 {
		return nil
	}
	b, err = r.readN(length)
	if err != nil {
		return err
	}
	v.SetBytes(b)
	return nil
}

type SpecialDecoder func(r *superReader, v reflect.Value) error

// SpecialStructDecoderMap Such types belong to the built-in types of the OPC UA protocol, but require special adaptation of the coding language.
var SpecialStructDecoderMap = map[string]SpecialDecoder{
	"NodeId":          NodeIdDecoder,
	"ExpandedNodeId":  ExpandedNodeIdDecoder,
	"ExtensionObject": ExtensionObjectDecoder,
	"DiagnosticInfo":  DiagnosticInfoDecoder,
	"LocalizedText":   LocalizedTextDecoder,
	"DataValue":       DataValueDecoder,
	"Variant":         VariantDecoder,
}

func NodeIdDecoder(r *superReader, v reflect.Value) error {
	tempData := uamsg.NodeId{}
	encodingType, err := r.readByte()
	if err != nil {
		return err
	}
	tempData.EncodingType = uamsg.NodeIdEncodingType(encodingType)

	// node id uses the lower 4 bits of byte
	switch tempData.EncodingType & 0x0f {
	case uamsg.TwoByte:
		dataByte, err := r.readByte()
		if err != nil {
			return err
		}
		tempData.Identifier = dataByte
	case uamsg.FourByte:
		dataByte, err := r.readByte()
		if err != nil {
			return err
		}
		tempData.Namespace = uint16(dataByte)
		dataBytes, err := r.readN(2)
		if err != nil {
			return err
		}
		tempData.Identifier = binary.LittleEndian.Uint16(dataBytes)
	case uamsg.Numeric:
		dataBytes, err := r.readN(2)
		if err != nil {
			return err
		}
		tempData.Namespace = binary.LittleEndian.Uint16(dataBytes)
		dataBytes, err = r.readN(4)
		if err != nil {
			return err
		}
		tempData.Identifier = binary.LittleEndian.Uint32(dataBytes)
	case uamsg.String:
		dataBytes, err := r.readN(2)
		if err != nil {
			return err
		}
		tempData.Namespace = binary.LittleEndian.Uint16(dataBytes)

		tempStr := ""
		err = StringDecoder(r, reflect.ValueOf(&tempData).Elem())
		if err != nil {
			return err
		}
		tempData.Identifier = tempStr
	case uamsg.GuidType:
		dataBytes, err := r.readN(2)
		if err != nil {
			return err
		}
		tempData.Namespace = binary.LittleEndian.Uint16(dataBytes)
		guid := &uamsg.Guid{}
		err = GuidDecoder(r, reflect.ValueOf(guid).Elem())
		if err != nil {
			return err
		}
		tempData.Identifier = guid
	case uamsg.ByteString:
		dataBytes, err := r.readN(2)
		if err != nil {
			return err
		}
		tempData.Namespace = binary.LittleEndian.Uint16(dataBytes)

		var tempByteStr []byte
		err = StringDecoder(r, reflect.ValueOf(&tempData).Elem())
		if err != nil {
			return err
		}
		tempData.Identifier = tempByteStr
	default:
		return fmt.Errorf("nodeId has unknown type identifier")
	}
	v.Set(reflect.ValueOf(tempData))
	return nil
}

func ExpandedNodeIdDecoder(r *superReader, v reflect.Value) error {
	tempData := uamsg.ExpandedNodeId{}
	nodeId := &uamsg.NodeId{}
	err := NodeIdDecoder(r, reflect.ValueOf(nodeId).Elem())
	if err != nil {
		return err
	}
	tempData.NodeId = nodeId

	if (tempData.EncodingType & uamsg.NamespaceUriFlag) != 0 {
		if err = StringDecoder(r, reflect.ValueOf(&tempData.NamespaceUri).Elem()); err != nil {
			return err
		}
	}
	if (tempData.EncodingType & uamsg.ServerIndexFlag) != 0 {
		dataBytes, err := r.readN(4)
		if err != nil {
			return err
		}
		tempData.ServerIndex = binary.LittleEndian.Uint32(dataBytes)
	}
	v.Set(reflect.ValueOf(tempData))
	return nil
}

func ExtensionObjectDecoder(r *superReader, v reflect.Value) error {
	tempData := uamsg.ExtensionObject{}
	typeId := &uamsg.NodeId{}
	err := NodeIdDecoder(r, reflect.ValueOf(typeId).Elem())
	if err != nil {
		return err
	}
	tempData.TypeId = typeId
	encoding, err := r.readByte()
	if err != nil {
		return err
	}
	tempData.Encoding = encoding
	switch tempData.Encoding {
	case 0x00:
	case 0x01:
		fallthrough
	case 0x02:
		err = StringDecoder(r, reflect.ValueOf(&tempData.Body).Elem())
		if err != nil {
			return err
		}
	}
	v.Set(reflect.ValueOf(tempData))
	return nil
}

func DiagnosticInfoDecoder(r *superReader, v reflect.Value) error {
	tempData := uamsg.DiagnosticInfo{}
	encodingMask, err := r.readByte()
	if err != nil {
		return err
	}
	tempData.EncodingMask = encodingMask

	if tempData.EncodingMask&uamsg.SymbolicIdFlag != 0 {
		dataBytes, err := r.readN(4)
		if err != nil {
			return err
		}
		tempData.SymbolicId = int32(binary.LittleEndian.Uint32(dataBytes))
	}
	if tempData.EncodingMask&uamsg.NamespaceFlag != 0 {
		dataBytes, err := r.readN(4)
		if err != nil {
			return err
		}
		tempData.NamespaceUri = int32(binary.LittleEndian.Uint32(dataBytes))
	}
	if tempData.EncodingMask&uamsg.LocaleFlag != 0 {
		dataBytes, err := r.readN(4)
		if err != nil {
			return err
		}
		tempData.Locale = int32(binary.LittleEndian.Uint32(dataBytes))
	}
	if tempData.EncodingMask&uamsg.LocalizedTextFlag != 0 {
		dataBytes, err := r.readN(4)
		if err != nil {
			return err
		}
		tempData.LocalizedText = int32(binary.LittleEndian.Uint32(dataBytes))
	}
	if tempData.EncodingMask&uamsg.AdditionalInfoFlag != 0 {
		err := StringDecoder(r, reflect.ValueOf(&tempData.AdditionalInfo).Elem())
		if err != nil {
			return err
		}
	}
	if tempData.EncodingMask&uamsg.InnerStatusCodeFlag != 0 {
		dataBytes, err := r.readN(4)
		if err != nil {
			return err
		}
		tempData.InnerStatusCode = binary.LittleEndian.Uint32(dataBytes)
	}
	if tempData.EncodingMask&uamsg.InnerDiagnosticInfoFlag != 0 {
		innerDiagnosticInfo := &uamsg.DiagnosticInfo{}
		err := DiagnosticInfoDecoder(r, reflect.ValueOf(innerDiagnosticInfo).Elem())
		if err != nil {
			return err
		}
		tempData.InnerDiagnosticInfo = innerDiagnosticInfo
	}
	v.Set(reflect.ValueOf(tempData))
	return nil
}

func LocalizedTextDecoder(r *superReader, v reflect.Value) error {
	tempData := uamsg.LocalizedText{}
	encodingMask, err := r.readByte()
	if err != nil {
		return err
	}
	tempData.EncodingMask = encodingMask

	if tempData.EncodingMask&0x01 != 0 {
		err := StringDecoder(r, reflect.ValueOf(&tempData.Locale).Elem())
		if err != nil {
			return err
		}
	}
	if tempData.EncodingMask&0x02 != 0 {
		err := StringDecoder(r, reflect.ValueOf(&tempData.Text).Elem())
		if err != nil {
			return err
		}
	}
	v.Set(reflect.ValueOf(tempData))
	return nil
}

func DataValueDecoder(r *superReader, v reflect.Value) error {
	return nil
}

func VariantDecoder(r *superReader, v reflect.Value) error {
	return nil
}

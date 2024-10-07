package uamsg

type ReadRequest struct {
	Header             *RequestHeader
	MaxAge             Duration
	TimestampsToReturn TimestampsToReturnEnum
	NodesToRead        []*ReadValueId
}

type ReadValueId struct {
	NodeIdToRead *NodeId
	AttributeId  IntegerId
	IndexRange   NumericRange
	DataEncoding *QualifiedName
}

type ReadResponse struct {
	Header          *ResponseHeader
	Results         []*DataValue
	DiagnosticInfos []DiagnosticInfo
}

type DataValue struct {
	EncodingMask      byte
	Value             *Variant
	ResultStatusCode  StatusCode
	SourceTimestamp   uint64
	SourcePicoSeconds uint16
	ServerTimestamp   uint64
	ServerPicoSeconds uint16
}

type Variant struct {
	EncodingMask          byte
	ArrayLength           int32
	Value                 interface{}
	ArrayDimensionsLength int32
	ArrayDimensions       []int32
}

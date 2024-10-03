package uamsg

import "time"

type ReadRequest struct {
	Header             *RequestHeader
	MaxAge             time.Duration
	TimestampsToReturn TimestampsToReturnEnum
	NodesToRead        []ReadValueId
}

type ReadValueId struct {
	NodeIdToRead *NodeId
	AttributeId  IntegerId
	IndexRange   NumericRange
	DataEncoding QualifiedName
}

type ReadResponse struct {
	Header          *ResponseHeader
	Results         []DataValue
	DiagnosticInfos []DiagnosticInfo
}

type DataValue struct {
	Value             BaseDataType
	ResultStatusCode  StatusCode
	SourceTimestamp   uint32
	SourcePicoSeconds uint64
	ServerTimestamp   uint32
	ServerPicoSeconds uint64
}

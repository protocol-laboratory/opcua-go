package uamsg

type BrowseRequest struct {
	Header                        *RequestHeader
	View                          *ViewDescription
	RequestedMaxReferencesPerNode uint32
	NodesToBrowse                 []*BrowseDescription
}

type ViewDescription struct {
	ViewId      *NodeId
	Timestamp   uint64
	ViewVersion uint32
}

type BrowseDescription struct {
	NodeId          *NodeId
	BrowseDirection uint32
	ReferenceTypeId *NodeId
	IncludeSubtypes bool
	NodeClassMask   uint32
	ResultMask      uint32
}

type BrowseResponse struct {
	Header  *ResponseHeader
	Results []*BrowseResult
}

type BrowseResult struct {
	StatusCode        *StatusCode
	ContinuationPoint []byte
	References        []*ReferenceDescription
	DiagnosticInfos   []*DiagnosticInfo
}

type ReferenceDescription struct {
	ReferenceTypeId *NodeId
	IsForward       bool
	NodeId          *ExpandedNodeId
	BrowseName      *QualifiedName
	DisplayName     *LocalizedText
	NodeClass       uint32
	TypeDefinition  *ExpandedNodeId
}

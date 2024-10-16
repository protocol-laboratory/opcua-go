package uamsg

type GetEndpointsRequest struct {
	Header      *RequestHeader
	EndpointUrl string
	LocaleIds   []string
	ProfileUris []string
}

type GetEndpointsResponse struct {
	Header    *ResponseHeader
	Endpoints []EndpointDescription
}

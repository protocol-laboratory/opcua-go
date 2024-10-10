package errorcode

const (
	BadDecoding = iota
	BadEncoding
)

// ErrorCodes definition`link： https://github.com/OPCFoundation/UA-Nodeset/blob/UA-1.05.03-2023-12-15/Schema/StatusCode.csv
// used to fill error response
var ErrorCodes = map[uint32]string{
	BadDecoding: "BAD_DECODING",
}

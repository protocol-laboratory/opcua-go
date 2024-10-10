package util

type BuiltInType interface {
	int32 | int64 |
		uint8 | uint32 | uint64 |
		string | bool
}

func GetPtr[T BuiltInType](v T) *T {
	return &v
}

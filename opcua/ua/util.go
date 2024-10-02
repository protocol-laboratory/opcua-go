package ua

func StrLen(str string) int {
	return 4 + len([]byte(str))
}

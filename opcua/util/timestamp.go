package util

import "time"

const (
	// OpcuaDateTimeTicksToUnixEpoch ((1970 - 1601) * 365 + 3 * 24 + 17) * 86400 * 10 * 1000 * 1000
	OpcuaDateTimeTicksToUnixEpoch int64 = 116444736000000000
)

func GetCurrentUaTimestamp() uint64 {
	ticks := time.Now().UTC().UnixNano()/100 + OpcuaDateTimeTicksToUnixEpoch
	if ticks < 0 {
		ticks = 0
	}

	return uint64(ticks)
}

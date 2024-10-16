package opcua

import "sync/atomic"

type ChannelIdGen struct {
	nextChannelId atomic.Uint32
}

func (c *ChannelIdGen) next() uint32 {
	return c.nextChannelId.Add(1)
}

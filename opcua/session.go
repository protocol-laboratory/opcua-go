package opcua

import (
	"time"

	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
	"github.com/protocol-laboratory/opcua-go/opcua/util"
	"golang.org/x/exp/rand"
)

type Session struct {
	sessionId   uamsg.NodeId
	sessionName string
	serverNonce []byte

	// TODO should check expiration -- func (s *Session) IsExpired() bool
	requestedSessionTimeout time.Duration
	maxResponseMessageSize  uint32
}

func newSession(sessionName string, requestedSessionTimeout uamsg.Duration, maxResponseMessageSize uint32) *Session {
	return &Session{
		// TODO should assign one if no session provided
		sessionId:               getUniqueSessionId(),
		sessionName:             sessionName,
		serverNonce:             util.GenerateRandomBytes(32),
		requestedSessionTimeout: time.Duration(requestedSessionTimeout) * time.Millisecond,
		maxResponseMessageSize:  maxResponseMessageSize,
	}
}

func getUniqueSessionId() uamsg.NodeId {
	return uamsg.NodeId{
		EncodingType: uamsg.GuidType,
		Namespace:    1,
		Identifier: uamsg.Guid{
			Data1: rand.Uint32(),
			Data2: uint16(rand.Uint32() % 65535),
			Data3: uint16(rand.Uint32() % 65535),
			Data4: rand.Uint64(),
		},
	}
}

func getUniqueSessionAuthenticationToken() uamsg.SessionAuthenticationToken {
	return uamsg.SessionAuthenticationToken{
		EncodingType: uamsg.ByteString,
		Namespace:    0,
		Identifier:   util.GenerateRandomBytes(32),
	}
}

package opcua

import (
	"github.com/libgox/gocollections/syncx"
	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
)

type SessionManager struct {
	sessions syncx.Map[uamsg.SessionAuthenticationToken, *Session]
}

func newSessionManager() *SessionManager {
	return &SessionManager{}
}

func (sm *SessionManager) add(sessionAuthenticationToken uamsg.SessionAuthenticationToken, session *Session) {
	sm.sessions.Store(sessionAuthenticationToken, session)
}

func (sm *SessionManager) delete(sessionAuthenticationToken uamsg.SessionAuthenticationToken) {
	sm.sessions.Delete(sessionAuthenticationToken)
}

func (sm *SessionManager) get(sessionAuthenticationToken uamsg.SessionAuthenticationToken) (*Session, bool) {
	return sm.sessions.Load(sessionAuthenticationToken)
}

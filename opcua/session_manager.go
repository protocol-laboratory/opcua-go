package opcua

import (
	"github.com/libgox/gocollections/syncx"
)

type SessionManager struct {
	sessions syncx.Map[string, *Session]
}

func newSessionManager() *SessionManager {
	return &SessionManager{}
}

func (sm *SessionManager) add(sessionAuthenticationToken string, session *Session) {
	sm.sessions.Store(sessionAuthenticationToken, session)
}

func (sm *SessionManager) delete(sessionAuthenticationToken string) {
	sm.sessions.Delete(sessionAuthenticationToken)
}

func (sm *SessionManager) get(sessionAuthenticationToken string) (*Session, bool) {
	return sm.sessions.Load(sessionAuthenticationToken)
}

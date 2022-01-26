package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"sync"
	"time"
)

type Manager struct {
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, maxlifetime: maxlifetime}, nil
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(sessionId string) (session Session, err error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	if sessionId == "" {
		sid := manager.sessionId()
		session, err = manager.provider.SessionInit(sid)
	} else {
		session, err = manager.provider.SessionRead(sessionId)
	}
	return
}

// SessionDestroy Destroys session with given ID
func (manager *Manager) SessionDestroy(sessionId string) error {
	if sessionId != "" {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		err := manager.provider.SessionDestroy(sessionId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (manager *Manager) GC(ctx context.Context) {
	for true {
		if ctx.Err() != nil {
			return
		}

		time.AfterFunc(time.Duration(manager.maxlifetime), func() {
			manager.lock.Lock()
			defer manager.lock.Unlock()
			manager.provider.SessionGC(manager.maxlifetime)
		})
	}
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var provides = make(map[string]Provider)

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil,
// it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

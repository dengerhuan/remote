package session

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"log"
	"sync"
)

var SessionManagerInst = NewManager()

type SessionManager interface {
	netty.ActiveHandler
	netty.InactiveHandler
	SessionAt(id int64) *Session
	Size() int
}

func NewManager() SessionManager {
	return &sessionManager{sessions: make(map[int64]Session, 64),}
}

type Session struct {
	context netty.HandlerContext
	attr    map[string]interface{}
}

func (s *Session) Attr(key string) (val interface{}, ok bool) {
	val, ok = s.attr[key]
	return val, ok
}

func (s *Session) SetAttr(key string, attr interface{}) {
	s.attr[key] = attr
}

type sessionManager struct {
	sessions map[int64]Session
	mutex    sync.RWMutex
}

func (s *sessionManager) Size() int {
	s.mutex.RLock()
	size := len(s.sessions)
	s.mutex.RUnlock()
	return size

}

func (s *sessionManager) SessionAt(id int64) *Session {
	s.mutex.RLock()
	se, _ := s.sessions[id]
	s.mutex.RUnlock()
	return &se
}

func (s *sessionManager) HandleActive(ctx netty.ActiveContext) {

	log.Println("session active:", ctx.Channel().RemoteAddr())

	s.mutex.RLock()
	s.sessions[ctx.Channel().ID()] = Session{context: ctx, attr: make(map[string]interface{}, 10)}
	s.mutex.RUnlock()
	ctx.HandleActive()
}

func (s *sessionManager) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	log.Println("session inactive ", ctx.Channel().RemoteAddr())
	s.mutex.RLock()
	delete(s.sessions, ctx.Channel().ID())
	s.mutex.RUnlock()
	ctx.HandleInactive(ex)
}

func (c *sessionManager) HandleException(ex netty.Exception) {
	fmt.Println(ex)
}

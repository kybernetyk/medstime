package main

import (
	//    "os"
	"crypto/md5"
	"crypto/rand"
	"math/big"
	"fmt"
	"github.com/hoisie/web.go"
	"sync"
	"time"
	"log"
)

type SessionManager struct {
	sessions map[string]*Session //sessions by sessionid

	//TODO: needs mutexing
	mu sync.RWMutex
}

func NewSessionManager() *SessionManager {
	mgr := new(SessionManager)
	mgr.sessions = make(map[string]*Session)

	return mgr
}

func (self *SessionManager) CurrentSession(ctx *web.Context) (session *Session) {

	session_id, ok := ctx.GetSecureCookie("session_id")
	log.Println("session_id = ", session_id, "ok = ", ok)
	if !ok {
		session = self.createSession(ctx)
		return
	}

	session, ok = self.sessionForSessionId(session_id)
	log.Println("session => ", session, " ok => ", ok)
	if !ok {
		session = self.createSession(ctx)
		return
	}

	return
}

func (self *SessionManager) DestroyCurrentSession(ctx *web.Context) {
	session := self.CurrentSession(ctx)
	delete(self.sessions, session.Id)
	ctx.SetSecureCookie("session_id", "", session.TimeoutAfter)
}

func (self *SessionManager) sessionForSessionId(ses_id string) (session *Session, ok bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	log.Println("retrieving session for id ", ses_id)
	session, ok = self.sessions[ses_id]
	log.Println("session: ", session, " ok: ", ok)
	if !ok {
		return
	}

	//check for timeout
	now := time.Now()
	if (now.Sub(session.LastActive)) > session.TimeoutAfter {
		log.Println("session timed out!")
		log.Println("last active: ", session.LastActive)
		log.Println("sub(active): ", now.Sub(session.LastActive))
		log.Println("TimeoutAfter: ", session.TimeoutAfter)
		delete(self.sessions, session.Id)
		ok = false
		return
	}
	//	session.LastActive = now
	ok = true
	return
}

func (self *SessionManager) createSession(ctx *web.Context) (session *Session) {
	self.mu.Lock()
	defer self.mu.Unlock()

	ses := new(Session)
	ses.Data = make(map[string]interface{})
	ses.LastActive = time.Now()
	ses.TimeoutAfter = 3600 * time.Second

	rnd1, _ := rand.Int(rand.Reader, big.NewInt(0xffffffff))
	rnd2, _ := rand.Int(rand.Reader, big.NewInt(0xffffffff))

	ses.Id = md5Hash(fmt.Sprintf("%d%d%d", rnd1 , time.Now(), rnd2))

	self.sessions[ses.Id] = ses
	session = self.sessions[ses.Id]

	ctx.SetSecureCookie("session_id", session.Id, session.TimeoutAfter)
	log.Println("new session: %#v", session)
	return
}

func md5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

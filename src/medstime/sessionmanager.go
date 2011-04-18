package main

import (
	//    "os"
	"rand"
	"fmt"
	"time"
	"sync"
	"crypto/md5"
	"web"
)

type SessionManager struct {
	Sessions map[string]*Session //sessions by sessionid

	//TODO: needs mutexing
	mu sync.RWMutex
}


func NewSessionManager() *SessionManager {
	mgr := new(SessionManager)
	mgr.Sessions = make(map[string]*Session)

	return mgr
}


func (self *SessionManager) SessionForSessionId(ses_id string) (session *Session, ok bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	session, ok = self.Sessions[ses_id]
	if !ok {
		//  err = os.NewError("No Session with this ID found!")
		return
	}

	//check for timeout
	now := time.Seconds()
	if (now - session.LastActive) > session.TimeoutAfter {
		//err = os.NewError("Session timed out")
		//self.Sessions[ses_id] = nil, false //make a cleanup method that will be called periodically
		ok = false
		return
	}
	session.LastActive = now
	ok = true
	return
}

func (self *SessionManager) CurrentSession(ctx *web.Context) (session *Session) {
	session_id, ok := ctx.GetSecureCookie("session_id")
	if !ok {
		session = self.createSession(ctx)
		return
	}

	session, ok = self.SessionForSessionId(session_id)
	if !ok {
		session = self.createSession(ctx)
		return
	}

	return
}

func (self *SessionManager) createSession(ctx *web.Context) (session *Session) {
	self.mu.Lock()
	defer self.mu.Unlock()

	ses := new(Session)
	ses.Data = make(map[string]interface{})
	ses.LastActive = time.Seconds()
	ses.TimeoutAfter = 60 * 60 //1 hour

	ses.Id = md5Hash(fmt.Sprintf("%d%d%d", rand.Int31(), time.Seconds(), rand.Int31()))

	self.Sessions[ses.Id] = ses
	session = self.Sessions[ses.Id]

	ctx.SetSecureCookie("session_id", session.Id, session.TimeoutAfter)
	return
}

func md5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return fmt.Sprintf("%x", hasher.Sum())
}

package main

import (
	//    "os"
	"fmt"
	"time"
	"sync"
	"crypto/md5"
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

func (self *SessionManager) CreateSession(setter CookieSetter) (session *Session, ok bool) {
	self.mu.Lock()
	defer self.mu.Unlock()

	ses := new(Session)
	ses.Data = make(map[string]interface{})
	ses.LastActive = time.Seconds()
	ses.TimeoutAfter = 60 * 60 //1 hour
	self.Sessions[ses.Id] = ses

	session = self.Sessions[ses.Id]

	setter.SetSecureCookie("session_id", session.Id, session.TimeoutAfter)
	
    ok = true
    return
}

func (self *SessionManager) CreateSessionForAccount(setter CookieSetter, acc Account) (session *Session, ok bool) {
    session, ok = self.CreateSession(setter)
    if !ok {
        return
    }

    session.Id = md5Hash(fmt.Sprintf("%s%d%s", acc.Username, time.Seconds(), acc.Password))
	session.AccountId = acc.Id

	return
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

func (self *SessionManager) CurrentSession(getter CookieGetter) (session *Session, ok bool) {
	session_id, ok := getter.GetSecureCookie("session_id")
	if !ok {
		//err = os.NewError("No Cookie")
		ok = false
		return
	}

	session, ok = self.SessionForSessionId(session_id)

	return
}


func md5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return fmt.Sprintf("%x", hasher.Sum())
}

//let's add a little helper to web.go's web.Context
type CookieSetter interface {
    SetSecureCookie(name string, val string, age int64)
}

type CookieGetter interface {
    GetSecureCookie(name string) (string, bool)
}


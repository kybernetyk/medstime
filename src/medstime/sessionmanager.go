package main

import (
    "os"
    "fmt"
    "time"
    "sync"
    "crypto/md5"
)

type SessionManager struct {
    Sessions    map[string]*Session  //sessions by sessionid
    
    //TODO: needs mutexing
    mu  sync.RWMutex
}

var g_SessionManager *SessionManager

func SharedSessionManager() *SessionManager {
    if g_SessionManager == nil {
        g_SessionManager = new(SessionManager)
        g_SessionManager.Sessions = make(map[string]*Session)
    }
    
    return g_SessionManager
}

func (self *SessionManager) NewSessionForAccount(acc Account) (session *Session, err os.Error) {
    self.mu.Lock()
    defer self.mu.Unlock()
  
    ses := new(Session)
    ses.Id = md5Hash(fmt.Sprintf("%s%d%s",acc.Username, time.Seconds(), acc.Password))
    ses.AccountId = acc.Id
    ses.LastActive = time.Seconds()
    ses.TimeoutAfter = 60 * 60  //1 hour
    self.Sessions[ses.Id] = ses
    
    session = self.Sessions[ses.Id]
    return
}

func (self *SessionManager) SessionForSessionId(ses_id string) (session *Session, err os.Error) {
    self.mu.RLock()
    defer self.mu.RUnlock()
    
    var ok bool
    session, ok = self.Sessions[ses_id]
    if (!ok) {
        err = os.NewError("No Session with this ID found!")
        return
    }
    
    //check for timeout
    now := time.Seconds()
    if (now - session.LastActive) > session.TimeoutAfter {
        err = os.NewError("Session timed out")
        //self.Sessions[ses_id] = nil, false //make a cleanup method that will be called periodically
        return
    }
    session.LastActive = now
    
    return
}

func md5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return fmt.Sprintf("%x", hasher.Sum())
}


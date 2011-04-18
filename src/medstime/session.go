package main

import (
//    "os"
)

type Session struct { 
	Id string

	LastActive   int64 //unix timestamp
	TimeoutAfter int64 //seconds

	AccountId int64
}

//let's add a little helper to web.go's web.Context
type CookieSetter interface {
    SetSecureCookie(name string, val string, age int64)
}

type CookieGetter interface {
    GetSecureCookie(name string) (string, bool)
}

func Cookie_StoreSession(setter CookieSetter, session *Session) {
    setter.SetSecureCookie("session_id", session.Id, session.TimeoutAfter)
}

func Cookie_RetrieveSession(getter CookieGetter) (session *Session, ok bool) {
    session_id, ok := getter.GetSecureCookie("session_id")
    if !ok {
        //err = os.NewError("No Cookie")
        ok = false
        return
    }
    
    mgr := SharedSessionManager()
    session, ok = mgr.SessionForSessionId(session_id)
        
    return 
}

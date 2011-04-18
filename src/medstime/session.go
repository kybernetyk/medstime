package main

import (
    "time"
    "web"
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

//retrieve session from cookie and return it, or redirect to location if no valid session found
func SessionOrRedirect(ctx *web.Context, location string) *Session {
    session, ok := Cookie_RetrieveSession(ctx)
    if !ok {
        ctx.Redirect(301, location)
    } else {
        session.LastActive = time.Seconds()
    }
    return session
}

//if the user session is valid redirect to location and return true. else do nothing and return false
func RedirectIfSession(ctx *web.Context, location string) bool {
    session, ok := Cookie_RetrieveSession(ctx)
    if ok {
        session.LastActive = time.Seconds()
        ctx.Redirect(301, location)
        return true
    } 
    return false
}


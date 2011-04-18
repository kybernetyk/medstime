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
	
	Data map[string]interface{}
}



//retrieve session from cookie and return it, or redirect to location if no valid session found
func SessionOrRedirect(ctx *web.Context, location string) *Session {
    session, ok := app.SessionMgr.CurrentSession(ctx)
    if !ok {
        ctx.Redirect(301, location)
    } else {
        session.LastActive = time.Seconds()
    }
    return session
}

//if the user session is valid redirect to location and return true. else do nothing and return false
func RedirectIfSession(ctx *web.Context, location string) bool {
    session, ok := app.SessionMgr.CurrentSession(ctx)
    if ok {
        session.LastActive = time.Seconds()
        ctx.Redirect(301, location)
        return true
    } 
    return false
}


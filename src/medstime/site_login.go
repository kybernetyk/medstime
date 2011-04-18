package main

import (
    "web"
    "mustache"
//    "fmt"
)

const (
    err_None = ""
    err_NoUsername = "1"
    err_NoPass = "2"
    err_LoginFailed = "3"
    err_Critical = "99"
)

var errmap = map[string]string {
    err_None: "",
    err_NoUsername: "No Username",
    err_NoPass: "No Password",
    err_LoginFailed: "Login Wrong",
    err_Critical: "Critical Error! O M G!",
}

func loginGet(ctx *web.Context) {
    session := app.SessionMgr.CurrentSession(ctx)
    if session.GetBool("logged_in") {
        ctx.Redirect(301, "/account")
        return
    }
   

    m := map[string]string {
        
    }
    err, ok := ctx.Params["err"]
    if ok {
        e, o := errmap[err]
        if o {
            m["Error"] = e
        }
    }
        
    s := mustache.RenderFile("templ/login.mustache", &m)
    ctx.WriteString(s)
}

func loginPost(ctx *web.Context) {
    username, ok := ctx.Params["username"]
    if !ok {
        ctx.Redirect(301, "/login?err=1")
        return
    }

    password, ok := ctx.Params["password"]
    if !ok {
        ctx.Redirect(301, "/login?err=2")
        return
    }
    
    accmgr := NewAccountManager()
    acc, ok := accmgr.AccountForUsernamePassword(username, password)
    if !ok {
        ctx.Redirect(301, "/login?err=3")
        return
    }
    
    ses := app.SessionMgr.CurrentSession(ctx)
    ses.Set("account_id", acc.Id)
    ses.Set("logged_in", true)
    
    ctx.Redirect(301, "/account")
}
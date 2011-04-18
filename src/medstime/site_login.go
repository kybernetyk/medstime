package main

import (
    "web"
    "mustache"
//    "fmt"
)

func loginGet(ctx *web.Context) {
    session := app.SessionMgr.CurrentSession(ctx)
    if session.GetBool("logged_in") {
        ctx.Redirect(301, "/account")
        return
    }
   

    m := map[string]string {
        
    }
    
    estr, ok := GetErrorString(ctx)
    if ok {
        m["Error"] = estr
    }
        
    s := mustache.RenderFile("templ/login.mustache", &m)
    ctx.WriteString(s)
}

func loginPost(ctx *web.Context) {
    username, ok := ctx.Params["username"]
    if !ok {
        ctx.Redirect(301, "/login?err=" + err_LoginNoUsername)
        return
    }

    password, ok := ctx.Params["password"]
    if !ok {
        ctx.Redirect(301, "/login?err=" + err_LoginNoPass)
        return
    }
    
    accmgr := NewAccountManager()
    acc, ok := accmgr.AccountForUsernamePassword(username, password)
    if !ok {
        ctx.Redirect(301, "/login?err=" + err_LoginFailed)
        return
    }
    
    ses := app.SessionMgr.CurrentSession(ctx)
    ses.Set("account_id", acc.Id)
    ses.Set("logged_in", true)
    
    ctx.Redirect(301, "/account")
}

func logoutGet(ctx *web.Context) {
    app.SessionMgr.DestroyCurrentSession(ctx)
    ctx.Redirect(301, "/")
}
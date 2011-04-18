package main

import (
    "web"
    "mustache"
//    "fmt"
)

func signupGet(ctx *web.Context) {
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
        
    s := mustache.RenderFile("templ/signup.mustache", &m)
    ctx.WriteString(s)
}

func signupPost(ctx *web.Context) {
    username, ok := ctx.Params["username"]
    if !ok || len(username) == 0 {
        ctx.Redirect(301, "/signup?err=" + err_SignupUsernameInvalid)
        return
    }

    password, ok := ctx.Params["password"]
    if !ok || len(password) == 0 {
        ctx.Redirect(301, "/signup?err=" + err_SignupPasswordInvalid)
        return
    }
    
    acc := Account{
        Username: username,
        Password: password,
        Plan: 0,
        Id: 0,
    }
    
    accmgr := NewAccountManager()
    acc_id, err := accmgr.CreateAccount(acc)
    if err != nil {
        ctx.Redirect(301, "/signup?err=" + err.String())
        return
    }
    
    ses := app.SessionMgr.CurrentSession(ctx)
    ses.Set("account_id", acc_id)
    ses.Set("logged_in", true)
    
    ctx.Redirect(301, "/account")
}
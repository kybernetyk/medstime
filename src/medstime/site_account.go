package main

import (
    "web"
    "mustache"
    "fmt"
//    "time"
)

func accountGet(ctx *web.Context) {
    session := app.SessionMgr.CurrentSession(ctx)
    if !session.GetBool("logged_in") {
        ctx.Redirect(301, "/login")
        return
    }

    accmgr := NewAccountManager()
    acc, _ := accmgr.AccountForAccountId(session.GetInt64("account_id"))
    
    m := map[string]string {
        "Debug": fmt.Sprintf("%#v<br><br>%#v", session, acc),
    }
    
    mgr := NewScheduleManager()
    items := mgr.ScheduleItemsForAccount(acc)
    fmt.Printf("%#v - %d\n", items, len(items))

    s := mustache.RenderFile("templ/account.mustache", &m)
    ctx.WriteString(s)
}


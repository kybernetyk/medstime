package main

import (
    "web"
    "mustache"
//    "fmt"
)

func indexGet(ctx *web.Context) string {
    ses := app.SessionMgr.CurrentSession(ctx)

    m := map[string]interface{}{}
    
    acc_id := ses.GetInt("account_id")
    if acc_id > 0 {
        m["Accid"] = acc_id
    }
    
    s := mustache.RenderFile("templ/index.mustache", &m)
    return s
}
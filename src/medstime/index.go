package main

import (
    "web"
    "mustache"
)

func index(ctx *web.Context) string {
    tmpl, _ := mustache.ParseFile("templ/index.mustache")
    
    session, err := Cookie_RetrieveSession(ctx)
    if err != nil {
        acc := Account{Id: 1234, Username: "joorek", Password: "warbird"}
        ses, err := SharedSessionManager().NewSessionForAccount(acc)
        if err != nil {
            return "couldn't create new session!"
        }
        Cookie_StoreSession(ctx, ses)
        return "stored"
    }

    m := map[string]interface{} {
        "Userinfo": map[string]string {"Name": "Jarek"},
        "SessionId": session.Id,
        "SessionAcc": session.AccountId,
    }
    s := tmpl.Render(&m)
    
    return s
}
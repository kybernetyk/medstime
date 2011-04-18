package main

import (
    "web"
//    "mustache"
    "fmt"
//    "time"
)

func accountGet(ctx *web.Context) {
    session := SessionOrRedirect(ctx, "/login")
    if session == nil {
        return
    }
    
    ctx.WriteString(fmt.Sprintf("%#v", session))
    
    // m := map[string]string {
    //     
    // }
    // err, ok := ctx.Params["err"]
    // if ok {
    //     e, o := errmap[err]
    //     if o {
    //         m["Error"] = e
    //     }
    // }
    //     
    // s := mustache.RenderFile("templ/login.mustache", &m)
    // ctx.WriteString(s)
}


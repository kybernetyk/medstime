package main

import (
	"web"
	"mustache"
	    "fmt"
)

func indexGet(ctx *web.Context) string {
	ses := app.SessionMgr.CurrentSession(ctx)
	fmt.Println(ses)
	m := map[string]interface{}{}

	acc_id := ses.GetInt("account_id")
	if acc_id > 0 {
		m["Accid"] = acc_id
	}
	fmt.Println(m)
	s := mustache.RenderFile("templ/index.mustache", &m)
	return s
}

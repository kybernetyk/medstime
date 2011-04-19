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

	m := map[string]interface{}{
		"Debug": fmt.Sprintf("%#v<br><br>%#v", session, acc),
	}

	mgr := NewScheduleManager()
	// id := mgr.AddScheduleItemToAccount(acc)
	// fmt.Printf("added schedule item %d to acc\n", id)
	// 
	// items := mgr.ScheduleItemsForAccount(acc)
	// fmt.Printf("%#v - %d\n", items, len(items))
	// si := items[0]
	// si.Hour ++
	// si.Message = "Hallo Du!"
	// mgr.UpdateScheduleItem(si)

	items := mgr.ScheduleItemsForAccount(acc)
	fmt.Printf("%#v - %d\n", items, len(items))

	m["Schedules"] = items

	s := mustache.RenderFile("templ/account.mustache", &m)
	ctx.WriteString(s)
}

func hoursList(selected_hour string) []interface{} {
    type MyHour struct {
		Hour     string
		Selected string
	}

	var hours []interface{}
	for i := 0; i < 24; i++ {
		selected := ""
		hstr := fmt.Sprintf("%d", i)
		if hstr == selected_hour {
			selected = "selected"
		}
		h := MyHour{
			Hour:     hstr,
			Selected: selected,
		}
		hours = append(hours, h)
	}
	
	return hours
}

func minutesList(selected_minute string) []interface{} {
    type MyMinute struct {
		Minute   string
		Selected string
	}


	var minutes []interface{}
	for i := 0; i < 60; i += 15 {
		selected := ""
		mstr := fmt.Sprintf("%d", i)
		if selected_minute == mstr {
			selected = "selected"
		}

		m := MyMinute{
			Minute:   mstr,
			Selected: selected,
		}
		minutes = append(minutes, m)
	}
	
	return minutes
	
}

func accountNewScheduleGet(ctx *web.Context) {
	session := app.SessionMgr.CurrentSession(ctx)
	if !session.GetBool("logged_in") {
		ctx.Redirect(301, "/login")
		return
	}

	m := map[string]interface{}{
		"Debug": ctx.Params,
	}

	type MyHour struct {
		Hour     string
		Selected string
	}

	m["Hours"] = hoursList("")
    m["Minutes"] = minutesList("")

	s := mustache.RenderFile("templ/account_newschedule.mustache", &m)
	ctx.WriteString(s)
}

func accountNewSchedulePost(ctx *web.Context) {
	session := app.SessionMgr.CurrentSession(ctx)
	if !session.GetBool("logged_in") {
		ctx.Redirect(301, "/login")
		return
	}

	//blah blah error

	// ctx.Redirect(301, "/account/new_schedule?selected_hour=" + ctx.Params["hour"] + "&err=Loli")
	m := map[string]interface{}{
		"Debug":   ctx.Params,
		"Message": ctx.Params["message"],
	}

	m["Hours"] = hoursList(ctx.Params["hour"])
	m["Minutes"] = minutesList(ctx.Params["minute"])

	if e, ok := ctx.Params["err"]; ok {
		m["Error"] = e
	}
	s := mustache.RenderFile("templ/account_newschedule.mustache", &m)
	ctx.WriteString(s)

}

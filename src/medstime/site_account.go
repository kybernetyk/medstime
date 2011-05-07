package main

import (
	"web"
	"mustache"
	"fmt"
	"strconv"
	"log"
)

func accountGet(ctx *web.Context) {
	session := app.SessionMgr.CurrentSession(ctx)
	if !session.GetBool("logged_in") {
		ctx.Redirect(301, "/account/login")
		return
	}

	ctx.Redirect(301, "/account/main")
}


func accountMainGet(ctx *web.Context) {
	session := app.SessionMgr.CurrentSession(ctx)

	log.Printf("session: %#v\n", session)

	if !session.GetBool("logged_in") {
		ctx.Redirect(301, "/account/login")
		return
	}

	accmgr := NewAccountManager()
	acc, _ := accmgr.AccountForAccountId(session.GetInt("account_id"))

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
	log.Printf("%#v - %d\n", items, len(items))

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
		ctx.Redirect(301, "/account/login")
		return
	}

	m := map[string]interface{}{
		"Debug":   ctx.Params,
		"Hours":   hoursList(""),
		"Minutes": minutesList(""),
	}

	s := mustache.RenderFile("templ/account_newschedule.mustache", &m)
	ctx.WriteString(s)
}

func accountNewSchedulePost(ctx *web.Context) {
	session := app.SessionMgr.CurrentSession(ctx)
	if !session.GetBool("logged_in") {
		ctx.Redirect(301, "/account/login")
		return
	}

	error := ""
	if len(ctx.Params["message"]) == 0 {
		error = "message no gief?"
		goto bailout
	}

	accmgr := NewAccountManager()
	acc, _ := accmgr.AccountForAccountId(session.GetInt("account_id"))

	ihr, _ := strconv.Atoi(ctx.Params["hour"])
	imn, _ := strconv.Atoi(ctx.Params["minute"])

	mgr := NewScheduleManager()
	scitms := mgr.ScheduleItemsForAccountAndOffset(acc, SecondsFromMidnight(ihr, imn))
	if len(scitms) > 0 {
		error = "a schedule for this time exists already!"
		goto bailout
	}

bailout:
	if len(error) > 0 {
		m := map[string]interface{}{
			"Debug":   ctx.Params,
			"Message": ctx.Params["message"],
			"Error":   error,
			"Hours":   hoursList(ctx.Params["hour"]),
			"Minutes": minutesList(ctx.Params["minute"]),
		}

		s := mustache.RenderFile("templ/account_newschedule.mustache", &m)
		ctx.WriteString(s)
		return
	}

	//get scheduled item count for items {accid, hour, minutes} -> if > 0 bail out
	sched := mgr.AddScheduleItemToAccount(acc)
	sched.OffsetFromMidnight = SecondsFromMidnight(ihr, imn)
	//sched.Hour = ihr
	//sched.Minute = imn
	sched.Message = ctx.Params["message"]
	mgr.UpdateScheduleItem(sched)
	ctx.Redirect(301, "/account")
}

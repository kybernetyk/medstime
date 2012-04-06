package main

import (
	"github.com/hoisie/web.go"
	"launchpad.net/mgo"
	"log"
)

type Application struct {
	SessionMgr *SessionManager
}

var app Application
var mgoSession *mgo.Session

func GetDB() (db *mgo.Database, session *mgo.Session) {
	session = mgoSession.Copy()
	tmp := session.DB("medstime")
	db = tmp
	return
}

func main() {
	log.Println("starting up ...")
	var err error

	log.Println("connecting to database ...")
	mgoSession, err = mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer mgoSession.Close()

	log.Println("creating session manager ...")
	app = Application{
		SessionMgr: NewSessionManager(),
	}

	log.Println("setting routes ...")
	web.Config.CookieSecret = "7C19QRmwf3mHZ9CPAaPQ0hsWeufKd"
	web.Get("/", indexGet)

	web.Get("/account/login", loginGet) //if logged in -> redirect, else show login form
	web.Post("/account/login", loginPost)

	web.Get("/account/logout", logoutGet)

	web.Get("/signup", signupGet)
	web.Post("/signup", signupPost)

	web.Get("/account", accountGet)
	web.Get("/account/main", accountMainGet)
	web.Get("/account/new_schedule", accountNewScheduleGet)
	web.Post("/account/new_schedule", accountNewSchedulePost)

	// web.Get("/post", post)
	// 
	// web.Get("/rss.xml", rss)
	// web.Get("/index.php/feed/", rss)
	// web.Get("/index.php/feed/atom/", rss)
	// 
	// 
	// web.Get("/admin/edit", editGet)
	// web.Post("/admin/edit", editPost)
	// 
	// web.Get("/admin", adminGet)
	// web.Post("/admin", adminPost)

	log.Println("and running!")
	web.Run("0.0.0.0:5555")

}

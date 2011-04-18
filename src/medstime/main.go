package main

import (
	"web"
	//	"fmt"
)

type Application struct {
    Db *MongoDB
    SessionMgr *SessionManager
}

var app Application

func main() {
	db := NewMongoDB()
	db.Connect()
    
    app = Application {
        Db: db,
        SessionMgr: NewSessionManager(),
    }
	
	web.Config.CookieSecret = "7C19QRmwf3mHZ9CPAaPQ0hsWeufKd"
	web.Get("/", indexGet)

	web.Get("/login", loginGet)
	web.Post("/login", loginPost)
	
	web.Get("/logout", logoutGet)

	web.Get("/signup", signupGet)
	web.Post("/signup", signupPost)


	web.Get("/account", accountGet)
	
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

	web.Run("0.0.0.0:5555")

}

package main

import (
	"launchpad.net/mgo"
	"os"
	"time"
	"log"
)

var control_chan chan string = make(chan string)

var mgoSession *mgo.Session
func GetDB() (db *mgo.Database, session *mgo.Session) {
	session = mgoSession.Copy()
	tmp := session.DB("medstime")
	db = &tmp
	return
}

func seconds(n int64) int64 {
	return 1000000000 * n
}

func minutes(n int64) int64 {
	return seconds(60 * n)
}

func main() {
	var err os.Error
	mgoSession, err = mgo.Mongo("127.0.0.1")
	if err != nil {
		log.Fatal(err.String())
		return
	}
	defer mgoSession.Close()

	log.Println("Ok, ready to go!")

	ticker := time.NewTicker(minutes(1))
L:
	for {
		select {
		case msg := <-control_chan:
			if msg == "quit" {
				break L
			}
		case <-ticker.C:
			go DoJob()
		}
	}
}

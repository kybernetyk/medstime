package main

import (
	"launchpad.net/mgo"
	"log"
	"time"
)

var control_chan chan string = make(chan string)

var mgoSession *mgo.Session

func GetDB() (db *mgo.Database, session *mgo.Session) {
	session = mgoSession.Copy()
	tmp := session.DB("medstime")
	db = tmp
	return
}

func seconds(n time.Duration) time.Duration {
	return time.Duration(1000000000 * n)
}

func minutes(n time.Duration) time.Duration {
	return seconds(60 * n)
}

func main() {
	var err error

	log.Println("Starting SMS demon.")

	log.Println("dialing mongo db ...")
	mgoSession, err = mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal(err.Error())
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

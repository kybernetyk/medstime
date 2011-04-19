package main

import "fmt"

//a schedule plan - a account has one
type Schedule struct {
	Id        int64
	AccountId int64 //the parent account
}

//an item in the schedule plan
type ScheduleItem struct {
	Id                 int64
	ScheduleId         int64  //the parent schedule
	Message            string //the message to be sent
	OffsetFromMidnight int64  //offset from midnight in seconds (to determine fire time)
}

//helper methods for mustache formatting - are called by mustche
func (itm ScheduleItem) Hour() string {
	h := itm.OffsetFromMidnight / (60 * 60)
	return fmt.Sprintf("%.2d", h)
}

func (itm ScheduleItem) Minute() string {
	m := (itm.OffsetFromMidnight % (60 * 60)) / 60
	return fmt.Sprintf("%.2d", m)
}

//get number of seconds from midnight for a given hour:minute pair
func SecondsFromMidnight(hour, minute int64) int64 {
	seconds := minute*60 + hour*(60*60)
	return seconds
}

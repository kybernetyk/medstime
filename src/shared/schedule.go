package main

import "fmt"

//a schedule plan - a account has one
type Schedule struct {
	Id        int
	AccountId int //the parent account
}

//an item in the schedule plan
type ScheduleItem struct {
	Id                 int
	ScheduleId         int  //the parent schedule
	Message            string //the message to be sent
	OffsetFromMidnight int  //offset from midnight in seconds (to determine fire time)
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

func HourMinute(offset int) (hour, minute int) {
	minute = (offset % (60 * 60)) / 60
	hour = offset / (60 * 60)
	return
}
//get number of seconds from midnight for a given hour:minute pair
func SecondsFromMidnight(hour, minute int) int {
	if hour < 0 || hour > 23 {
		hour = 0
	}
	if minute < 0 || minute > 59 {
		minute = 0
	}
	seconds := minute*60 + hour*(60*60)
	return seconds
}

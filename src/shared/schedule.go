package main

type ScheduleItem struct {
	Id         int64
	ScheduleId int64
	Message    string
	Hour       int
	Minute     int
}

type Schedule struct {
	Id        int64
	AccountId int64
}

package main

import (
	"launchpad.net/gobson/bson"
	// "launchpad.net/mgo"
)

type ScheduleManager struct{}

const (
	col_schedules      = "schedules"
	col_schedule_items = "scheduleitems"
)

func NewScheduleManager() *ScheduleManager {
	return &ScheduleManager{}
}

func (self *ScheduleManager) AddScheduleItemToAccount(account Account) ScheduleItem {
	schedule, ok := self.scheduleForAccountId(account.Id)
	if !ok {
		schedule = self.createScheduleForAccount(account)
	}

	si := ScheduleItem{
		ScheduleId: schedule.Id,
	}

	return self.addScheduleItem(si)
}

func (self *ScheduleManager) ScheduleItemsForAccount(account Account) []ScheduleItem {
	schedule, ok := self.scheduleForAccountId(account.Id)
	if !ok {
		return nil
	}

	items, ok := self.scheduleItemsForScheduleId(schedule.Id)
	if !ok {
		return nil
	}
	return items
}

func (self *ScheduleManager) UpdateScheduleItem(si ScheduleItem) {
	m := bson.M{"id": si.Id}

	db, ses := GetDB()
	defer ses.Close()

	db.C("scheduleitems").Update(m, si)
}

func (self *ScheduleManager) ScheduleItemsForOffset(offset int) []ScheduleItem {
	qry := bson.M{
		"$query": bson.M{"offsetfrommidnight": offset},
	}

	db, ses := GetDB()
	defer ses.Close()

	iter, err := db.C("scheduleitems").Find(qry).Iter()
	if err != nil {
		return nil
	}

	var items []ScheduleItem
	for {
		item := ScheduleItem{}
		err = iter.Next(&item)
		if err != nil {
			break
		}
		items = append(items, item)
	}
	return items
}

func (self *ScheduleManager) ScheduleItemsForAccountAndOffset(account Account, offset int) []ScheduleItem {
	schedule, ok := self.scheduleForAccountId(account.Id)
	if !ok {
		return nil
	}

	qry := bson.M{
		"$query": bson.M{"scheduleid": schedule.Id, "offsetfrommidnight": offset},
	}

	db, ses := GetDB()
	defer ses.Close()

	iter, err := db.C("scheduleitems").Find(qry).Iter()
	if err != nil {
		return nil
	}

	var items []ScheduleItem
	for {
		item := ScheduleItem{}
		err = iter.Next(&item)
		if err != nil {
			break
		}
		items = append(items, item)
	}
	return items
}


func (self *ScheduleManager) createScheduleForAccount(account Account) Schedule {
	sc := Schedule{
		AccountId: account.Id,
	}

	return self.createSchedule(sc)
}

func (self *ScheduleManager) createSchedule(schedule Schedule) Schedule {
	if _, ok := self.scheduleForAccountId(schedule.AccountId); ok {
		return schedule
	}

	db, ses := GetDB()
	defer ses.Close()

	count, _ := db.C("schedules").Count()
	count++
	schedule.Id = count
	db.C("schedules").Insert(schedule)
	return schedule
}


func (self *ScheduleManager) scheduleForAccountId(acc_id int) (schedule Schedule, ok bool) {
	qry := bson.M{
		"$query": bson.M{"accountid": acc_id},
	}

	db, ses := GetDB()
	defer ses.Close()

	err := db.C("schedules").Find(qry).One(&schedule)
	if err != nil {
		ok = false
		return
	}
	ok = true
	return
}


func (self *ScheduleManager) scheduleItemsForScheduleId(sched_id int) (items []ScheduleItem, ok bool) {
	qry := bson.M{
		"$query":   bson.M{"scheduleid": sched_id},
		"$orderby": bson.M{"offsetfrommidnight": 1},
	}

	db, ses := GetDB()
	defer ses.Close()

	iter, err := db.C("scheduleitems").Find(qry).Iter()
	if err != nil {
		ok = false
		return
	}

	for {
		item := ScheduleItem{}
		err = iter.Next(&item)
		if err != nil {
			break
		}
		items = append(items, item)
	}

	ok = true
	return
}


func (self *ScheduleManager) addScheduleItem(si ScheduleItem) ScheduleItem {
	db, ses := GetDB()
	defer ses.Close()

	count, _ := db.C("scheduleitems").Count()
	count++
	si.Id = count
	db.C("scheduleitems").Insert(si)
	return si
}

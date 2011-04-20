package main

import (
	"fmt"
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
	fmt.Printf("schedule id: %d\n", schedule)
	if !ok {
		return nil
	}

	items, ok := self.scheduleItemsForScheduleId(schedule.Id)
	fmt.Printf("items: %v\n", items)
	if !ok {
	    fmt.Printf("omg bai")
		return nil
	}
	return items
}

func (self *ScheduleManager) UpdateScheduleItem(si ScheduleItem) {
	m := bson.M{"id": si.Id}
	app.Db.C("scheduleitems").Update(m, si)
}

func (self *ScheduleManager) ScheduleItemsForAccountAndOffset(account Account, offset int) []ScheduleItem {
	schedule, ok := self.scheduleForAccountId(account.Id)
	if !ok {
		return nil
	}

	qry := bson.M{
		"$query": bson.M{"scheduleid": schedule.Id, "offsetfrommidnight": offset},
	}
	
	iter, err := app.Db.C("scheduleitems").Find(qry).Iter()
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
    fmt.Println(qry)
	return items
}


func (self *ScheduleManager) createScheduleForAccount(account Account) Schedule {
	sc := Schedule{
		AccountId: account.Id,
	}
	
	fmt.Printf("createScheduleForAccount %#v\n",account)

	return self.createSchedule(sc)
}

func (self *ScheduleManager) createSchedule(schedule Schedule) Schedule {
	if _, ok := self.scheduleForAccountId(schedule.AccountId); ok {
		return schedule
	}


	count, _ := app.Db.C("schedules").Count()
	count++
	schedule.Id = count
    app.Db.C("schedules").Insert(schedule)
	return schedule
}


func (self *ScheduleManager) scheduleForAccountId(acc_id int) (schedule Schedule, ok bool) {
	qry := bson.M{
		"$query": bson.M{"accountid": acc_id},
	}

    err := app.Db.C("schedules").Find(qry).One(&schedule)
    fmt.Printf("scheduleForAccountId: %d = %#v\n", acc_id , schedule)
    if err != nil {
        fmt.Println(err.String())
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
	
	iter, err := app.Db.C("scheduleitems").Find(qry).Iter()
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
	count, _ := app.Db.C("scheduleitems").Count()
	count++
	si.Id = count
    app.Db.C("scheduleitems").Insert(si)
	
	//app.Db.Insert(col_schedule_items, si)

	return si
}

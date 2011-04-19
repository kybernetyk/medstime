package main

import (
//	"os"
	"github.com/mikejs/gomongo/mongo"
)

type ScheduleManager struct{}

const (
    col_schedules = "schedules"
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
    m := querymap{"id": si.Id}
    app.Db.Update(col_schedule_items, si, m)
}






func (self *ScheduleManager) createScheduleForAccount(account Account) Schedule{
    sc := Schedule{
        AccountId: account.Id,
    }
    
    return self.createSchedule(sc)
}

func (self *ScheduleManager) createSchedule(schedule Schedule) Schedule {
    if _, ok := self.scheduleForAccountId(schedule.AccountId); ok {
        return schedule
    }
    
    qry := querymap{}
    count := app.Db.Count(col_schedules, qry)
    count++
    schedule.Id = count
    app.Db.Insert(col_schedules, schedule)

    return	schedule
}


func (self *ScheduleManager) scheduleForAccountId(acc_id int64) (schedule Schedule, ok bool) {
    	qry := querymap{
    		"$query": querymap{"accountid": acc_id},
    	}
    
    	docs, err := app.Db.Query(col_schedules, qry, 0, 1)
    	if err != nil || len(docs) == 0 {
    		ok = false
    		return
    	}

    	err = mongo.Unmarshal(docs[0].Bytes(), &schedule)
    	if err != nil {
    	    ok = false
    	    return
    	}

        ok = true
        return
}


func (self *ScheduleManager) scheduleItemsForScheduleId(sched_id int64) (items []ScheduleItem, ok bool) {
    	qry := querymap{
    		"$query": querymap{"scheduleid": sched_id},
    	}
    
    	docs, err := app.Db.Query(col_schedule_items, qry, 0, 0)
    	if err != nil || len(docs) == 0 {
    		ok = false
    		return
    	}

        for _, itembson := range docs {
            item := ScheduleItem{}
            mongo.Unmarshal(itembson.Bytes(), &item)
            items = append(items, item)
        }
        
        ok = true
        return
}


func (self *ScheduleManager) addScheduleItem(si ScheduleItem) ScheduleItem {
    qry := querymap{}
    count := app.Db.Count(col_schedule_items, qry)
    count++
    si.Id = count
    app.Db.Insert(col_schedule_items, si)
    
    return si
}


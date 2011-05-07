package main

import (
	"fmt"
	"time"
	"log"
	"github.com/jsz/gosms/sms"
)

func getScheduleItems(offset int) []ScheduleItem {
	mgr := NewScheduleManager()
	items := mgr.ScheduleItemsForOffset(offset)

	return items
}

func getSMSInterfaceForItem(item ScheduleItem) (iface SMSInterface, ok bool) {
	//query db for account blah blah

	iface = SMSInterface{
		Id:        1234,
		AccountId: 14414,
		TelNumber: "491737649349",
	}
	ok = true
	return
}

func sendSMS(number, message string) {
	log.Printf("\t\tsending '%s' to '%s' ...\n", message, number)

	s := sms.NewBulkSMSSMSSender("joorek", "warbird")
	//s.Testmode = 1     //don't send the sms, just perform an API supported test
	s.RoutingGroup = 2 //let's use the cheap eco route

	receivers := []string{number} //put a proper tel# here in

	//let's see how much this sms would cost us
	_, quote := s.GetQuote(receivers, message)
	price := quote * 3.75 * 0.01 //quote is in credits. 1 credit = 3.75 eur cent

	log.Printf("\t\t\tPrice for SMS(%s): %.4f EUR\n", number, price)

	//send the sms
	if err := s.Send(receivers, message); err != nil {
		log.Printf("\t\tcouldn't send sms to '%s': %s\n",number, err.String())
		return
	}

	fmt.Printf("\t\tsms sent to '%s'!\n", number)
}

func DoJob() {
	now := time.LocalTime()

	hour := now.Hour
	minute := now.Minute

	offset := SecondsFromMidnight(hour, minute)
	log.Printf("doing job on %.2d:%.2d -> %d ...\n", hour, minute, offset)

	items := getScheduleItems(offset)
	if items == nil {
		log.Printf("\tno items to send ...\n")
		return
	}

	for _, item := range items {
		iface, ok := getSMSInterfaceForItem(item)
		if !ok {
			log.Printf("\tNo SMS interface found for item %#v!\n", item)
			continue
		}
		go sendSMS(iface.TelNumber, item.Message)
	}

}

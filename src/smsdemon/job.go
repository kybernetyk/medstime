package main

import (
	"fmt"
	"time"
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
		TelNumber: "48516520678",
	}
	ok = true
	return
}

func sendSMS(number, message string) {
	fmt.Printf("sending '%s' to '%s' ...\n", message, number)

	s := sms.NewBulkSMSSMSSender("joorek", "warbird")
	s.Testmode = 1     //don't send the sms, just perform an API supported test
	s.RoutingGroup = 2 //let's use the cheap eco route

	receivers := []string{number} //put a proper tel# here in

	//let's see how much this sms would cost us
	_, quote := s.GetQuote(receivers, message)
	price := quote * 3.75 * 0.01 //quote is in credits. 1 credit = 3.75 eur cent

	fmt.Printf("\tPrice for SMS(%s): %.4f EUR\n", number, price)

	//send the sms
	if err := s.Send(receivers, message); err != nil {
		fmt.Printf("\tcouldn't send sms to '%s': %s\n",number, err.String())
		return
	}

	fmt.Printf("\tsms sent to '%s'!\n", number)
}

func DoJob() {
	now := time.LocalTime()

	hour := now.Hour
	minute := now.Minute

	hour = 13
	minute = 30
	offset := SecondsFromMidnight(hour, minute)

	fmt.Printf("doing job on %.2d:%.2d -> %d ...\n", hour, minute, offset)

	items := getScheduleItems(offset)
	if items == nil {
		return
	}

	for _, item := range items {
		iface, ok := getSMSInterfaceForItem(item)
		if !ok {
			fmt.Printf("No SMS interface found for item %#v!\n", item)
			continue
		}
		go sendSMS(iface.TelNumber, item.Message)
	}

}

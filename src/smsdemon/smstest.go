package main

import (
	"github.com/jsz/gosms/sms"
	"fmt"
)

func main() {
	s := sms.NewBulkSMSSMSSender("joorek", "warbird")
	s.Testmode = 1     //don't send the sms, just perform an API supported test
	s.RoutingGroup = 2 //let's use the cheap eco route

	msg := "tabletten nehmen!"
	receivers := []string{"491737649349"} //put a proper tel# here in

	//let's see how much this sms would cost us
	_, quote := s.GetQuote(receivers, msg)
	price := quote * 3.75 * 0.01 //quote is in credits. 1 credit = 3.75 eur cent

	fmt.Printf("the sms will cost us %.4f EUR\n", price)

	//send the sms
	if err := s.Send(receivers, msg); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("sms sent!")
}

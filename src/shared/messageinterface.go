package main

type SMSInterface struct {
	Id	int
	AccountId	int

	TelNumber	string
}

type MailInterface struct {
	Id int
	AccountId int

	Address string
}

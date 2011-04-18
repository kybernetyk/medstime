package main

type SMSInterface struct {
	Id	int64
	AccountId	int64

	TelNumber	string
}

type MailInterface struct {
	Id int64
	AccountId int64

	Address string
}

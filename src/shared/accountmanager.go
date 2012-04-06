package main

import (
	"errors"
	"launchpad.net/mgo/bson"
)

type AccountManager struct{}

func NewAccountManager() *AccountManager {
	return new(AccountManager)
}

func (self *AccountManager) AccountForEmailPassword(email, password string) (account Account, ok bool) {
	account, ok = self.AccountForEmail(email)
	if !ok {
		return
	}
	if account.Email == email && account.Password == password {
		ok = true
		return
	}

	ok = false
	return
}

func (self *AccountManager) AccountForEmail(email string) (account Account, ok bool) {
	qry := bson.M{
		"$query": bson.M{"email": email},
	}

	db, ses := GetDB()
	defer ses.Close()

	err := db.C("accounts").Find(qry).One(&account)
	if err != nil {
		ok = false
		return
	}

	ok = true
	return
}

func (self *AccountManager) AccountForAccountId(acc_id int) (account Account, ok bool) {
	qry := bson.M{
		"$query": bson.M{"id": acc_id},
	}

	db, ses := GetDB()
	defer ses.Close()

	err := db.C("accounts").Find(qry).One(&account)
	if err != nil {
		ok = false
		return
	}

	ok = true
	return
}

func (self *AccountManager) UpdateAccount(account Account) int {
	//m := bson.M{"id": account.Id}
	qry := bson.M{
		"id": account.Id,
	}

	db, ses := GetDB()
	defer ses.Close()

	err := db.C("accounts").Update(qry, &account)
	if err != nil {
		return 0
	}
	return account.Id
}

func (self *AccountManager) CreateAccount(account Account) (acc_id int, err error) {
	if _, ok := self.AccountForEmail(account.Email); ok {
		err = errors.New(err_SignupEmailExists)
		return
	}

	if _, ok := self.AccountForAccountId(account.Id); ok {
		err = errors.New(err_SignupEmailExists)
		return
	}

	db, ses := GetDB()
	defer ses.Close()

	count, _ := db.C("accounts").Count()
	count++
	account.Id = count

	err = db.C("accounts").Insert(&account)
	if err != nil {
		err = errors.New(err_Critical)
		return
	}
	acc_id = account.Id
	return
}

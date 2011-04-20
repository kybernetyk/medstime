package main

import (
	"os"
	"launchpad.net/gobson/bson"
    "fmt"
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
    fmt.Printf("%#v\n", account)
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

    err := app.Db.C("accounts").Find(qry).One(&account)
    if err != nil {
        fmt.Println("AccountForEmail:" + err.String());
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

    err := app.Db.C("accounts").Find(qry).One(&account)
    if err != nil {
        fmt.Println("AccountForAccountId: " + err.String());
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
    err := app.Db.C("accounts").Update(qry, &account)
	if  err != nil {
		return 0
	}
	return account.Id
}

func (self *AccountManager) CreateAccount(account Account) (acc_id int, err os.Error) {
	if _, ok := self.AccountForEmail(account.Email); ok {
		err = os.NewError(err_SignupEmailExists)
		return
	}

	if _, ok := self.AccountForAccountId(account.Id); ok {
		err = os.NewError(err_SignupEmailExists)
		return
	}

	count, _ := app.Db.C("accounts").Count()
	count++
	account.Id = count

    err = app.Db.C("accounts").Insert(&account)
	if err != nil {
		err = os.NewError(err_Critical)
		return
	}
	acc_id = account.Id
	return
}

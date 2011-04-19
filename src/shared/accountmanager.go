package main

import (
	"os"
	"github.com/mikejs/gomongo/mongo"
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
	qry := querymap{
		"$query": querymap{"email": email},
	}

//	var docs []mongo.BSON
	docs, err := app.Db.Query(col_accounts, qry, 0, 1)
	if err != nil || len(docs) == 0 {
		ok = false
		return
	}

	err = mongo.Unmarshal(docs[0].Bytes(), &account)
	if err != nil {
	    ok = false
	    return
	}

	ok = true
	return
}

func (self *AccountManager) AccountForAccountId(acc_id int64) (account Account, ok bool) {
	qry := querymap{
		"$query": querymap{"id": acc_id},
	}

//	var docs []mongo.BSON
	docs, err := app.Db.Query(col_accounts, qry, 0, 1)
	if err != nil || len(docs) == 0 {
		ok = false
		return
	}

	err = mongo.Unmarshal(docs[0].Bytes(), &account)
	if err != nil {
	    ok = false
	    return
	}

	ok = true
	return
}

func (self *AccountManager) UpdateAccount(account Account) int64 {
    m := querymap{"id": account.Id}
    ok := app.Db.Update(col_accounts, account, m)
    if !ok {
        return 0
    }
    return account.Id
}

func (self *AccountManager) CreateAccount(account Account) (acc_id int64, err os.Error) {
	if _, ok := self.AccountForEmail(account.Email); ok {
		err = os.NewError(err_SignupEmailExists)
		return
	}

	if _, ok := self.AccountForAccountId(account.Id); ok {
		err = os.NewError(err_SignupEmailExists)
		return
	}

    qry := querymap{}
    count := app.Db.Count(col_accounts, qry)
    count++
    account.Id = count
    
    if ok := app.Db.Insert(col_accounts, account); !ok {
        err = os.NewError(err_Critical)
        return
    }
    acc_id = account.Id
    return
}

package main
import (
    "os"
)

type AccountManager struct {
    
}

func NewAccountManager() *AccountManager {
    acc := new(AccountManager)
    return acc
}

func (self *AccountManager) AccountForUsernamePassword(username, password string) (account Account, ok bool) {
    var err os.Error
    account, err = app.Db.GetAccountForUsername(username)
    if err != nil {
        ok = false
        return
    }
    
    if account.Username == username && account.Password == password {
        ok = true
        return
    }
    
    ok = false
    return
}

func (self *AccountManager) AccountForAccountId(acc_id int64) (account Account, ok bool) {
    var err os.Error
    account, err = app.Db.GetAccountForAccountId(acc_id)
    if err != nil {
        ok = false
        return
    }
    
    ok = true
    return
}

func (self *AccountManager) StoreAccount(account Account) int64 {
    id, _ := app.Db.StoreAccount(account)
    return id
}

func (self *AccountManager) CreateAccount(account Account) (acc_id int64, err os.Error) {
   _, ok := self.AccountForUsernamePassword(account.Username, account.Password)
   if ok {
       err = os.NewError(err_SignupUsernameExists)
       return
   }

   _, ok = self.AccountForAccountId(account.Id)
   if ok {
       err = os.NewError(err_SignupUsernameExists)
       return
   }

   acc_id = self.StoreAccount(account)

   
   return
}
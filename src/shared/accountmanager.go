package main

type AccountManager struct {
    
}

func NewAccountManager() *AccountManager {
    acc := new(AccountManager)
    return acc
}

func (self *AccountManager) AccountForUsernamePassword(username, password string) (account Account, ok bool) {
    ok = true
    account.Id = 456
    return
}
package main

import (
    "web"
)

const (
    err_LoginNoUsername = "1"
    err_LoginNoPass = "2"
    err_LoginFailed = "3"
    
    err_SignupUsernameExists = "11"
    err_SignupUsernameInvalid = "12"

    err_SignupEmailExists = "13"
    err_SignupEmailInvalid = "14"
    
    err_SignupPasswordInvalid = "15"
    
    err_Critical = "99"
)

var errmap = map[string]string {
    err_LoginNoUsername: "No Username",
    err_LoginNoPass: "No Password",
    err_LoginFailed: "Login Wrong",

    err_SignupUsernameExists: "This Username Exists already!",
    err_SignupUsernameInvalid: "Invalid Username",
    err_SignupEmailExists: "Another Account with this eMail exists already!",
    err_SignupEmailInvalid: "Not a valid email Address!",
    err_SignupPasswordInvalid: "Not a valid Password!",

    err_Critical: "Critical Error! O M G!",
}

func GetErrorString(ctx *web.Context) (estr string, ok bool) {
    err, found := ctx.Params["err"]
    if found {
        estr, ok = errmap[err]
        return
    }
    ok = false
    return
}
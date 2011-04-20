package main

import (
)

type Session struct { 
	Id string

	LastActive   int64 //unix timestamp
	TimeoutAfter int64 //seconds

	Data map[string]interface{}
}

//generics would be cool
func (self *Session) Set(key string, val interface{}) {
    if val == nil {
        self.Data[key] = nil, false
        return
    }
    self.Data[key] = val
}

func (self *Session) Get(key string) interface{} {
    v, ok := self.Data[key]
    if !ok {
        return nil
    }
    return v
}

func (self *Session) GetBool(key string) bool {
    v, ok := self.Data[key]
    if !ok {
       return false
    }
    return v.(bool)
}

func (self *Session) GetInt(key string) int {
    v, ok := self.Data[key]
    if !ok {
       return 0
    }
    return v.(int)
}


func (self *Session) GetString(key string) string {
    v, ok := self.Data[key]
    if !ok {
        return ""
    }
    return v.(string)
}
package KiiLib_Go

import (
	rj "github.com/fkmhrk-go/rawjson"
)

type User struct {
	ID   string
	Data rj.RawJsonObject
}

type UserAPI interface {
}

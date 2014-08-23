package KiiLib_Go

import (
	rj "github.com/fkmhrk-go/rawjson"
)

type AppAPI interface {
	Login(identifier, password string) (*User, string, error)
	LoginAsAdmin(clientId, clientSecret string) (*User, string, error)
	SignUp(info rj.RawJsonObject, password string) (*User, error)

	//	UserAPI() UserAPI
}

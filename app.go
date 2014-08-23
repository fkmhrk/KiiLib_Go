package KiiLib_Go

import (
	rj "github.com/fkmhrk-go/rawjson"
)

type AppAPI interface {
	Login(identifier, password string) *User
	LoginAsAdmin(clientId, clientSecret string) *User

	SignUp(info rj.RawJsonObject, password string) *User

	UserAPI() UserAPI
}

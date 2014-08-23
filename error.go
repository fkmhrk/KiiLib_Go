package KiiLib_Go

import (
	"fmt"
	rj "github.com/fkmhrk-go/rawjson"
)

type KiiError struct {
	Status int
	Body   rj.RawJsonObject
}

func (e KiiError) Error() string {
	code, _ := e.Body.String("errorCode")
	return fmt.Sprintf("HTTP %d CODE=%s", e.Status, code)
}

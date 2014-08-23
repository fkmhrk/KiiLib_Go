package KiiLib_Go

import (
	rj "github.com/fkmhrk-go/rawjson"
)

type HttpClientFactory interface {
	NewClient() HttpClient
}

type HttpMethod int

const (
	HTTP_GET HttpMethod = iota
	HTTP_POST
	HTTP_PUT
	HTTP_DELETE
)

type HttpHeader map[string]string

type HttpClient interface {
	SendJsonRequest(method HttpMethod,
		url string,
		header HttpHeader,
		body rj.RawJsonObject) (*HttpResponse, error)
}

type HttpResponse struct {
	Status int
	Body   rj.RawJsonObject
	Etag   string
}

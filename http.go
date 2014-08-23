package KiiLib_Go

import (
	rj "github.com/fkmhrk-go/rawjson"
)

type HttpClientFactory interface {
	NewClient() HttpClient
}

type HttpMethod string

const (
	HTTP_GET    HttpMethod = "GET"
	HTTP_POST   HttpMethod = "POST"
	HTTP_PUT    HttpMethod = "PUT"
	HTTP_DELETE HttpMethod = "DELETE"
)

type HttpHeader map[string]string

type HttpClient interface {
	SendJsonRequest(methdo HttpMethod,
		url string,
		header HttpHeader,
		body rj.RawJsonObject) (*HttpResponse, error)
}

type HttpResponse struct {
	status int
	body   rj.RawJsonObject
	etag   string
}

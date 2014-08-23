package KiiLib_Go

import (
	"bytes"
	"encoding/json"
	"fmt"
	rj "github.com/fkmhrk-go/rawjson"
	"net/http"
)

type KiiHttpClientFactory struct {
}

func (f *KiiHttpClientFactory) NewClient() HttpClient {
	return &kiiHttpClient{}
}

// utility
func (m HttpMethod) String() string {
	switch m {
	case HTTP_GET:
		return "GET"
	case HTTP_POST:
		return "POST"
	case HTTP_PUT:
		return "PUT"
	case HTTP_DELETE:
		return "DELETE"
	}
	return ""
}

type kiiHttpClient struct {
}

func (c *kiiHttpClient) SendJsonRequest(method HttpMethod,
	url string,
	header HttpHeader,
	body rj.RawJsonObject) (*HttpResponse, error) {
	// to byte array
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("failed to encode json")
		return nil, err
	}
	// create reader
	r := bytes.NewReader(b)
	req, err := http.NewRequest(method.String(), url, r)
	if err != nil {
		fmt.Printf("failed to create request")
		return nil, err
	}
	// set header
	for key, value := range header {
		req.Header.Add(key, value)
	}
	//	req.Header.Add("content-type", "application/json")
	//	req.Header.Add("x-kii-appid", "<your App Id>")
	//	req.Header.Add("x-kii-appkey", "<your App Key>")
	// create client
	client := &http.Client{}
	// send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to execute")
		return nil, err
	}
	// to Json
	var respJson rj.RawJsonObject
	if resp.StatusCode == 204 {
		respJson, _ = rj.ObjectFromString("{}")
	} else {
		respJson, err = rj.ObjectFromReader(resp.Body)
		if err != nil {
			return nil, err
		}
	}
	etag := resp.Header.Get("ETag")

	return &HttpResponse{
		Status: resp.StatusCode,
		Body:   respJson,
		Etag:   etag,
	}, nil
}

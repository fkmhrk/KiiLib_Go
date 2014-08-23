package KiiLib_Go

import (
	"testing"
)

func TestHttpImpl_0000_POSTRequest(t *testing.T) {
	f := &KiiHttpClientFactory{}
	c := f.NewClient()

	url := "https://api-jp.kii.com/api/oauth2/token"
	headers := HttpHeader{
		"content-type": "application/json",
		"x-kii-appid":  "<your app ID>",
		"x-kii-appkey": "<your app Key>",
	}
	body := map[string]interface{}{
		"username": "fkm",
		"password": "123456",
	}
	resp, err := c.SendJsonRequest(HTTP_POST, url, headers, body)
	if err != nil {
		t.Errorf("error : %s", err)
	}
	if resp.Status == 200 || resp.Status == 400 || resp.Status == 404 {
		// OK (maybe...)
	} else {
		t.Errorf("unexpected status : %d", resp.Status)
	}
	//for k, v := range resp.Body {
	//	t.Errorf("body : %s => %s ", k, v)
	//}
}

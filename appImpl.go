package KiiLib_Go

import (
	rj "github.com/fkmhrk-go/rawjson"
)

type kiiAppAPI struct {
	factory HttpClientFactory
	appId   string
	appKey  string
	baseURL string
}

func NewAppAPI(appId, appKey, baseURL string,
	factory HttpClientFactory) *kiiAppAPI {
	return &kiiAppAPI{
		appId:   appId,
		appKey:  appKey,
		baseURL: baseURL,
		factory: factory,
	}
}

func (a *kiiAppAPI) Login(identifier, password string) (*User, string, error) {
	client := a.factory.NewClient()
	url := a.baseURL + "/oauth2/token"
	headers := HttpHeader{
		"content-type": "application/json",
		"x-kii-appid":  a.appId,
		"x-kii-appkey": a.appKey,
	}
	body := map[string]interface{}{
		"username": identifier,
		"password": password,
	}
	resp, err := client.SendJsonRequest(HTTP_POST, url, headers, body)
	return handleLoginResponse(resp, err)
}

func (a *kiiAppAPI) LoginAsAdmin(clientId, clientSecret string) (*User, string, error) {
	client := a.factory.NewClient()
	url := a.baseURL + "/oauth2/token"
	headers := HttpHeader{
		"content-type": "application/json",
		"x-kii-appid":  a.appId,
		"x-kii-appkey": a.appKey,
	}
	body := map[string]interface{}{
		"client_id":     clientId,
		"client_secret": clientSecret,
	}
	resp, err := client.SendJsonRequest(HTTP_POST, url, headers, body)
	return handleLoginResponse(resp, err)
}

func handleLoginResponse(resp *HttpResponse, err error) (*User, string, error) {
	if err != nil {
		return nil, "", err
	}
	if resp.Status != 200 {
		return nil, "", KiiError{
			Status: resp.Status,
			Body:   resp.Body,
		}
	}
	// get userId and access token
	id, _ := resp.Body.String("id")
	token, _ := resp.Body.String("access_token")

	return &User{
		ID: id,
	}, token, nil
}

func (a *kiiAppAPI) SignUp(info rj.RawJsonObject, password string) (*User, error) {
	client := a.factory.NewClient()
	url := a.baseURL + "/apps/" + a.appId + "/users"
	headers := HttpHeader{
		"content-type": "application/json",
		"x-kii-appid":  a.appId,
		"x-kii-appkey": a.appKey,
	}
	body := map[string]interface{}{
		"password": password,
	}
	for key, value := range info {
		body[key] = value
	}
	resp, err := client.SendJsonRequest(HTTP_POST, url, headers, body)
	if err != nil {
		return nil, err
	}
	if resp.Status != 201 {
		return nil, KiiError{
			Status: resp.Status,
			Body:   resp.Body,
		}
	}
	// get userId
	id, _ := resp.Body.String("userID")
	return &User{
		ID:   id,
		Data: info,
	}, nil
}

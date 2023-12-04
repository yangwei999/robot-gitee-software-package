package useradapterimpl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/opensourceways/server-common-lib/utils"
)

type omTokenReq struct {
	GrantType string `json:"grant_type"`
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type omTokenResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
}

type omUserInfoResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	Identities []Identities `json:"identities"`
	Username   string       `json:"username"`
	Email      string       `json:"email"`
}

type Identities struct {
	LoginName string `json:"login_name"`
	Identity  string `json:"identity"`
}

type omClient struct {
	cli utils.HttpClient
	cfg *Config
}

func NewAdapterImpl(c *Config) *omClient {
	return &omClient{
		cli: utils.NewHttpClient(3),
		cfg: c,
	}
}

func (o *omClient) GetEmail(username string) (string, error) {
	user, err := o.getUserInfo(username)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (o *omClient) getToken() (string, error) {
	request := omTokenReq{
		GrantType: "token",
		AppId:     o.cfg.AppId,
		AppSecret: o.cfg.AppSecret,
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, o.cfg.TokenEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	v := new(omTokenResp)
	if _, err = o.cli.ForwardTo(req, v); err != nil {
		return "", err
	}
	if v.Status != 200 {
		return "", errors.New(v.Msg)
	}

	return v.Token, nil
}

func (o *omClient) getUserInfo(username string) (d Data, err error) {
	token, err := o.getToken()
	if err != nil {
		return
	}

	url := o.cfg.UserEndpoint + username
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Set("token", token)

	v := new(omUserInfoResp)
	if _, err = o.cli.ForwardTo(req, v); err != nil {

		return
	}

	if v.Code != 200 {
		err = errors.New(v.Msg)

		return
	}

	return v.Data, nil
}

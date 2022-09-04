package app

import (
	"errors"
	json "github.com/json-iterator/go"
)

func (t *AppCore) loginaCode(body []byte) (*loginARequest, error) {
	var req loginARequest
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Data.Msg)
	}
	if req.Data.Status != 1 {
		return nil, errors.New(req.Data.Msg)
	}
	return &req, nil
}

type loginARequest struct {
	Code int `json:"code"`
	Data struct {
		CookieInfo struct {
			AccountId   int    `json:"account_id"`
			CookieToken string `json:"cookie_token"`
			CreateTime  int    `json:"create_time"`
			CurTime     int    `json:"cur_time"`
			Email       string `json:"email"`
			IsAdult     int    `json:"is_adult"`
			IsRealname  int    `json:"is_realname"`
			Mobile      string `json:"mobile"`
		} `json:"cookie_info"`
		Msg    string `json:"msg"`
		Sign   string `json:"sign"`
		Status int    `json:"status"`
	} `json:"data"`
}

func (t *AppCore) loginbCode(body []byte) (*loginBRequest, error) {
	var req loginBRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, err
	}
	if req.Retcode != 0 {
		return nil, errors.New(req.Message)
	}
	return &req, nil
}

type loginBRequest struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Name  string `json:"name"`
			Token string `json:"token"`
		} `json:"list"`
	} `json:"data"`
}

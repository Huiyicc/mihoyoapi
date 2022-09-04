package app

import (
	"errors"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
	"io"
	"net/http"
	"strings"
)

// httpGet 用于内部请求
func (t *AppCore) httpGet(req request.RequestStruct, dsType int, headerFunc func(*http.Header)) ([]byte, error) {
	uri := req.Url
	if req.Query != "" {
		uri += "?" + req.Query
	}
	var (
		resp *http.Response
		err  error
	)
	cli := &http.Client{}
	requ, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	requ.Header = t.Cookies.GetHeadersMap(req, dsType)
	if headerFunc != nil {
		headerFunc(&requ.Header)
	}
	//处理返回结果
	resp, _ = cli.Do(requ)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return data, nil
}

// httpPost 用于内部请求
func (t *AppCore) httpPost(req request.RequestStruct, dsType int, headerFunc func(*http.Request)) ([]byte, error) {
	uri := req.Url
	if req.Query != "" {
		uri += "?" + req.Query
	}
	var (
		resp *http.Response
		err  error
		cli  http.Client
	)

	/*	ProxyURL, _ := url.Parse("http://127.0.0.1:8080")
		cli = http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(ProxyURL),
			},
		}*/

	requ, err := http.NewRequest("POST", uri, strings.NewReader(req.Body.GetData()))
	if err != nil {
		return nil, err
	}
	requ.Header = t.Cookies.GetHeadersMap(req, dsType)
	if headerFunc != nil {
		headerFunc(requ)
	}
	//处理返回结果
	resp, _ = cli.Do(requ)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return data, nil
}

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

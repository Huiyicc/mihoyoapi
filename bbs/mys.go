package bbs

import (
	"errors"
	"github.com/Huiyicc/mihoyoapi/Cookies"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	"io"
	"net/http"
	"strings"
	"time"
)

type AppResponse struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	ERROR_LOGIN_FAILURE        = errors.New("登录失效")
	ERROR_NO_ROLE              = errors.New("请先去米游社绑定角色")
	ERROR_UPPER_LIMIT          = errors.New("查询已达今日上限")
	ERROR_USER_IS_PRIVATE      = errors.New("此账号数据未公开")
	ERROR_INTERFACE_ERROR      = errors.New("接口错误")
	ERROR_INTERFACE_ERROR_1002 = errors.New("接口错误:-1002")
)

type BBSCore struct {
}

func (t *BBSCore) Init(server string) error {
	if server == define.SERVER_TIANKONGDAO || server == define.SERVER_SHIJIESHU {
		return nil
	}
	return errors.New("不支持的平台")
}

// checkTime 数据更新中，请稍后再试
func (t *BBSCore) checkTime() error {
	now := time.Now()
	hour := now.Hour()
	min := now.Minute()
	second := now.Second()
	if hour == 23 && min == 59 && second >= 58 {
		return errors.New("数据更新中，请稍后再试")
	}
	if hour == 0 && min == 0 && second <= 3 {
		return errors.New("数据更新中，请稍后再试")
	}
	return nil
}

func (t *BBSCore) GetUrl(Type string) (request.RequestStruct, error) {
	reqs, ifSet := request.UrlMap[Type]
	if !ifSet {
		return request.RequestStruct{}, errors.New("请求方案不存在")
	}
	cmp := request.RequestStruct{
		Url:   reqs.Url,
		Query: reqs.Query,
		Sign:  reqs.Sign,
		Body:  reqs.Body,
	}
	return cmp, nil
}

func (t *BBSCore) Get(r request.RequestStruct, c *Cookies.CookiesCore) ([]byte, error) {
	url := r.Url + "?" + r.Query
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.GetHeadersMap(r)
	if req.Header.Get("Cookie") == "" {
		req.Header.Set("Cookie", c.GetCookies())
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (t *BBSCore) Post(r request.RequestStruct, c *Cookies.CookiesCore) ([]byte, error) {
	url := r.Url
	req, err := http.NewRequest("POST", url, strings.NewReader(r.Body.GetData()))
	if err != nil {
		return nil, err
	}
	req.Header = c.GetHeadersMap(r)
	if req.Header.Get("Cookie") == "" {
		req.Header.Set("Cookie", c.GetCookies())
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (t *BBSCore) SwitchCode(code int, msg string) error {
	switch code {
	case 0:
		return nil
	case -1:
	case -100:
	case 1001:
	case 10001:
	case 10103:
		return ERROR_LOGIN_FAILURE
	case 1008:
		return ERROR_NO_ROLE
	case 10101:
		return ERROR_UPPER_LIMIT
	case 10102:
		if msg == "Data is not public for the user" {
			return ERROR_USER_IS_PRIVATE
		}
		return ERROR_NO_ROLE
	case -1002:
		return ERROR_INTERFACE_ERROR_1002
	}
	return ERROR_INTERFACE_ERROR
}

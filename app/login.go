package app

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mihoyoapi/Cookies"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	"strconv"
)

// Login 用于登陆米游社app,返回登陆成功后cookies
func (t *AppCore) Login() (string, error) {
	if err := t.login1(); err != nil {
		return "", err
	}
	if err := t.login2(); err != nil {
		return "", err
	}
	return t.Cookies.GetCookies(), nil
}
func (t *AppCore) login1() error {
	requ := request.UrlMap[define.MIHOYOAPP_API_LOGINA]
	requ.Query = fmt.Sprintf(requ.Query, t.Cookies.LoginTicket)
	data, err := t.httpGet(requ, 1, nil)
	if err != nil {
		return err
	}
	res, err := t.loginaCode(data)
	if err != nil {
		return err
	}
	t.Cookies.Stuid = strconv.Itoa(res.Data.CookieInfo.AccountId)
	t.Cookies.CookiesMap["stuid"] = t.Cookies.Stuid
	return nil
}
func (t *AppCore) login2() error {
	requ := request.UrlMap[define.MIHOYOAPP_API_LOGINB]
	requ.Query = fmt.Sprintf(requ.Query, t.Cookies.LoginTicket, t.Cookies.Stuid)
	data, err := t.httpGet(requ, 1, nil)
	if err != nil {
		return err
	}
	res, err := t.loginbCode(data)
	if err != nil {
		return err
	}
	dlLen := len(res.Data.List)
	if dlLen == 0 {
		return errors.New("api错误")
	}
	for i := 0; i < dlLen; i++ {
		t.Cookies.CookiesMap[res.Data.List[i].Name] = res.Data.List[i].Token
	}
	t.Cookies.Stoken = t.Cookies.CookiesMap["stoken"]
	t.Cookies.Stuid = t.Cookies.CookiesMap["stuid"]
	return nil
}

// LoginToCookiesStr 用于使用cookies文本登陆,需先调用Login方法换取token
func (t *AppCore) LoginToCookiesStr(cookies string) error {
	var (
		err error
		c   Cookies.AppCookies
	)
	//解析token
	if err = c.ParseForLoginApp(cookies); err != nil {
		return err
	}
	t.Cookies = &c
	t.TasksInfo, err = t.GetTasksList()
	return err
}

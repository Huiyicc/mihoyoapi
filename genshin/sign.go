package genshin

import (
	"errors"
	"fmt"
	app2 "github.com/Huiyicc/mihoyoapi/bbs"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
)

type SignCore struct {
	app     app2.BBSCore
	cookies *request.CookiesCore
}

func (t *SignCore) Init(c *request.CookiesCore) error {
	if err := t.app.Init(c.Region); err != nil {
		return err
	}
	t.cookies = &request.CookiesCore{}
	err := t.cookies.Parse(c.GetCookies())
	return err
}

// Info 用于获取签到信息
func (t *SignCore) Info() (*SignInfo, error) {
	req, _ := t.app.GetUrl(define.MYSINFO_API_BBSSIGN_INFO)
	req.Query = fmt.Sprintf(req.Query, t.cookies.Region, t.cookies.GameUID)
	data, err := t.app.Get(req, t.cookies)
	if err != nil {
		return nil, err
	}
	var rep signInfoResponse
	if err = json.Unmarshal(data, &rep); err != nil {
		return nil, err
	}
	if rep.Retcode == -100 {
		return nil, errors.New("签到失败,cookies失效")
	}
	if rep.Retcode != 0 {
		return nil, errors.New("签到失败,未知错误 " + rep.Message)
	}
	rtmp := rep.Data
	return &rtmp, nil
}

// GenShiSign 用于原神签到,签到前请使用info判断是否签到,不然重复签到会被风控
func (t *SignCore) GenShiSign() error {
	var (
		err  error
		data []byte
	)
	req, _ := t.app.GetUrl(define.MYSINFO_API_BBSSIGN)
	req.Query = fmt.Sprintf(req.Query, t.cookies.Region, t.cookies.GameUID)
	req.Body["region"] = t.cookies.Region
	req.Body["uid"] = t.cookies.GameUID
	if data, err = t.app.Post(req, t.cookies); err != nil {
		return err
	}
	var resp genShiSignInfoResponse
	json.Unmarshal(data, &req)
	if resp.Data.Success != 0 {
		return errors.New("验证码失败")
	}
	if resp.Retcode == -5003 {
		return errors.New("已签到,注意：多次重复签到可能会触发账号风控")
	}
	if resp.Retcode == 0 && resp.Data.Success == 0 {
		return nil
	}
	return errors.New("签到失败,未知原因")
}

type signInfoResponse struct {
	Retcode int      `json:"retcode"`
	Message string   `json:"message"`
	Data    SignInfo `json:"data"`
}

type SignInfo struct {
	TotalSignDay  int    `json:"total_sign_day"`  //总签到天数
	Today         string `json:"today"`           //今天
	IsSign        bool   `json:"is_sign"`         //是否已经签到
	FirstBind     bool   `json:"first_bind"`      //是否第一次绑定
	IsSub         bool   `json:"is_sub"`          //
	MonthFirst    bool   `json:"month_first"`     //第一个月
	SignCntMissed int    `json:"sign_cnt_missed"` //漏签天数
}

type genShiSignInfoResponse struct {
	Retcode int            `json:"retcode"`
	Message string         `json:"message"`
	Data    GenShiSignInfo `json:"data"`
}
type GenShiSignInfo struct {
	Code      string `json:"code"`
	RiskCode  int    `json:"risk_code"`
	Gt        string `json:"gt"`
	Challenge string `json:"challenge"`
	Success   int    `json:"success"`
}

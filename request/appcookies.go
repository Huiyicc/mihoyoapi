package request

import (
	"errors"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/tools"
	"strings"
)

type AppCookies struct {
	GameUID     string
	Region      string
	LoginTicket string
	Stuid       string
	Stoken      string
	CookiesMap  map[string]string
}

func NewAppcookies(cookies string) (*AppCookies, error) {
	c := AppCookies{}
	err := c.ParseForLoginTicket(cookies)
	return &c, err
}

// ParseForLoginTicket 从user.mihoyo.com的cookies解析到对象
func (t *AppCookies) ParseForLoginTicket(cookies string) error {
	lst := strings.Split(cookies, ";")
	lens := len(lst)
	for i := 0; i < lens; i++ {
		lst[i] = strings.TrimSpace(lst[i])
	}
	maplist := []string{"login_ticket", "login_uid"}
	lenm := len(maplist)
	t.CookiesMap = make(map[string]string)
	for i := 0; i < lenm; i++ {
		for i1 := 0; i1 < lens; i1++ {
			if len(lst[i1]) > len(maplist[i]) {
				if lst[i1][:len(maplist[i])] == maplist[i] {
					value := lst[i1][len(maplist[i])+1:]
					t.CookiesMap[maplist[i]] = tools.DeepCopyStr(value)
					break
				}
			}
			index := strings.Index(lst[i1], "=")
			t.CookiesMap[lst[i1][:index]] = tools.DeepCopyStr(lst[i1][index+1:])
		}
		if t.CookiesMap[maplist[i]] == "" {
			return errors.New("cookies不完整")
		}
	}
	t.LoginTicket = t.CookiesMap["login_ticket"]
	t.Stuid = t.CookiesMap["login_uid"]
	t.CookiesMap["stuid"] = t.Stuid
	return nil
}

func (t *AppCookies) UpdateGameInfo(Region, GameUID string) {
	t.GameUID = GameUID
	t.Region = Region
}

func (t *AppCookies) GetHeadersMap(req RequestStruct, dsType int) map[string][]string {
	rm := make(map[string][]string)
	rm["x-rpc-app_version"] = []string{"2.35.2"}
	rm["x-rpc-client_type"] = []string{"5"}
	rm["x-rpc-device_id"] = []string{tools.GetUUID()}
	rm["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 12; " + define.DEVICE + " AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.73 Mobile Safari/537.36 miHoYoBBS/2.35.2"}
	if req.Sign {
		rm["X-Requested-With"] = []string{"com.mihoyo.hyperion"}
		rm["x-rpc-platform"] = []string{"android"}
		rm["x-rpc-device_model"] = []string{define.DEVICE}
		rm["x-rpc-device_name"] = []string{define.DEVICE}
		rm["x-rpc-channel"] = []string{"miyousheluodi"}
		rm["x-rpc-sys_version"] = []string{"6.0.1"}
		rm["Referer"] = []string{"https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=true&act_id=e202009291139501&utm_source=bbs&utm_medium=mys&utm_campaign=icon"}
		switch dsType {
		case 1:
			rm["DS"] = []string{GetDsSign()}
		case 2:
			rm["x-rpc-app_version"] = []string{"2.34.1"}
			rm["x-rpc-client_type"] = []string{"2"}
			rm["DS"] = []string{GetDsSign2()}
			rm["Referer"] = []string{"https://app.mihoyo.com"}
			rm["User-Agent"] = []string{"okhttp/4.8.0"}
		}
	} else {
		body := ""
		if req.Body == nil {
			body = ""
		} else {
			body = req.Body.GetData()
		}
		switch dsType {
		case 1:
			rm["DS"] = []string{GetDs(t.Region, req.Query, body)}
		case 2:
			rm["x-rpc-client_type"] = []string{"2"}
			rm["User-Agent"] = []string{"okhttp/4.8.0"}
			rm["Referer"] = []string{"https://app.mihoyo.com"}
			rm["DS"] = []string{GetDs2(req.Query, body)}
		}
	}
	rm["Cookie"] = []string{t.GetCookies()}
	return rm
}

// GetCookies 获取cookies的文本
func (t *AppCookies) GetCookies() string {
	str := ""
	for k, v := range t.CookiesMap {
		str += k + "=" + v + "; "
	}
	return str
}

func (t *AppCookies) ParseForLoginApp(cookies string) error {
	lst := strings.Split(cookies, ";")
	lens := len(lst)
	for i := 0; i < lens; i++ {
		lst[i] = strings.TrimSpace(lst[i])
	}
	maplist := []string{"stuid", "stoken"}
	//maplist := []string{"login_uid", "login_ticket", "ltoken", "ltuid"}
	lenm := len(maplist)
	t.CookiesMap = make(map[string]string)
	for i := 0; i < lenm; i++ {
		for i1 := 0; i1 < lens; i1++ {
			if len(lst[i1]) > len(maplist[i]) {
				if lst[i1][:len(maplist[i])] == maplist[i] {
					value := lst[i1][len(maplist[i])+1:]
					t.CookiesMap[maplist[i]] = tools.DeepCopyStr(value)
					break
				}
			}
		}
		if t.CookiesMap[maplist[i]] == "" {
			return errors.New("cookies不完整")
		}
	}
	t.Stuid = t.CookiesMap["stuid"]
	t.Stoken = t.CookiesMap["stoken"]
	return nil
}

package Cookies

import (
	"errors"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	"github.com/Huiyicc/mihoyoapi/tools"
	json "github.com/json-iterator/go"
	"io"
	"net/http"
	"strings"
)

type CookiesCore struct {
	CookiesMap map[string]string
	AccountId  string
	GameUID    string
	Region     string
}

// Parse 给定一个cookies,自动解析
func (t *CookiesCore) Parse(cookies string) error {
	lst := strings.Split(cookies, ";")
	lens := len(lst)
	for i := 0; i < lens; i++ {
		lst[i] = strings.TrimSpace(lst[i])
	}
	maplist := []string{"ltoken", "ltuid", "cookie_token", "account_id"}
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
	return t.CheckCookies()
}

// GetCookies 获取cookies的文本
func (t *CookiesCore) GetCookies() string {
	str := ""
	for k, v := range t.CookiesMap {
		str += k + "=" + v + "; "
	}
	return str
}

// CheckCookies 检查cookies是否有效,并返回默认角色的UID
func (t *CookiesCore) CheckCookies() error {
	url := "https://api-takumi.mihoyo.com/binding/api/getUserGameRolesByCookie?game_biz=hk4e_cn"
	cli := http.Client{}
	header, _ := http.NewRequest("GET", url, nil)
	header.Header.Set("cookie", t.GetCookies())
	req, err := cli.Do(header)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	var res CookiesResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	if res.Retcode != 0 {
		return errors.New(res.Message)
	}
	if len(res.Data.List) == 0 {
		return errors.New("米游社账号未绑定原神角色！")
	}
	//米游社默认展示的角色
	for i := 0; i < len(res.Data.List); i++ {
		if res.Data.List[i].IsChosen {
			t.GameUID = res.Data.List[i].GameUid
			t.Region = res.Data.List[i].Region
			t.AccountId = t.CookiesMap["account_id"]
			return nil
		}
	}
	return errors.New("无默认角色")
}

func (t *CookiesCore) Get(key string) string {
	return t.CookiesMap[key]
}

func (t *CookiesCore) GetHeadersMap(req request.RequestStruct) map[string][]string {
	rm := make(map[string][]string)
	if req.Sign {
		rm["x-rpc-app_version"] = []string{"2.35.2"}
		rm["x-rpc-client_type"] = []string{"5"}
		rm["x-rpc-device_id"] = []string{tools.GetUUID()}
		rm["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 12; " + define.DEVICE + " AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.73 Mobile Safari/537.36 miHoYoBBS/2.35.2"}
		rm["X-Requested-With"] = []string{"com.mihoyo.hyperion"}
		rm["x-rpc-platform"] = []string{"android"}
		rm["x-rpc-device_model"] = []string{define.DEVICE}
		rm["x-rpc-device_name"] = []string{define.DEVICE}
		rm["x-rpc-channel"] = []string{"miyousheluodi"}
		rm["x-rpc-sys_version"] = []string{"6.0.1"}
		rm["Referer"] = []string{"https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=true&act_id=e202009291139501&utm_source=bbs&utm_medium=mys&utm_campaign=icon"}
		rm["DS"] = []string{request.GetDsSign()}
	} else {
		rm["x-rpc-app_version"] = []string{"2.34.1"}
		rm["x-rpc-client_type"] = []string{"5"}
		rm["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 12; " + define.DEVICE + " AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.73 Mobile Safari/537.36 miHoYoBBS/2.35.2"}
		body := ""
		if req.Body == nil {
			body = ""
		} else {
			body = req.Body.GetData()
		}
		rm["DS"] = []string{request.GetDs(t.Region, req.Query, body)}
	}
	return rm
}

type CookiesResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []CookiesUserinfo `json:"list"`
	} `json:"data"`
}

type CookiesUserinfo struct {
	GameBiz    string `json:"game_biz"`
	Region     string `json:"region"`
	GameUid    string `json:"game_uid"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	IsChosen   bool   `json:"is_chosen"`
	RegionName string `json:"region_name"`
	IsOfficial bool   `json:"is_official"`
}

type BBSCookiesCore struct {
	CookiesMap map[string]string
	AccountId  string
	GameUID    string
	Region     string
}

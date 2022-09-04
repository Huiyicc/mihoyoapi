package genshin

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
	"net/http"
)

// GetBindingInfo 用于获取绑定的游戏信息
func (t *GameCore) GetBindingInfo() (GenShinInfo, error) {
	info := GenShinInfo{}
	req := request.UrlMap[define.MIHOYOAPP_API_BINDINGO].Copy()
	req.Query = fmt.Sprintf(req.Query, "hk4e_cn")
	cli := request.NewClient(t.Cookies)
	data, err := cli.Get(req, 1, func(r *http.Request) {
		r.Header["Referer"] = []string{"https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=true&act_id=e202009291139501&utm_source=bbs&utm_medium=mys&utm_campaign=icon"}
	})
	if err != nil {
		return info, err
	}
	var resp getBindingInfoResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return info, errors.New(string(data))
	}
	if err = info.parse(resp); err != nil {
		return info, err
	}
	return info, nil
}

type GenShinInfo struct {
	GameBiz    string `json:"game_biz"`
	Region     string `json:"region"`
	GameUid    string `json:"game_uid"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	IsChosen   bool   `json:"is_chosen"`
	RegionName string `json:"region_name"`
	IsOfficial bool   `json:"is_official"`
}

func (t *GenShinInfo) parse(r getBindingInfoResponse) error {
	if r.Retcode != 0 {
		return errors.New(r.Message)
	}
	if len(r.Data.List) == 0 {
		return errors.New("无默认角色")
	}
	t.GameBiz = r.Data.List[0].GameBiz
	t.Region = r.Data.List[0].Region
	t.GameUid = r.Data.List[0].GameUid
	t.Nickname = r.Data.List[0].Nickname
	t.Level = r.Data.List[0].Level
	t.IsChosen = r.Data.List[0].IsChosen
	t.RegionName = r.Data.List[0].RegionName
	t.IsOfficial = r.Data.List[0].IsOfficial
	return nil
}

/*func (t *getBindingInfoResponse) handleCode() error {
	if t.Retcode == 0 {
		return nil
	}
	return errors.New(t.Message)
}*/

type getBindingInfoResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			GameBiz    string `json:"game_biz"`
			Region     string `json:"region"`
			GameUid    string `json:"game_uid"`
			Nickname   string `json:"nickname"`
			Level      int    `json:"level"`
			IsChosen   bool   `json:"is_chosen"`
			RegionName string `json:"region_name"`
			IsOfficial bool   `json:"is_official"`
		} `json:"list"`
	} `json:"data"`
}

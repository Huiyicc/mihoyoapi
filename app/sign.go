package app

import (
	"errors"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
)

const (
	BBS_BH3 = "崩坏3"
	BBS_YS  = "原神"
	BBS_BH2 = "崩坏2"
	BBS_WD  = "未定事件簿"
	BBS_DBY = "大别野"
	BBS_SR  = "崩坏：星穹铁道"
	BBS_ZZZ = "绝区零"
)

var (
	mihoyobbs_List = map[string]bbsSetting{
		BBS_BH3: {
			ID:      "1",
			ForumID: "1",
			Name:    "崩坏3",
			Url:     "https://bbs.mihoyo.com/bh3/",
		},
		BBS_YS: {
			ID:      "2",
			ForumID: "26",
			Name:    "原神",
			Url:     "https://bbs.mihoyo.com/ys/",
		},
		BBS_BH2: {
			ID:      "3",
			ForumID: "30",
			Name:    "崩坏2",
			Url:     "https://bbs.mihoyo.com/bh2/",
		},
		BBS_WD: {
			ID:      "4",
			ForumID: "37",
			Name:    "未定事件簿",
			Url:     "https://bbs.mihoyo.com/wd/",
		},
		BBS_DBY: {
			ID:      "5",
			ForumID: "34",
			Name:    "大别野",
			Url:     "https://bbs.mihoyo.com/dby/",
		},
		BBS_SR: {
			ID:      "6",
			ForumID: "52",
			Name:    "崩坏：星穹铁道",
			Url:     "https://bbs.mihoyo.com/bh3/",
		},
		BBS_ZZZ: {
			ID:      "8",
			ForumID: "57",
			Name:    "绝区零",
			Url:     "https://bbs.mihoyo.com/zzz/",
		},
	}
)

// 米哈游分区设置
type bbsSetting struct {
	ID      string
	ForumID string
	Name    string
	Url     string
}

type bbsSignReqStruct struct {
	Gids string `json:"gids"`
}

// Sign 用于签到分区,对应常量 BBS_
// 成功返回本次签到获得的米游币数量
func (t *AppCore) Sign(bbsType string) (int, error) {
	uris, isOk := mihoyobbs_List[bbsType]
	if !isOk {
		return 0, errors.New("分区不存在")
	}
	resq := request.UrlMap[define.MIHOYOAPP_API_SIGN].Copy()
	resq.Body["gids"] = uris.ID
	cli := request.NewClient(t.Cookies)
	data, err := cli.Post(resq, 2, nil)
	if err != nil {
		return 0, err
	}
	var resp signResponst
	if err = json.Unmarshal(data, &resp); err != nil {
		return 0, errors.New(string(data))
	}
	return resp.Data.Points, nil
}

type signResponst struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Points int `json:"points"`
	} `json:"data"`
}

package request

import (
	"github.com/Huiyicc/mihoyoapi/define"
	json "github.com/json-iterator/go"
)

type BodyMap map[string]interface{}

func (t BodyMap) GetData() string {
	data, _ := json.Marshal(t)
	//t["body"] = string(data)
	return string(data)
}

type RequestStruct struct {
	Url    string
	Query  string
	Sign   bool
	Body   BodyMap
	AddMap map[string]string
}

var UrlMap = map[string]RequestStruct{
	//首页宝箱
	define.MYSINFO_API_INDEX: {
		Url:   "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/index",
		Query: "role_id=%s&server=%s",
	},
	//深渊
	define.MYSINFO_API_SPIRALABYSS: {
		Url:   "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/spiralAbyss",
		Query: "role_id=%s&server=%s",
	},
	//角色详情
	define.MYSINFO_API_CHARACTER: {
		Url: "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/character",
		Body: map[string]interface{}{
			"role_id": "",
			"server":  "",
		},
	},
	//树脂
	define.MYSINFO_API_DAILYNOTE: {
		Url:   "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/dailyNote",
		Query: "role_id=%s&server=%s",
	},
	//签到信息
	define.MYSINFO_API_BBSSIGN_INFO: {
		Url:   "https://api-takumi.mihoyo.com/event/bbs_sign_reward/info",
		Query: "act_id=e202009291139501&region=%s&uid=%s",
		Sign:  true,
	},
	//签到奖励
	define.MYSINFO_API_BBSSIGN_HOME: {
		Url:   "https://api-takumi.mihoyo.com/event/bbs_sign_reward/home",
		Query: "act_id=e202009291139501&region=%s&uid=%s",
		Sign:  true,
	},
	//签到
	define.MYSINFO_API_BBSSIGN: {
		Url: "https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign",
		Body: map[string]interface{}{
			"act_id": "e202009291139501",
			"region": "",
			"uid":    "",
		},
		Sign: true,
	},
	//详情
	define.MYSINFO_API_DETAIL: {
		Url:   "https://api-takumi.mihoyo.com/event/e20200928calculate/v1/sync/avatar/detail",
		Query: "uid=%s&region=%s&avatar_id=%s",
	},
	//札记
	define.MYSINFO_API_YSLEDGER: {
		Url:   "https://hk4e-api.mihoyo.com/event/ys_ledger/monthInfo",
		Query: "month=%s&bind_uid=%s&bind_region=%s",
	},
	//养成计算器
	define.MYSINFO_API_COMPUTE: {
		Url: "https://api-takumi.mihoyo.com/event/e20200928calculate/v2/compute",
	},
	//角色技能
	define.MYSINFO_API_AVATARSKILL: {
		Url:   "https://api-takumi.mihoyo.com/event/e20200928calculate/v1/avatarSkill/list",
		Query: "avatar_id=%s",
	},
	// app登陆第一阶段
	define.MIHOYOAPP_API_LOGINA: {
		Url:   "https://webapi.account.mihoyo.com/Api/cookie_accountinfo_by_loginticket",
		Query: "login_ticket=%s",
		Sign:  true,
	},
	// app登陆第二阶段
	define.MIHOYOAPP_API_LOGINB: {
		Url:   "https://api-takumi.mihoyo.com/auth/api/getMultiTokenByLoginTicket",
		Query: "login_ticket=%s&token_types=3&uid=%s",
		Sign:  true,
	},
	// app任务列表
	define.MIHOYOAPP_API_TASKS_LIST: {
		Url: "https://bbs-api.mihoyo.com/apihub/sapi/getUserMissionsState",
		//Url:   "https://api-takumi.mihoyo.com/apihub/wapi/getUserMissionsState",
		Query: "point_sn=myb",
		Sign:  true,
	},
	//app内讨论区签到
	define.MIHOYOAPP_API_SIGN: {
		Url:  "https://bbs-api.mihoyo.com/apihub/app/api/signIn",
		Body: make(map[string]interface{}),
	},
	//获取app内某讨论区帖子列表
	define.MIHOYOAPP_API_FORUM_LIST: {
		Url:   "https://bbs-api.mihoyo.com/post/api/getForumPostList",
		Query: "forum_id=%s&is_good=false&is_hot=false&page_size=20&sort_type=1",
	},
	//看帖
	define.MIHOYOAPP_API_FORUM_LOOK: {
		Url:   "https://bbs-api.mihoyo.com/post/api/getPostFull",
		Query: "post_id=%s",
	},
	//分享帖子
	define.MIHOYOAPP_API_FORUM_SHARE: {
		Url:   "https://bbs-api.mihoyo.com/apihub/api/getShareConf",
		Query: "entity_id=%s&s&entity_type=1",
		Sign:  true,
	},
	//点赞帖子
	define.MIHOYOAPP_API_FORUM_LIKE: {
		Url: "https://bbs-api.mihoyo.com/apihub/sapi/upvotePost",
		Body: map[string]interface{}{
			"post_id":   "",
			"is_cancel": false,
		},
		Sign: true,
	},
}

package app

import "github.com/Huiyicc/mihoyoapi/Cookies"

const (
	TASKS_MISSION_ID_BBS_SIGN       = 58 //讨论区签到
	TASKS_MISSION_ID_BBS_READ_POSTS = 59 //看帖子
	TASKS_MISSION_ID_BBS_LIKE_POSTS = 60 //给帖子点赞
	TASKS_MISSION_ID_BBS_SHARE      = 61 //分享帖子
)

// NewAppCoreFromCookiesStr 使用cookies文本初始化
func NewAppCoreFromCookiesStr(cookies string) (*AppCore, error) {
	c, err := Cookies.NewAppcookies(cookies)
	if err != nil {
		return nil, err
	}
	r := AppCore{
		Cookies: c,
	}
	r.Init(c)
	return &r, err
}

// NewAppCoreFromCookies 使用cookies结构体初始化
func NewAppCoreFromCookies(cookies *Cookies.AppCookies) (*AppCore, error) {
	r := AppCore{
		Cookies: cookies,
	}
	r.Init(cookies)
	return &r, nil
}

// AppCore 为app类
type AppCore struct {
	TasksInfo *TasksList
	Cookies   *Cookies.AppCookies
}

// Init 用于初始化类相关参数
func (t *AppCore) Init(cookies *Cookies.AppCookies) {
	t.Cookies = cookies
}

func Test() {

}

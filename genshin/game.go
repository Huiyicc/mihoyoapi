package genshin

import (
	"github.com/Huiyicc/mihoyoapi/request"
)

type GameCore struct {
	Cookies *request.AppCookies
}

func NewGameCore(cookies *request.AppCookies) GameCore {
	return GameCore{
		Cookies: cookies,
	}
}

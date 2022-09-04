package note

import (
	"fmt"
	"github.com/Huiyicc/mihoyoapi/bbs"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
)

type NoteCoreNew struct {
	cookies *request.AppCookies
}

func (t *NoteCoreNew) Init(c *request.AppCookies) {
	t.cookies = c
}

func (t *NoteCoreNew) Info() (MysNoteResponse, error) {
	req := request.UrlMap[define.MYSINFO_API_DAILYNOTE]
	req.Query = fmt.Sprintf(req.Query, t.cookies.GameUID, t.cookies.Region)
	cli := request.NewClient(t.cookies)
	var (
		body []byte
		err  error
		note MysNoteResponse
	)
	if body, err = cli.Get(req, 1, nil); err != nil {
		return MysNoteResponse{}, err
	}
	if err = json.Unmarshal(body, &note); err != nil {
		return MysNoteResponse{}, err
	}
	if err = SwitchCode(note.Retcode, note.Message); err != nil {
		return MysNoteResponse{}, err
	}
	note.Data.Region = t.cookies.Region
	note.Data.GameUID = t.cookies.GameUID
	return note, nil
}

func SwitchCode(code int, msg string) error {
	switch code {
	case 0:
		return nil
	case -1:
	case -100:
	case 1001:
	case 10001:
	case 10103:
		return bbs.ERROR_LOGIN_FAILURE
	case 1008:
		return bbs.ERROR_NO_ROLE
	case 10101:
		return bbs.ERROR_UPPER_LIMIT
	case 10102:
		if msg == "Data is not public for the user" {
			return bbs.ERROR_USER_IS_PRIVATE
		}
		return bbs.ERROR_NO_ROLE
	case -1002:
		return bbs.ERROR_INTERFACE_ERROR_1002
	}
	return bbs.ERROR_INTERFACE_ERROR
}

package app

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
)

// ShareForum 分享帖子
func (t *AppCore) ShareForum(postID string) error {
	req := request.UrlMap[define.MIHOYOAPP_API_FORUM_SHARE]
	req.Query = fmt.Sprintf(req.Query, postID)
	data, err := t.httpGet(req, 1, nil)
	if err != nil {
		return err
	}
	var resp shareForumResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return err
	}
	if err = resp.handleCode(); err != nil {
		return err
	}
	return nil
}

type shareForumResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Icon    string `json:"icon"`
		Url     string `json:"url"`
	} `json:"data"`
}

func (t *shareForumResponse) handleCode() error {
	if t.Retcode == 0 {
		return nil
	}
	return errors.New(t.Message)
}

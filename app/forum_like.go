package app

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
)

// ForumLike 点赞帖子
func (t *AppCore) ForumLike(postID string, isCancel bool) error {
	req := request.UrlMap[define.MIHOYOAPP_API_FORUM_LIKE].Copy()
	req.Body["post_id"] = postID
	req.Body["is_cancel"] = isCancel
	fmt.Println(req.Body)
	cli := request.NewClient(t.Cookies)
	data, err := cli.Post(req, 2, nil)
	if err != nil {
		return err
	}
	var resp forumLikeResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return errors.New(string(data))
	}
	if resp.Retcode != 0 {
		return errors.New(resp.Message)
	}
	return nil
}

type forumLikeResponse struct {
	Retcode int
	Message string
}

package app

import (
	"fmt"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
)

// ForumLike 点赞帖子
func (t *AppCore) ForumLike(postID string, isCancel bool) error {
	req := request.UrlMap[define.MIHOYOAPP_API_FORUM_LIKE]
	req.Body["post_id"] = postID
	req.Body["is_cancel"] = isCancel
	fmt.Println(req.Body)
	data, err := t.httpPost(req, 2, nil)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

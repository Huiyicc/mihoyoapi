package app

import (
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/request"
	json "github.com/json-iterator/go"
)

// GetTasksList 用于获取已完成的米游币任务列表
func (t *AppCore) GetTasksList() (*TasksList, error) {
	r := request.UrlMap[define.MIHOYOAPP_API_TASKS_LIST].Copy()
	cli := request.NewClient(t.Cookies)
	data, err := cli.Get(r, 1, nil)
	if err != nil {
		return nil, err
	}
	var pesp tatsksListResp
	if err = json.Unmarshal(data, &pesp); err != nil {
		return nil, err
	}
	var info TasksList
	return &info, info.parse(pesp)
}

// UpdateTasksList 用于更新结构体内任务信息
func (t *AppCore) UpdateTasksList() error {
	list, err := t.GetTasksList()
	if err != nil {
		return err
	}
	t.TasksInfo = list
	return nil
}

// GetUnfinishedTaskList 用于获取未完成的任务列表
func (t *AppCore) GetUnfinishedTaskList() ([]int, error) {
	var unfinishedTaskListMap = map[int]int{
		TASKS_MISSION_ID_BBS_SIGN:       0,
		TASKS_MISSION_ID_BBS_READ_POSTS: 0,
		TASKS_MISSION_ID_BBS_LIKE_POSTS: 0,
		TASKS_MISSION_ID_BBS_SHARE:      0,
	}
	list, err := t.GetTasksList()
	if err != nil {
		return nil, err
	}
	//先统计已做任务
	r := make([]int, 0, 4)
	for _, states := range list.TasksList {
		unfinishedTaskListMap[states.MissionId]++
	}
	//从字典中筛出未做任务
	mlen := []int{TASKS_MISSION_ID_BBS_SIGN, TASKS_MISSION_ID_BBS_READ_POSTS, TASKS_MISSION_ID_BBS_LIKE_POSTS, TASKS_MISSION_ID_BBS_SHARE}
	for i := 0; i < len(mlen); i++ {
		if unfinishedTaskListMap[mlen[i]] == 0 {
			r = append(r, mlen[i])
		}
	}
	return r, nil
}

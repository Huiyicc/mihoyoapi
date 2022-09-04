package app

import "errors"

// tatsksListResp 为米游币任务列表结构体
type tatsksListResp struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		States []struct {
			MissionId     int    `json:"mission_id"`     //任务id
			Process       int    `json:"process"`        //
			HappenedTimes int    `json:"happened_times"` //happened_times
			IsGetAward    bool   `json:"is_get_award"`   //
			MissionKey    string `json:"mission_key"`    //mission key
		} `json:"states"` //已完成的任务
		AlreadyReceivedPoints int  `json:"already_received_points"` //今日已获得积分
		TotalPoints           int  `json:"total_points"`            //总积分
		TodayTotalPoints      int  `json:"today_total_points"`      //今日最多可获得积分
		IsUnclaimed           bool `json:"is_unclaimed"`            //
		CanGetPoints          int  `json:"can_get_points"`          //今日还可以获取的总积分
	} `json:"data"`
}

// TasksList 为任务信息结构体
type TasksList struct {
	AlreadyReceivedPoints int           //今日已获得积分
	TotalPoints           int           //总积分
	TodayTotalPoints      int           //今日最多可获得积分
	CanGetPoints          int           //今日还可以获取的总积分
	IsUnclaimed           bool          //is unclaimed
	TasksList             []TasksStates //已完成的任务列表
}

type TasksStates struct {
	MissionId     int    `json:"mission_id"`     //任务id
	Process       int    `json:"process"`        //
	HappenedTimes int    `json:"happened_times"` //happened_times
	IsGetAward    bool   `json:"is_get_award"`   //
	MissionKey    string `json:"mission_key"`    //mission key
}

// parse 用于将内部response json转换为对外的数据结构体
func (t *TasksList) parse(r tatsksListResp) error {
	if r.Retcode != 0 {
		return errors.New(r.Message)
	}
	t.CanGetPoints = r.Data.CanGetPoints
	t.IsUnclaimed = r.Data.IsUnclaimed
	t.AlreadyReceivedPoints = r.Data.AlreadyReceivedPoints
	t.TodayTotalPoints = r.Data.TodayTotalPoints
	t.TotalPoints = r.Data.TotalPoints
	t.TasksList = make([]TasksStates, len(r.Data.States), len(r.Data.States))
	for i := 0; i < len(r.Data.States); i++ {
		t.TasksList[i].MissionId = r.Data.States[i].MissionId
		t.TasksList[i].MissionKey = r.Data.States[i].MissionKey
		t.TasksList[i].Process = r.Data.States[i].Process
		t.TasksList[i].HappenedTimes = r.Data.States[i].HappenedTimes
		t.TasksList[i].IsGetAward = r.Data.States[i].IsGetAward
	}
	return nil
}

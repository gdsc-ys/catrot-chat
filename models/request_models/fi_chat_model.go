package requestmodel

import (
	"catrot-chat/models"
)

type MessageSendRq struct {
	models.DefaultData
	MRID       int `json:"mrid"`
	Uid        int `json:"uid"`
	CounterUid int `json:"counter_uid"`
}

type ImageSendRq struct {
	MessageSendRq
}

type UnreadMsgModel struct {
	models.DefaultData
	MRID       int  `json:"mrid"`
	Fid        int  `json:"fid"`
	CounterSub bool `json:"counter_sub"`
}

type UnreadMsgRoomModel struct {
	models.DefaultData
	Fid int `json:"fid"`
}

type GetUpdatedInfoMsgRoomModel struct {
	Fid  int   `json:"fid"`
	MRID []int `json:"mrid"`
}

type MessageSendRqText struct {
	MessageSendRq
	Msg          string `json:"msg"`
	MsgType      any    `json:"msg_type"`
}

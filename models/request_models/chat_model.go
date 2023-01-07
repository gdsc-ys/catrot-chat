package requestmodel

import (
	"catrot-chat/models"
)

type MessageSendRq struct {
	models.DefaultData
	MRID       int `json:"mrid"`
	CounterUid int `json:"counter_uid"`
}

type ImageSendRq struct {
	MessageSendRq
}

type UnreadMsgModel struct {
	models.DefaultData
	MRID       int  `json:"mrid"`
	CounterSub bool `json:"counter_sub"`
}

type UnreadMsgRoomModel struct {
	models.DefaultData
}

type GetUpdatedInfoMsgRoomModel struct {
	MRID []int `json:"mrid"`
}

type MessageSendRqText struct {
	MessageSendRq
	Msg          string `json:"msg"`
	MsgType      any    `json:"msg_type"`
}

type ChangeRoomNameModel struct {
	MessageSendRq
	RoomName  		string `json:"room_name"`
}

type CreateMsgRoomModel struct {
	MessageSendRq
	RoomType       	string 	`json:"room_type"`
}

type MsgRoomDetailModel struct {
	MessageSendRq
	RoomId    		int `json:"room_id"`
	Uid    			int `json:"uid"`
}

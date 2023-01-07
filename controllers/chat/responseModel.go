package fichat

type sendImageResp200 struct {
	InsertedMessageId string `json:"insertedMessageId"`
}

type sendMessageResp200 struct {
	InsertedMessageId int    `json:"insertedMessageId"`
	RcvAutoMsg        string `json:"rcv_auto_msg"`
	RcvAutoMsgId      int    `json:"rcv_auto_msg_id"`
	IsMedalReceivable int    `json:"is_medal_receivable"`
	SentenceLinkId    int    `json:"sentence_link_id"`
	IsBool            bool   `json:"is_bad"`
}

type unreadMsgList struct {
	Id       int         `json:"id"`
	Fid      int         `json:"fid"`
	FiName   string      `json:"fi_name"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
	SendData string      `json:"send_date"`
}

type unreadMsgListRs struct {
	MsgList []unreadMsgList `json:"msg_list"`
}

type unreadMsgRoomListRs struct {
	Uid       int    `json:"uid"`
	Fid       int    `json:"fid"`
	MRID      int    `json:"mrid"`
	UnreadCnt int    `json:"unread_cnt"`
	RoomName  string `json:"room_name"`
	Msg       string `json:"msg"`
	Friend    bool   `json:"friend"`
	ImgUrl    string `json:"img_url"`
	SendData  string `json:"send_date"`
	IsSub     bool   `json:"is_sub"`
}

type getUpdatedInfoMsgRoomListRs struct {
	Uid      int    `json:"uid"`
	Fid      int    `json:"fid"`
	MRID     int    `json:"mrid"`
	RoomName string `json:"room_name"`
	Friend   bool   `json:"friend"`
	ImgUrl   string `json:"img_url"`
	IsSub    bool   `json:"is_sub"`
}

type emptyResp struct {
}


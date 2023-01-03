package fichat

import (
	requestmodel "catrot-chat/models/request_models"
	"context"
)

func getUnreadMsgList(ctx context.Context, req requestmodel.UnreadMsgModel, fullReturn bool) []map[string]interface{} {
	//TODO : read from db

	var unreadMsgList = []map[string]interface{}{}

	return unreadMsgList
}

func deleteMsgUnreadLog(mrid int, fid int) {
}

type unreadMsgRoomListModel struct {
	id       int
	mrid     int
	fid      int
	lastMqid int
	regDate  string
	isGroup  string
	lc       string
}

func getUnreadMsgRoomList(ctx context.Context, fid int) []unreadMsgRoomListModel {
	//TODO : read from db

	var unreadMsgRoomList []unreadMsgRoomListModel

	return unreadMsgRoomList
}

func getSingleUnreadCount(ctx context.Context, FID int, MRID int) int {
	//TODO : read from db

	var messageCnt int

	return messageCnt
}

func getSingleUnreadMsgRoomData(ctx context.Context, data unreadMsgRoomListModel) map[string]interface{} {
	//TODO : read from db

	resultData := make(map[string]interface{})

	return resultData
}


func getMsgRoomName(ctx context.Context, mrid int) string {
	//TODO : read from db

	return "no_name"
}

func getSingleMsgRoomData(ctx context.Context, mrid int, fid int) map[string]interface{} {

	//TODO : read from db

	resultData := make(map[string]interface{})

	return resultData
}

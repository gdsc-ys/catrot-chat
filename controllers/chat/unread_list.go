package fichat

import (
	requestmodel "catrot-chat/models/request_models"
	"context"

	db "catrot-chat/services/mysql"
)

func getUnreadMsgList(ctx context.Context, req requestmodel.UnreadMsgModel) []map[string]interface{} {
	sql := `
	SELECT T1.*
	FROM (
		SELECT id, send_uid as uid, msg, messageType, DATE_FORMAT(reg_date, "%Y-%m-%dT%H:%i:%s") as send_date 
		FROM catrot_chat.messages as msgs
		WHERE room_id = ?
	) as T1
	JOIN (
		SELECT last_mid
		FROM catrot_chat.unread_messages
		WHERE room_id = ?
		AND uid = ?
		LIMIT 1    
	) as T2
	ON T1.id >= T2.last_mid           
	`

	result, err := db.SlaveBun.QueryContext(ctx, sql, req.MRID, req.MRID, req.UID)
	if err != nil {
		return []map[string]interface{}{}
	}
	defer result.Close()

	var unreadMsgList = []map[string]interface{}{}
	for result.Next() {
		var id int
		var uid int
		var msg string
		var messageType any
		var sendDate string
		result.Scan(&id, &uid, &msg, &messageType, &sendDate)

		unreadMsgList = append(unreadMsgList, map[string]interface{}{
			"id":        	id,
			"uid":       	uid,
			"msg":       	msg,
			"messageType":	messageType,
			"send_date": 	sendDate,
		})
	}

	return unreadMsgList

}

func deleteMsgUnreadLog(ctx context.Context, room_id int, uid int) {
	result, err := db.SlaveBun.QueryContext(ctx, `
	DELETE FROM catrot_chat.unread_messages
	WHERE room_id=? and uid=?
	`, room_id, uid)

	if err != nil {
		print(err)
	}
	defer result.Close()

}

type unreadMsgRoomListModel struct {
	id       	int
	room_id     int
	uid      	int
	last_mid 	int
	reg_date  	string
}

func getUnreadMsgRoomList(ctx context.Context, uid int) []unreadMsgRoomListModel {
	result, err := db.SlaveBun.QueryContext(ctx, `
	SELECT *
	FROM catrot_chat.unread_messages
	WHERE uid=?
	`, uid)

	if err != nil {
		return make([]unreadMsgRoomListModel, 0)
	}
	defer result.Close()

	var unreadMsgRoomList []unreadMsgRoomListModel
	for result.Next() {
		var id int
		var room_id int
		var uid int
		var last_mid int
		var reg_date string
		result.Scan(&id, &room_id, &uid, &last_mid, &reg_date)

		unreadMsgRoomList = append(unreadMsgRoomList, unreadMsgRoomListModel{
			id:       	id,
			room_id: 	room_id,
			uid:      	uid,
			last_mid: 	last_mid,
			reg_date:  	reg_date,
		})
	}

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

func getSingleMsgRoomData(ctx context.Context, roomId int, fid int) map[string]interface{} {

	//TODO : read from db

	resultData := make(map[string]interface{})

	return resultData
}

func chageChatRoomName(ctx context.Context, name string, roomId int, uid int){
	result, err := db.SlaveBun.QueryContext(ctx, `
	UPDATE catrot_chat.chat_room_member
	SET room_name=?
	WHERE room_id=? and uid=?;
	`, name, roomId, uid)

	if err != nil {
		print(err)
	}
	defer result.Close()
}

func createChatRoom(ctx context.Context, uid int, counterUid int, roomType string) int {
	result, err := db.SlaveBun.QueryContext(ctx, `
	INSERT INTO catrot_chat.chat_room (create_uid, room_type)
	VALUES(?, ?)
	`, uid, counterUid, roomType)

	if err != nil {
		print(err)
		return -1
	}
	defer result.Close()

	for result.Next() {
		var room_id int
		result.Scan(&room_id)

		createRoomMember(ctx, uid, room_id)
		createRoomMember(ctx, counterUid, room_id)
		return room_id
	}

	return -1
}


func createRoomMember(ctx context.Context, uid int, roomid int) {
	result, err := db.SlaveBun.QueryContext(ctx, `
	INSERT INTO catrot_chat.chat_room_member (room_id, uid)
	VALUES(?, ?)
	`, roomid, uid)

	if err != nil {
		print(err)
	}
	defer result.Close()
}
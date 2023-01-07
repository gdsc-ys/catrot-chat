package fichat

import (
	requestmodel "catrot-chat/models/request_models"
	db "catrot-chat/services/mysql"
	"context"
	"sync"
)

type messageData struct {
	textMsg     string
	msgType     any
	reqData     *requestmodel.MessageSendRq
	insertedId  int64
	ctx         context.Context
}


func newMessageDateForText(ctx context.Context, msg string, msgType any, reqData *requestmodel.MessageSendRq) *messageData {
	return &messageData{
		textMsg:     msg,
		msgType:     msgType,
		reqData:     reqData,
		ctx:         ctx,
	}
}

func (m *messageData) InsertMessageQueue() int64 {
	m.textInsertToMsgQueue()
	return m.insertedId
}

// s3에 업로드 하고 주소를 msg_queue에 인서트 한다.
func (m *messageData) imageInsertToMsgQueue() {
}

func (m *messageData) uploadToS3(key string, body *[]byte, contentType string, w *sync.WaitGroup) {
}

// 텍스트를 msg_queue에 업로드 한다.
func (m *messageData) textInsertToMsgQueue() {
	res, err := db.MasterBun.ExecContext(m.ctx, `
		INSERT INTO catrot_chat.messages
		(msg, send_uid, room_id, messageType)
		VALUES (?, ?, ?, ?)
	`, m.textMsg, m.reqData.UID, m.reqData.MRID, m.msgType)

	if err != nil {
		m.insertedId = 0
	} else {
		id, _ := res.LastInsertId()
		m.insertedId = id
	}

	if m.insertedId > 0 {
		_, err := db.MasterBun.ExecContext(m.ctx, `
			INSERT INTO catrot_chat.unread_messages
			(room_id, uid, last_mid)
			VALUES (?, ?, ?)
		`, m.reqData.MRID, m.reqData.UID, m.insertedId)

		if(err != nil){
			print(err)
		}
	}
}

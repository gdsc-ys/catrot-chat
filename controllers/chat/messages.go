package fichat

import (
	"catrot-chat/models"
	requestmodel "catrot-chat/models/request_models"
	"context"
	"sync"

	"github.com/gin-gonic/gin"
)

type messageData struct {
	textMsg     string
	msgType     any
	reqData     *requestmodel.MessageSendRq
	insertedId  int64
	ctx         context.Context
}

func newMessageDataForImage(c *gin.Context, reqData *requestmodel.ImageSendRq, cd *models.DefaultData) (*messageData, error) {
	var data messageData

	data = messageData{
		msgType:     "img",
		reqData:     &reqData.MessageSendRq,
		ctx:         c.Request.Context(),
	}
	return &data, nil
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
	// res, err := db.MasterBun.ExecContext(m.ctx, `
	// 	INSERT INTO catrot-chat.msg_queue
	// 	(msg, send_fid, mrid, lc, cc, data)
	// 	VALUES (?, ?, ?, ?, ?, ?)
	// `, m.textMsg, m.reqData.Fid, m.reqData.MRID, m.reqData.LC, m.reqData.CC, m.msgType)

	// if err != nil {
	// 	m.insertedId = 0
	// } else {
	// 	id, _ := res.LastInsertId()
	// 	m.insertedId = id
	// }

	// if m.insertedId > 0 {
	// }
}


func (m *messageData) getPreviousMsg() string {
	return ""
}

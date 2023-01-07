package fichat

import (
	requestmodel "catrot-chat/models/request_models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary SendMessage
// @Schemes
// @Description Send Msg
// @Param			param		body		requestmodel.MessageSendRqText	true "data"
// @Tags chat
// @Accept  json
// @Produce json
// @Success 200 {object} sendMessageResp200
// @Router /chat/send_msg [post]
func SendMessage(c *gin.Context) {
	var reqData requestmodel.MessageSendRqText

	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("(%v)", err),
		})
		c.Abort()
		return
	}

	if reqData.CounterUid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"insertedMessageId": "DELETE_" + time.Now().String(),
		})
		return
	}

	messageData := newMessageDateForText(c.Request.Context(), reqData.Msg, reqData.MsgType, &reqData.MessageSendRq)
	insertId := messageData.InsertMessageQueue()

	pushModel := newModel(&reqData.MessageSendRq)
	go pushModel.SendPush(reqData.Msg)

	c.JSON(http.StatusOK, gin.H{"insertedMessageId":   insertId,})
	
}

// @Summary UnreadMsgList
// @Schemes
// @Description 안 읽은 메세지 리스트 받아오기, 받아온 후에는 unread_log에서 삭제 된다.
// @Param			param		body		requestmodel.UnreadMsgModel	true		"data"
// @Tags chat
// @accept json
// @Produce json
// @Success 200 {object} unreadMsgListRs
// @Failure 422 {object} emptyResp
// @Router /chat/unread_msg_list [post]
func UnreadMsgList(c *gin.Context) {
	var data requestmodel.UnreadMsgModel

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("(%v)", err),
		})
		c.Abort()
		return
	}

	var unreadMsgList []map[string]interface{} = getUnreadMsgList(c.Request.Context(), data)

	go deleteMsgUnreadLog(c.Request.Context(), data.MRID, data.UID)


	if len(unreadMsgList) > 0 {
		unreadMsgRoomJson, _ := json.Marshal(map[string]interface{}{
			"msg_list": unreadMsgList,
		})
		var unreadMsgRs unreadMsgListRs
		json.Unmarshal(unreadMsgRoomJson, &unreadMsgRs)
		c.JSON(http.StatusOK, unreadMsgRs)

	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg_list": []interface{}{},
		})
	}
}

// @Summary UnreadMsgRoomList
// @Schemes
// @Description `내 안읽은 메시지 대화방 목록 정보 가져오기`
// @Param			param		body		requestmodel.UnreadMsgRoomModel	true		"data"
// @Tags chat
// @accept json
// @Produce json
// @Success 200 {object} []unreadMsgRoomListRs
// @Failure 422 {object} emptyResp
// @Router /chat/unread_msg_room_list [post]
func UnreadMsgRoomList(c *gin.Context) {
	var data requestmodel.UnreadMsgRoomModel

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("(%v)", err),
		})
		c.Abort()
		return
	}

	myUnreadMsgRoomList := getUnreadMsgRoomList(c.Request.Context(), data.UID)

	var finalResult = []map[string]interface{}{}
	ch := make(chan map[string]interface{}, len(myUnreadMsgRoomList))

	for i := 0; i < len(myUnreadMsgRoomList); i++ {
		finalResult = append(finalResult, <-ch)
	}

	unreadMsgRoomJson, _ := json.Marshal(finalResult)
	var unreadMsgRoomRs []unreadMsgRoomListRs
	json.Unmarshal(unreadMsgRoomJson, &unreadMsgRoomRs)
	c.JSON(http.StatusOK, unreadMsgRoomRs)
}


// @Summary ChangeRoomName
// @Schemes
// @Description `채팅방 이름 변경`
// @Param			param		body		requestmodel.ChangeRoomNameModel	true		"data"
// @Tags chat
// @accept json
// @Produce json
// @Success 200 {object} sendMessageResp200
// @Failure 422 {object} emptyResp
// @Router /chat/change_room_name [post]
func ChangeRoomName(c *gin.Context) {
	var data requestmodel.ChangeRoomNameModel

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("(%v)", err),
		})
		c.Abort()
		return
	}

	chageChatRoomName(c.Request.Context(), data.RoomName, data.MRID, data.UID);

	c.JSON(http.StatusOK, gin.H{"roomName": data.RoomName,})
}


// @Summary CreateMsgRoom
// @Schemes
// @Description `채팅방 생성`
// @Param			param		body		requestmodel.CreateMsgRoomModel	true		"data"
// @Tags chat
// @accept json
// @Produce json
// @Success 200 {object} sendMessageResp200
// @Failure 422 {object} emptyResp
// @Router /chat/create_msg_room [post]
func CreateMsgRoom(c *gin.Context) {
	var data requestmodel.CreateMsgRoomModel

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("(%v)", err),
		})
		c.Abort()
		return
	}

	roomId := createChatRoom(c, data.UID, data.CounterUid, data.RoomType)

	if(roomId>0){
		c.JSON(http.StatusOK, gin.H{"roomId":  roomId})

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": 500})
	}
}

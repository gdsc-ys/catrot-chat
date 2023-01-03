package routers

import (
	chat "catrot-chat/controllers/chat"

	"github.com/gin-gonic/gin"
)

func setChatRoutes(router *gin.RouterGroup) {
	router.POST("/send_msg", chat.SendMessage)
	router.POST("/unread_msg_list", chat.UnreadMsgList)
	router.POST("/unread_msg_room_list", chat.UnreadMsgRoomList)
}

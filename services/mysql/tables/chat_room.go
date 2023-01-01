package bumcoming

import (
	"time"

	"github.com/uptrace/bun"
)

type ChatRoom struct {
	bun.BaseModel `bun:"table:catrot_chat.chat_room,alias:cr"`

	RoomID      int       	`bun:"room_id"`
	CreateUid   int       	`bun:"create_uid"`
	RoomType   	string   	`bun:"room_type"`
	PrevUid 	int       	`bun:"prev_uid"`
	PrevMsg 	string     	`bun:"prev_msg"`
	RegDate   	time.Time 	`bun:"reg_date"`
}

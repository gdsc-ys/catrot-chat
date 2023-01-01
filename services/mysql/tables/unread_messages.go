package bumcoming

import (
	"time"

	"github.com/uptrace/bun"
)

type UnreadMessages struct {
	bun.BaseModel `bun:"table:catrot_chat.unread_messages,alias:ums"`

	ID			int		  	`bun:"id"`
	RoomId    	int       	`bun:"room_id"`
	Uid     	int 		`bun:"uid"`
	LastMid		int       	`bun:"last_mid"`
	RegDate		time.Time 	`bun:"reg_date"`
}

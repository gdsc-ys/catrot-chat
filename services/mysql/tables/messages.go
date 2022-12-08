package bumcoming

import (
	"time"

	"github.com/uptrace/bun"
)

type Messages struct {
	bun.BaseModel `bun:"table:catrot_chat.messages,alias:mgs"`

	ID      	int    		`bun:"id"`
	Msg     	string    	`bun:"msg"`
	SendUid 	int       	`bun:"send_uid"`
	RoomID    	int     	`bun:"room_id"`
	MessageType	string    	`bun:"messageType"`
	RegDate		time.Time 	`bun:"reg_date"`
}

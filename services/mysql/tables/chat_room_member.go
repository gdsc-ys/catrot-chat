package bumcoming

import (
	"time"

	"github.com/uptrace/bun"
)

type ChatRoomMember struct {
	bun.BaseModel `bun:"table:catrot_chat.chat_room_member,alias:crm"`

	ID        	int			`bun:"id"`
	RoomID  	int  		`bun:"room_id"`
	Uid   		int			`bun:"uid"`
	RoomName 	string		`bun:"room_name"`
	PushState 	string		`bun:"push_state"`
	RegDate 	time.Time	`bun:"reg_date"`
}

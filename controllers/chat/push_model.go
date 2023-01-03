package fichat

import (
	requestmodel "catrot-chat/models/request_models"
	fcm_util "catrot-chat/services/firebase"
	"context"
	"fmt"
	"strconv"

	"firebase.google.com/go/v4/messaging"
)

type pushModel struct {
	uid    		int
	name   		string
	mrid   		int
	senderName 	string
	pushType    string
	message    	string
	token		string

	ctx context.Context
}

func newModel(data *requestmodel.MessageSendRq) *pushModel {

	model := pushModel{
		mrid:       data.MRID,
		pushType:   "chat_msg",
		ctx:        context.Background(),
	}
	model.setSenderFiName()

	return &model
}

// fi 이름을 가져오는 메서드
func (p *pushModel) setSenderFiName() {
}

// 푸시 보내는 메서드
func (p *pushModel) SendPush(message string) {

	p.message = message
	p.sendNotificationPush()
}


// 알림 보내는 푸시
func (p *pushModel) sendNotificationPush() {

	message := messaging.Message{
		Token: p.token,
		Data: map[string]string{
			"push_type": p.pushType,
			"mrid":      strconv.Itoa(p.mrid),
		},
		Notification: &messaging.Notification{
			Title: p.getPushTitle(),
			Body:  p.getPushBody(),
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Sound:     "Default",
				Tag:       fmt.Sprintf("fi_msg#%d#%d", p.mrid),
				ChannelID: "MessagePushChannelId",
			},
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Sound: "Default",
				},
			},
		},
	}

	fcm_util.FCMClient.Send(p.ctx, &message)
}

func (p *pushModel) getPushTitle() string {
	return p.senderName
}

func (p *pushModel) getPushBody() string {
	return "(" + p.name + " " + p.senderName + ") : " + p.message
}


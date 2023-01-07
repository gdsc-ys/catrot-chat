package fcm_util

import (
	"context"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type firebaseApp struct {
	*firebase.App
}

var FCMClient *messaging.Client

func SetFirebaseApp(wg *sync.WaitGroup) {
	defer wg.Done()

	opt := option.WithCredentialsFile("public/firebase/catrotapp-pby42-2ec310a29m.json")
	newApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	app := firebaseApp{
		newApp,
	}

	var wgSub sync.WaitGroup
	wgSub.Add(1)
	go app.setFCMClient(&wgSub)
	wgSub.Wait()
}

func (app firebaseApp) setFCMClient(wg *sync.WaitGroup) {
	defer wg.Done()

	client, err := app.Messaging(context.Background())
	if err != nil {
		panic(err)
	}
	FCMClient = client
}

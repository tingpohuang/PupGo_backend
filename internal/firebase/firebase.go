package firebase

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var (
	once sync.Once
)

var instance *firebase_controller

type firebase_controller struct {
	FApp *firebase.App
}

func GetApp() *firebase_controller {
	once.Do(func() {
		var err error
		settingJson, _ := json.Marshal(setting)
		opt := option.WithCredentialsJSON(settingJson)

		config := &firebase.Config{
			ProjectID: "pupgo-e03ef",
			// ServiceAccountID: "",
		}
		fapp, err := firebase.NewApp(context.Background(), config, opt)
		instance = &firebase_controller{
			FApp: fapp,
		}
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	})
	return instance
}

func (fctr *firebase_controller) SendNotification(ctx context.Context, message *messaging.Message) error {
	if ctx == nil {
		ctx = context.Background()
	}
	client, err := fctr.FApp.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
		return err
	}
	_, err = client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (fctr *firebase_controller) SendNotificationMultiDevices(ctx context.Context, message *messaging.MulticastMessage) error {
	if ctx == nil {
		ctx = context.Background()
	}
	client, err := fctr.FApp.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
		return err
	}
	_, err = client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

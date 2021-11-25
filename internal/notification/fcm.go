package notification

import (
	"context"
	"log"

	"github.com/tingpo/pupgobackend/internal/firebase"
	"github.com/tingpo/pupgobackend/internal/gorm"
)

type Notification struct {
}

func (n *Notification) SendFriendsInviteMessage(ctx context.Context, petId string, recommendId string, s *gorm.SQLCnter) {
	devices, err := s.FindUserDeviceID(ctx, petId)
	if err != nil {
		log.Fatal(err)
		return
	}
	tokens, err := UserDeviceToTokens(devices)
	if err != nil {
		log.Fatal(err)
		return
	}
	msg, err := n.generateFriendsInviteMessage(ctx, petId, recommendId, tokens)
	if err != nil {
		log.Fatal(err)
		return
	}
	app := firebase.GetApp()
	err = app.SendNotificationMultiDevices(ctx, msg)
	if err != nil {
		log.Fatal(err)
		return
	}
}
func (n *Notification) SendNewFriendMessage(ctx context.Context, petId string, recommendId string, s *gorm.SQLCnter) {
	devices, err := s.FindUserDeviceID(ctx, petId)
	if err != nil {
		log.Fatal(err)
		return
	}
	tokens, err := UserDeviceToTokens(devices)
	if err != nil {
		log.Fatal(err)
		return
	}
	msg, err := n.generateNewFriendMessage(ctx, petId, recommendId, tokens)
	if err != nil {
		log.Fatal(err)
		return
	}
	app := firebase.GetApp()
	err = app.SendNotificationMultiDevices(ctx, msg)
	if err != nil {
		log.Fatal(err)
		return
	}
}

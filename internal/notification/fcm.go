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
	tokens, err := PetIDToTokens(ctx, petId, s)
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
	tokens, err := PetIDToTokens(ctx, petId, s)
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
func (n *Notification) SendNewParticipantsMessage(ctx context.Context, petId string, recommendId string, s *gorm.SQLCnter) {
	tokens, err := PetIDToTokens(ctx, petId, s)
	if err != nil {
		log.Fatal(err)
		return
	}
	msg, err := n.generateNewParticipantsMessage(ctx, petId, recommendId, tokens)
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
func (n *Notification) SendEventJoinedMessage(ctx context.Context, eventId string, petId string, s *gorm.SQLCnter) {
	tokens, err := PetIDToTokens(ctx, petId, s)
	if err != nil {
		log.Fatal(err)
		return
	}
	msg, err := n.generateEventJoinedMessage(ctx, eventId, petId, tokens)
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

func (n *Notification) SendEventContentUpdateMessage(ctx context.Context, eventId string, s *gorm.SQLCnter) {
	tokens, err := EventIDToTokens(ctx, eventId, s)
	if err != nil {
		log.Fatal(err)
		return
	}
	msg, err := n.generateEventContentUpdateMessage(ctx, eventId, tokens)
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

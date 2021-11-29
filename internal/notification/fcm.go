package notification

import (
	"context"
	"fmt"
	"log"

	"github.com/tingpo/pupgobackend/internal/firebase"
	"github.com/tingpo/pupgobackend/internal/gorm"
)

type Notification struct {
}

func (n *Notification) SendFriendsInviteMessage(ctx context.Context, petId string, recommendId string, s *gorm.SQLCnter) {
	panic("abandon")
	// tokens, err := s.FindDeviceByPetID(ctx, petId)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	// msg, err := n.generateFriendsInviteMessage(ctx, petId, recommendId, tokens)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// app := firebase.GetApp()
	// err = app.SendNotificationMultiDevices(ctx, msg)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
}
func (n *Notification) SendNewFriendMessage(ctx context.Context, petId string, recommendId string, s *gorm.SQLCnter) {
	// write notification db
	nMsg := NewNotification()
	uid, err := s.GetUserIdbyPetId(ctx, petId)
	if err != nil {
		print(err, petId)
	}
	nMsg.User_id = *uid
	nMsg.Pet_id = recommendId
	nMsg.Notification_type = gorm.Notification_NewFriend

	if err = s.CreateNotification(ctx, nMsg); err != nil {
		print(err, nMsg)
	}
	//
	tokens, err := s.FindDeviceByPetID(ctx, petId)
	if err != nil {
		log.Print(err)
		return
	}

	msg, err := n.generateNewFriendMessage(ctx, petId, recommendId, tokens)
	if err != nil {
		log.Print(err)
		return
	}
	app := firebase.GetApp()
	err = app.SendNotificationMultiDevices(ctx, msg)
	if err != nil {
		log.Print(err)
		return
	}
}
func (n *Notification) SendNewParticipantsMessage(ctx context.Context, applicantId string, eventId string, s *gorm.SQLCnter) {
	holderId := s.FindHolderIdByEventId(ctx, eventId)
	tokens, err := s.FindDeviceByPetID(ctx, holderId)
	if err != nil {
		log.Print(err)
		return
	}
	// write notification db
	nMsg := NewNotification()
	uid, err := s.GetUserIdbyPetId(ctx, holderId)
	if err != nil {
		print(err, holderId)
	}
	nMsg.User_id = *uid
	nMsg.Pet_id = applicantId
	nMsg.Event_id = eventId
	nMsg.Notification_type = gorm.Notification_NewParticipants
	if err = s.CreateNotification(ctx, nMsg); err != nil {
		print(err, nMsg)
	}
	//

	if err = s.CreateNotification(ctx, nMsg); err != nil {
		print(err, nMsg)
	}
	msg, err := n.generateNewParticipantsMessage(ctx, eventId, applicantId, tokens)
	if err != nil {
		log.Print(err)
		return
	}
	app := firebase.GetApp()
	err = app.SendNotificationMultiDevices(ctx, msg)
	if err != nil {
		log.Print(err)
		return
	}
}
func (n *Notification) SendEventJoinedMessage(ctx context.Context, eventId string, petId string, s *gorm.SQLCnter) {

	// write notification db
	nMsg := NewNotification()
	uid, err := s.GetUserIdbyPetId(ctx, petId)
	if err != nil {
		print(err, petId)
	}
	nMsg.User_id = *uid
	nMsg.Pet_id = petId
	nMsg.Event_id = eventId
	nMsg.Notification_type = gorm.Notification_EventJoined
	if err = s.CreateNotification(ctx, nMsg); err != nil {
		fmt.Print(err, nMsg)
	}
	fmt.Print(nMsg)
	tokens, err := s.FindDeviceByPetID(ctx, petId)
	if err != nil {
		log.Print(err)
		return
	}
	msg, err := n.generateEventJoinedMessage(ctx, eventId, petId, tokens)
	if err != nil {
		log.Print(err)
		return
	}
	app := firebase.GetApp()
	err = app.SendNotificationMultiDevices(ctx, msg)
	if err != nil {
		log.Print(err)
		return
	}
}

func (n *Notification) SendEventContentUpdateMessage(ctx context.Context, eventId string, s *gorm.SQLCnter) {
	// write notification db

	nMsg := NewNotification()
	uids, err := s.FindUserIdListByEventId(ctx, eventId)
	if err != nil {
		print(err, eventId)
	}
	for i := 0; i < len(uids); i++ {
		uid := uids[i]
		nMsg.User_id = uid
		nMsg.Event_id = eventId
		nMsg.Notification_type = gorm.Notification_EventContentUpdate
		if err = s.CreateNotification(ctx, nMsg); err != nil {
			print(err, nMsg)
		}
	}
	//
	tokens, err := s.FindDeviceByAllParticipant(ctx, eventId)
	if err != nil {
		log.Print(err)
		return
	}
	msg, err := n.generateEventContentUpdateMessage(ctx, eventId, tokens)
	if err != nil {
		log.Print(err)
		return
	}
	app := firebase.GetApp()
	err = app.SendNotificationMultiDevices(ctx, msg)
	if err != nil {
		log.Print(err)
		return
	}
}

func (n *Notification) SendEventsToFriends(ctx context.Context, eventId string, s *gorm.SQLCnter) {
	// write notification db
	nMsg := NewNotification()
	nMsg.Event_id = eventId
	nMsg.Notification_type = gorm.Notification_EventsToFriends
	// uids, err := s.FindUserIdListByEventId(ctx, eventId)
	holderId := s.FindHolderIdByEventId(ctx, eventId)
	friendIds, err := s.GetFriendsPetIdByPetId(ctx, holderId)
	if err != nil {
		print(err, holderId, eventId)
	}
	friendUserIds, err := s.GetUserIdsbyPetIds(ctx, friendIds)
	if err != nil {
		print(err, holderId, eventId)
	}
	for i := 0; i < len(friendUserIds); i++ {
		friendUserId := friendUserIds[i]
		nMsg.User_id = friendUserId
		if err = s.CreateNotification(ctx, nMsg); err != nil {
			print(err, nMsg)
		}
	}
	tokens, err := s.FindDeviceByPetIDs(ctx, friendIds)
	if err != nil {
		// log.Print(err)
		return
	}
	msg, err := n.generateEventsToFriendsMessage(ctx, eventId, tokens)
	if err != nil {
		// log.Print(err)
		return
	}
	app := firebase.GetApp()
	err = app.SendNotificationMultiDevices(ctx, msg)
	if err != nil {
		// log.Print(err)
		return
	}
}

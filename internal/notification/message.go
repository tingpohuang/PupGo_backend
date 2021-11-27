package notification

import (
	"context"
	"errors"

	"firebase.google.com/go/messaging"
)

func (n *Notification) generateFriendsInviteMessage(ctx context.Context, pet_id string, recommend_id string, tokens []string) (m *messaging.MulticastMessage, err error) {
	if recommend_id == "" {
		return nil, errors.New("recommend_id is empty")
	}
	if tokens == nil {
		return nil, errors.New("tokens should not be empty")
	}
	m = &messaging.MulticastMessage{
		Data: map[string]string{
			"recommend_id":      recommend_id,
			"pet_id":            pet_id,
			"notification_type": "FriendsInvite",
			"click_action":      "FLUTTER_NOTIFICATION_CLICK",
		},
		Notification: &messaging.Notification{
			Title: "New Friends Invitation",
			Body:  "A new freinds is sending invitation to you!\n Click to open PetProfile!\n",
		},
		Tokens: tokens,
	}
	return m, err
}

func (n *Notification) generateNewFriendMessage(ctx context.Context, pet_id string, recommend_id string, tokens []string) (m *messaging.MulticastMessage, err error) {
	if recommend_id == "" {
		return nil, errors.New("recommend_id is empty")
	}
	if tokens == nil {
		return nil, errors.New("tokens should not be empty")
	}
	m = &messaging.MulticastMessage{
		Data: map[string]string{
			"notification_type": "NewFriend",
		},
		Notification: &messaging.Notification{
			Title: "New Friends!",
			Body:  "You get a new friends!\n Click to open PetProfile!\n",
		},
		Tokens: tokens,
	}
	return m, err
}
func (n *Notification) generateNewParticipantsMessage(ctx context.Context, eventId string, applicantId string, tokens []string) (m *messaging.MulticastMessage, err error) {
	if applicantId == "" {
		return nil, errors.New("applier is empty")
	}
	if tokens == nil {
		return nil, errors.New("tokens should not be empty")
	}
	m = &messaging.MulticastMessage{
		Data: map[string]string{
			"event_id":          eventId,
			"applicant_id":      applicantId,
			"notification_type": "NewParticipants",
		},
		Notification: &messaging.Notification{
			Title: "New applicant request!",
			Body:  "A new applicant want to joined the event you hold\n",
		},
		Tokens: tokens,
	}
	return m, err
}

func (n *Notification) generateEventJoinedMessage(ctx context.Context, event_id string, applicantId string, tokens []string) (m *messaging.MulticastMessage, err error) {
	if applicantId == "" {
		return nil, errors.New("applier is empty")
	}
	if tokens == nil {
		return nil, errors.New("tokens should not be empty")
	}
	m = &messaging.MulticastMessage{
		Data: map[string]string{
			"event_id":          event_id,
			"notification_type": "EventJoined",
		},
		Notification: &messaging.Notification{
			Title: "You joined a event!",
			Body:  "Holder accept your event applicant! congratulations!\n",
		},
		Tokens: tokens,
	}
	return m, err
}
func (n *Notification) generateEventContentUpdateMessage(ctx context.Context, event_id string, tokens []string) (m *messaging.MulticastMessage, err error) {
	if event_id == "" {
		return nil, errors.New("event_id is empty")
	}
	if tokens == nil {
		return nil, errors.New("tokens should not be empty")
	}
	m = &messaging.MulticastMessage{
		Data: map[string]string{
			"event_id":          event_id,
			"notification_type": "EventContentUpdate",
		},
		Notification: &messaging.Notification{
			Title: "A Event Content Updates",
			Body:  "Check the app for the event content update\n",
		},
		Tokens: tokens,
	}
	return m, err
}

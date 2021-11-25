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
			"recommend_id": recommend_id,
			"pet_id":       pet_id,
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
		Notification: &messaging.Notification{
			Title: "New Friends!",
			Body:  "You get a new friends!\n Click to open PetProfile!\n",
		},
		Tokens: tokens,
	}
	return m, err
}

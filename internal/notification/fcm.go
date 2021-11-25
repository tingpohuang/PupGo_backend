package notification

import (
	"context"
	"errors"
	"log"

	"firebase.google.com/go/messaging"
	"github.com/appleboy/go-fcm"
	"github.com/tingpo/pupgobackend/internal/gorm"
)

type Notification struct {
}

// var (
// 	token = "AAAAA_iwUx8:APA91bFWyMYgeeflt-a2IVZr4gHuz2TxR7er1RejiFr8H8HUHKWkW_ZfcRnAsydNwQyqJ4-g7gRe5K4-AzlR40qfa2FBNXFXq0OPoKZthNJSua0stkXKOrmNM7AyQkqwdPGEEbxxuZXH"
// 	uid   = "a1NNQdnfUgg96aLUzQA56MuYKHm2"
// )

func (Notification) SendFriendsPairsNotification(ctx context.Context, deviceToken string) error {

	msg := &fcm.Message{
		To: deviceToken,
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: "title",
			Body:  "body",
		},
	}
	client, err := fcm.NewClient("sample_api_key")
	if err != nil {
		return err
	}
	_, err = client.Send(msg)
	return err
}

func (n Notification) GenerateFriendsInvite(ctx context.Context, s *gorm.SQLCnter, pet_id string, recommend_id string) (m *messaging.MulticastMessage, err error) {
	if recommend_id == "" {
		return nil, errors.New("recommend_id is empty")
	}
	tokens, err := n.FindTokenByUID(ctx, s, recommend_id)
	if err != nil {
		log.Fatal(err)
		return nil, err
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
func (n Notification) FindTokenByUID(ctx context.Context, s *gorm.SQLCnter, uid string) ([]string, error) {
	data, err := s.FindUserDeviceID(ctx, uid)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ret := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		ret[i] = data[i].Device_id
	}
	return ret, nil
}

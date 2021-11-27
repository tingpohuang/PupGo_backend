package firebase

import (
	"context"
	"testing"

	"firebase.google.com/go/messaging"
	"github.com/stretchr/testify/assert"
)

func TestFireBaseMessage(t *testing.T) {
	assert := assert.New(t)
	f := GetApp()
	token := "eKWNaSb5FkHmrRiJ_4SUg-:APA91bGjnf2hGD_jUkpwUIMgp0z-UZJnFc-7YvdiJ-eX0KiGhevbh7bw9mhoT9XYXhdxFkpnI0h2SZSS_VPv_q6-esp7xld1fx7Wil-NKhbahjUgEk1sA3yc1h5INKpNitd2sEHua5Tt"
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Notification: &messaging.Notification{
			Title: "y",
			Body:  "pupgo",
		},
		Token: token,
	}
	err := f.SendNotification(context.Background(), message)
	assert.Nil(err)
}

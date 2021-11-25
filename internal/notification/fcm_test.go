package notification

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotification(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	v := Notification{}
	a := v.SendFriendsPairsNotification(ctx, "token")
	assert.Nil(a)
}

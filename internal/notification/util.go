package notification

import (
	"time"

	"github.com/google/uuid"
	"github.com/tingpo/pupgobackend/internal/gorm"
)

const ()

func NewNotification() *gorm.Notification {
	ret := &gorm.Notification{
		Notification_id: uuid.NewString(),
		Created_at:      time.Now(),
		Has_read:        false,
	}
	return ret
}

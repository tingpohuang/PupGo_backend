package notification

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tingpo/pupgobackend/internal/gorm"
)

func TestNotification(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
}

func TestSendEventJoinedMessage(t *testing.T) {
	// go test -timeout 30s -run ^TestSendEventJoinedMessage$ github.com/tingpo/pupgobackend/internal/notification
	ctx := context.Background()
	n := Notification{}
	mysqlConnector, err := gorm.GetConnectorFactory("mySQL")
	assert := assert.New(t)
	assert.Nil(err)
	if err != nil {
		panic(fmt.Errorf("Connect to DB failed: %w \n", err))
	}
	db := mysqlConnector.NewDBConnection()
	s := gorm.NewSQLCnter(db)
	n.SendEventJoinedMessage(ctx, gorm.Event_1_id, gorm.Pet_4_id, s)
}

func TestSendNewFriendMessage(t *testing.T) {
	// go test -timeout 30s -run ^TestSendNewFriendMessage$ github.com/tingpo/pupgobackend/internal/notification
	ctx := context.Background()
	n := Notification{}
	mysqlConnector, err := gorm.GetConnectorFactory("mySQL")
	if err != nil {
		panic(fmt.Errorf("Connect to DB failed: %w \n", err))
	}
	db := mysqlConnector.NewDBConnection()
	s := gorm.NewSQLCnter(db)
	n.SendNewFriendMessage(ctx, gorm.Pet_3_id, gorm.Pet_3_id, s)
}

func TestWriteNotification(t *testing.T) {
	// go test -timeout 30s -run ^TestWriteNotification$ github.com/tingpo/pupgobackend/internal/notification
	ctx := context.Background()
	mysqlConnector, err := gorm.GetConnectorFactory("mySQL")
	if err != nil {
		panic(fmt.Errorf("Connect to DB failed: %w \n", err))
	}
	assert := assert.New(t)
	db := mysqlConnector.NewDBConnection()
	s := gorm.NewSQLCnter(db)
	nMsg := NewNotification()
	uid, err := s.GetUserIdbyPetId(ctx, gorm.Pet_4_id)
	assert.Nil(err)
	nMsg.User_id = *uid
	// log.Fatal(nMsg.User_id)
	nMsg.Pet_id = gorm.Pet_4_id
	err = s.CreateNotification(ctx, nMsg)
	assert.Nil(err)
}

// "eNTUoWOXr0Z3oDzsNXiqKR:APA91bHbCAahTVQDx34Kif1XklEoPhtf0PEqaMXI0935xQ4bygZ046h-e9ApkrAkZbx6NIKFHNZpKfU1GHxIVMslUHVBUBPtChVnqIn8weNd_u6eNeMIiqv5sCcprw_J-G9TNsLZGeUH"

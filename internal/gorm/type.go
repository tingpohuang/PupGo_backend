package gorm

import (
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Name      string
	Cooldown  timestamp
	Create_at timestamp
	gender    int
	birthday  timestamp
}

type User_device struct {
	user_id   uuid.UUID `gorm:"primaryKey"`
	device_id uuid.UUID
}
type Pet_connection struct {
	id1 uuid.UUID `gorm:"primaryKey"`
	id2 uuid.UUID
}
type Pet_owner struct {
	user_id uuid.UUID `gorm:"primaryKey"`
	pet_id  uuid.UUID
}
type pet_recommend struct {
	id1    uuid.UUID `gorm:"primaryKey"`
	id2    uuid.UUID
	score  float64
	status int
}

type Event struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	Holder_Id       uuid.UUID
	Start_date      timestamp
	End_date        timestamp
	Image           string
	Limits_user_num int
	Limits_pet_num  int
	Description     string
}

type Pet struct {
	id           uuid.UUID `gorm:"primaryKey"`
	name         string
	image        string
	gender       int
	breed        string
	isCastration bool
	birthday     timestamp
}
type Event_participant struct {
	event_id       uuid.UUID `gorm:"primaryKey"`
	participant_id uuid.UUID
	pet_id         uuid.UUID
	status         int
}

//not exist time stamp
type timestamp struct {
}

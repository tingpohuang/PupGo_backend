package gorm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         []byte `gorm:"type:varbinary(16);primaryKey;"`
	Name       string
	Cooldown   time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Created_at time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	gender     int
	birthday   time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) error {

	uuid, err := uuid.New().MarshalBinary()
	u.Id = uuid
	u.Created_at = time.Now()
	return err
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
	User_id []byte `gorm:"primaryKey"`
	Pet_id  []byte
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
	Start_date      time.Time
	End_date        time.Time
	Image           string
	Limits_user_num int
	Limits_pet_num  int
	Description     string
}

type Pet struct {
	Id           []byte `gorm:"primaryKey"`
	Name         string
	Image        string
	Gender       int
	Breed        string
	IsCastration bool      `gorm:"column:isCastration"`
	Birthday     time.Time `gorm:"type:timestamp; default: NOW(); not null"`
}

func (p *Pet) BeforeCreate(db *gorm.DB) error {

	uuid, err := uuid.New().MarshalBinary()
	p.Id = uuid
	return err
}

type Event_participant struct {
	event_id       uuid.UUID `gorm:"primaryKey"`
	participant_id uuid.UUID
	pet_id         uuid.UUID
	status         int
}

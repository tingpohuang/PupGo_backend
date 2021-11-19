package gorm

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         string `gorm:"column:id",gorm:"type:VARCHAR(36);primaryKey;"`
	Name       string
	Cooldown   time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Created_at time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	gender     int
	birthday   time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) error {

	// uuid, err := uuid.New().MarshalBinary()
	// u.Id = uuid
	u.Created_at = time.Now()
	return nil
}

type User_device struct {
	user_id   string `gorm:"primaryKey"`
	device_id string
}
type Pet_connection struct {
	id1 string `gorm:"primaryKey"`
	id2 string
}
type Pet_owner struct {
	User_id string `gorm:"column:user_id", gorm:"primaryKey", gorm:"constraint:OnDelete:CASCADE"`
	Pet_id  string `gorm:"column:pet_id"`
}
type pet_recommend struct {
	id1    string `gorm:"primaryKey"`
	id2    string
	score  float64
	status int
}

type Event struct {
	Id              string `gorm:"primaryKey"`
	Holder_Id       string
	Start_date      time.Time
	End_date        time.Time
	Image           string
	Limits_user_num int
	Limits_pet_num  int
	Description     string
}

type Pet struct {
	Id           string `gorm:"primaryKey",gorm:"column:id"`
	Name         string
	Image        string
	Gender       int
	Breed        string
	IsCastration bool      `gorm:"column:isCastration"`
	Birthday     time.Time `gorm:"type:timestamp; default: NOW(); not null"`
}

// func (p *Pet) BeforeCreate(db *gorm.DB) error {

// 	uuid, err := uuid.New().MarshalBinary()
// 	// p.Id = uuid
// 	return err
// }

type Event_participant struct {
	event_id       string `gorm:"primaryKey"`
	participant_id string
	pet_id         string
	status         int
}

package gorm

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         string `gorm:"column:id",gorm:"type:VARCHAR(36);primaryKey"`
	Name       string
	Cooldown   time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Created_at time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Gender     int
	Birthday   time.Time `gorm:"type:timestamp; default: NOW(); not null"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {

	// uuid, err := uuid.New().MarshalBinary()
	// u.Id = uuid
	u.Created_at = time.Now()
	return nil
}

type User_device struct {
	User_id   string `gorm:"primaryKey"`
	Device_id string
}
type Pet_connection struct {
	id1 string `gorm:"type:VARCHAR(36);column:Id1;not null;default:null", gorm:"constraint:OnDelete:CASCADE"`
	id2 string `gorm:"type:VARCHAR(36);column:Id2;not null;default:null", gorm:"constraint:OnDelete:CASCADE"`
}

type Pet_owner struct {
	User_id string `gorm:"type:VARCHAR(36);column:user_id;not null;default:null", gorm:"constraint:OnDelete:CASCADE"`
	Pet_id  string `gorm:"type:VARCHAR(36);column:pet_id;not null;default:null", gorm:"primaryKey"`
}
type Pet_recommend struct {
	Id1    string `gorm:"type:VARCHAR(36);column:id1;not null;default:null", gorm:"constraint:OnDelete:CASCADE"`
	Id2    string `gorm:"type:VARCHAR(36);column:id2;not null;default:null", gorm:"constraint:OnDelete:CASCADE"`
	Score  float64
	Status int
}

type Event struct {
	Id             string `gorm:"type:VARCHAR(36);primaryKey"`
	Holder_Id      string
	Start_date     time.Time
	End_date       time.Time
	Image          string
	Limit_user_num int
	Limit_pet_num  int
	Description    string
}

type Pet struct {
	Id           string `gorm:"type:VARCHAR(36);column:id;primaryKey"`
	Name         string
	Image        string
	Gender       int
	Breed        string
	IsCastration bool      `gorm:"column:isCastration"`
	Birthday     time.Time `gorm:"type:timestamp; default: NOW(); not null"`
}

type Event_participant struct {
	event_id       string `gorm:"type:VARCHAR(36);OnDelete:CASCADE"`
	participant_id string `gorm:"type:VARCHAR(36);OnDelete:CASCADE"`
	pet_id         string `gorm:"type:VARCHAR(36);OnDelete:CASCADE"`
	status         int
}

package gorm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type User struct {
	Id         string `gorm:"column:id",gorm:"type:VARCHAR(36);primaryKey"`
	Name       string
	Cooldown   time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Created_at time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Gender     int
	Birthday   time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Email      string
}

type UserToken struct {
	User_id string `json:"user_id"`
	Token   string `json:"token"`
}

type Position struct {
	Lat  float64 `gorm:"column:lat"`
	Long float64 `gorm:"column:long"`
}

func (loc Position) GormDataType() string {
	return "geometry"
}

func (loc Position) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%v %v)", loc.Lat, loc.Long)},
	}
}

func (loc *Position) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	fmt.Printf("Scan value %v", v)
	bytes, ok := v.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", v))
	}

	var pos map[interface{}]interface{}
	json.Unmarshal(bytes, &pos)
	println(pos)

	return nil
}

type UserLocation struct {
	User_id  string
	Position Position
	Country  string
	State    string
	City     string
	Address  string
}

type EventLocation struct {
	Event_id string
	Position Position
	Country  string
	State    string
	City     string
	Address  string
}

func (u *User) BeforeCreate(db *gorm.DB) error {
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
	Event_id       string `gorm:"type:VARCHAR(36);OnDelete:CASCADE"`
	Participant_id string `gorm:"type:VARCHAR(36);OnDelete:CASCADE"`
	Pet_id         string `gorm:"type:VARCHAR(36);OnDelete:CASCADE"`
	Status         int
}

type Notification struct {
	Notification_id   string    `gorm:"type:VARCHAR(36);not null"`
	Notification_type int       `gorm:"type:int"`
	User_id           string    `gorm:"type:VARCHAR(36);OnDelete:CASCADE;not null"`
	Pet_id            string    `gorm:"type:VARCHAR(36)"`
	Event_id          string    `gorm:"type:VARCHAR(36)"`
	Created_at        time.Time `gorm:"type:timestamp; default: NOW(); not null"`
	Payload           string    `gorm:"type:VARCHAR(1024)"`
	Has_read          bool      `gorm:"type:Boolean"`
}

package gorm

type User struct {
	Id        []byte `gorm:"primaryKey"`
	Name      string
	Cooldown  timestamp
	Create_at timestamp
	gender    int
	birthday  timestamp
}

type User_device struct {
	user_id   []byte `gorm:"primaryKey"`
	device_id []byte
}
type Pet_connection struct {
	id1 []byte `gorm:"primaryKey"`
	id2 []byte
}
type Pet_owner struct {
	user_id []byte `gorm:"primaryKey"`
	pet_id  []byte
}
type pet_recommend struct {
	id1    []byte `gorm:"primaryKey"`
	id2    []byte
	score  float64
	status int
}

type Event struct {
	Id              []byte `gorm:"primaryKey"`
	Holder_Id       []byte
	Start_date      timestamp
	End_date        timestamp
	Image           string
	Limits_user_num int
	Limits_pet_num  int
	Description     string
}

type Pet struct {
	id           []byte `gorm:"primaryKey"`
	name         string
	image        string
	gender       int
	breed        string
	isCastration bool
	birthday     timestamp
}
type Event_participant struct {
	event_id       []byte `gorm:"primaryKey"`
	participant_id []byte
	pet_id         []byte
	status         int
}

//not exist time stamp
type timestamp struct {
}

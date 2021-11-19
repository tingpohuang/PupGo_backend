package gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	u1id = []byte("abcd")
	u2id = []byte("abce")
	u3id = []byte("abcf")
	p1id = []byte("abcd")
	p2id = []byte("bbcd")
	p3id = []byte("cbcd")
	p4id = []byte("dbcd")
	e1id = []byte("aaaa")
	e2id = []byte("bbbb")
	e3id = []byte("cccc")
	u1   = User{
		Id:     u1id,
		Name:   "User_1",
		gender: 1,
	}
	u2 = User{
		Id:     u2id,
		Name:   "User_2",
		gender: 1,
	}
	u3 = User{
		Id:     u3id,
		Name:   "User_1",
		gender: 0,
	}
	po1 = Pet_owner{
		user_id: u1id,
		pet_id:  p1id,
	}
	po2 = Pet_owner{
		user_id: u1id,
		pet_id:  p2id,
	}
	po3 = Pet_owner{
		user_id: u2id,
		pet_id:  p3id,
	}
	po4 = Pet_owner{
		user_id: u2id,
		pet_id:  p4id,
	}
	pc1 = Pet_connection{
		id1: p1id,
		id2: p3id,
	}
	pc2 = Pet_connection{
		id1: p1id,
		id2: p4id,
	}
	pr1 = pet_recommend{
		id1:    p2id,
		id2:    p3id,
		score:  0.03,
		status: 0,
	}
	pr2 = pet_recommend{
		id1:    p2id,
		id2:    p4id,
		score:  0.05,
		status: 0,
	}
	e1 = Event{
		Id:         e1id,
		Holder_Id:  u1id,
		Start_date: timestamp{},
		End_date:   timestamp{},
		Image: "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/holiday-logo-202111?wid=71&hei=87&fmt=jpeg&qlt=95&.v=1636070054000"
	}
	ep1 = Event_participant{
		event_id : e1id,
		participant_id: u1id,
		pet_id: p2id,
		status: 1,
	}
	ep1 = Event_participant{
		event_id : e1id,
		participant_id: u2id,
		pet_id: p3id,
		status: 1,
	}
	ep1 = Event_participant{
		event_id : e1id,
		participant_id: u2id,
		pet_id: p4id,
		status: 1,
	}
)

type gormTestor struct {
	gdb *gorm.DB
}

func TestInputValueToDatabse(t *testing.T) {
	// timestamp := time.Now()
	connector, err := GetConnectorFactory("mySQL")
	db := connector.NewDBConnection()
	assert := assert.New(t)
	assert.Nil(t, err)
	assert.NotNil(t, db)
	user := User{
		Name: "test",
		// Cooldown: time.Now(),
		gender: 1,
	}
	result := db.Create(&user) // pass pointer of data to Create
	assert.Nil(result.Error)
}

func InputUser(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Create(&u1) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Create(&u2) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Create(&u3) // pass pointer of data to Create
	assert.Nil(result.Error)
}

func InputPet(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
}

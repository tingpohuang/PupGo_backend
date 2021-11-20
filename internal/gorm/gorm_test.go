package gorm

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	u1id = uuid.NewString()
	u2id = uuid.NewString()
	u3id = uuid.NewString()
	p1id = uuid.NewString()
	p2id = uuid.NewString()
	p3id = uuid.NewString()
	p4id = uuid.NewString()
	e1id = uuid.NewString()
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
	p1 = Pet{
		Id: p1id,
	}
	p2 = Pet{
		Id: p2id,
	}
	p3 = Pet{
		Id: p3id,
	}
	p4 = Pet{
		Id: p4id,
	}
	po1 = Pet_owner{
		User_id: u1id,
		Pet_id:  p1id,
	}
	po2 = Pet_owner{
		User_id: u1id,
		Pet_id:  p2id,
	}
	po3 = Pet_owner{
		User_id: u2id,
		Pet_id:  p3id,
	}
	po4 = Pet_owner{
		User_id: u2id,
		Pet_id:  p4id,
	}
	pc1 = Pet_connection{
		id1: p1id,
		id2: p3id,
	}
	pc2 = Pet_connection{
		id1: p1id,
		id2: p4id,
	}
	pr1 = Pet_recommend{
		id1:    p2id,
		id2:    p3id,
		score:  0.03,
		status: 0,
	}
	pr2 = Pet_recommend{
		id1:    p2id,
		id2:    p4id,
		score:  0.05,
		status: 0,
	}
	e1 = Event{
		Id:             e1id,
		Holder_Id:      u1id,
		Start_date:     time.Now(),
		End_date:       time.Now(),
		Image:          "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/holiday-logo-202111?wid=71&hei=87&fmt=jpeg&qlt=95&.v=1636070054000",
		Limit_user_num: 0,
		Limit_pet_num:  0,
		Description:    "",
	}
	ep1 = Event_participant{
		event_id:       e1id,
		participant_id: u1id,
		pet_id:         p2id,
		status:         1,
	}
	ep2 = Event_participant{
		event_id:       e1id,
		participant_id: u2id,
		pet_id:         p3id,
		status:         1,
	}
	ep3 = Event_participant{
		event_id:       e1id,
		participant_id: u2id,
		pet_id:         p4id,
		status:         1,
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
	g := gormTestor{gdb}
	assert.Nil(err)
	assert.NotNil(db)

	InputUser(t, &g)
	InputPet(t, &g)
	InputPetOwner(t, &g)
	InputPetConnect(t, &g)
	InputPetRecommend(t, &g)
	InputEvent(t, &g)
	InputEventParticipant(t, &g)
}

func InputUser(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("users").Create(&u1) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("users").Create(&u2) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("users").Create(&u3) // pass pointer of data to Create
	assert.Nil(result.Error)
}

func InputPet(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("pet").Create(&p1) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("pet").Create(&p2) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("pet").Create(&p3) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("pet").Create(&p4) // pass pointer of data to Create
	assert.Nil(result.Error)
}

func InputPetOwner(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("petowner").Create(&po1)
	assert.Nil(result.Error)
	result = g.gdb.Table("petowner").Create(&po2)
	assert.Nil(result.Error)
	result = g.gdb.Table("petowner").Create(&po3)
	assert.Nil(result.Error)
	result = g.gdb.Table("petowner").Create(&po4)
	assert.Nil(result.Error)
}

func InputPetConnect(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Exec("INSERT INTO pet_connection VALUES (?, ?)", pc1.id1, pc1.id2)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO pet_connection VALUES (?, ?)", pc2.id1, pc2.id2)
	assert.Nil(result.Error)
}

func InputPetRecommend(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Exec("INSERT INTO pet_recommend VALUES (?, ?, ?, ?)", pr1.id1, pr1.id2, pr1.score, pr1.status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO pet_recommend VALUES (?, ?, ?, ?)", pr2.id1, pr2.id2, pr2.score, pr2.status)
	assert.Nil(result.Error)
}

func InputEvent(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("event").Create(&e1)
	assert.Nil(result.Error)
}

func InputEventParticipant(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep1.event_id, ep1.participant_id, ep1.pet_id, ep1.status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep2.event_id, ep2.participant_id, ep2.pet_id, ep2.status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep3.event_id, ep3.participant_id, ep3.pet_id, ep3.status)
	assert.Nil(result.Error)

}

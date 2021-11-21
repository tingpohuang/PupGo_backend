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
		Gender: 1,
	}
	u2 = User{
		Id:     u2id,
		Name:   "User_2",
		Gender: 1,
	}
	u3 = User{
		Id:     u3id,
		Name:   "User_1",
		Gender: 0,
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
		Id1:    p2id,
		Id2:    p3id,
		Score:  0.03,
		Status: 0,
	}
	pr2 = Pet_recommend{
		Id1:    p2id,
		Id2:    p4id,
		Score:  0.05,
		Status: 0,
	}
	e1 = Event{
		Id:             e1id,
		Holder_Id:      p1id,
		Start_date:     time.Now(),
		End_date:       time.Now(),
		Image:          "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/holiday-logo-202111?wid=71&hei=87&fmt=jpeg&qlt=95&.v=1636070054000",
		Limit_user_num: 0,
		Limit_pet_num:  0,
		Description:    "",
	}
	ep1 = Event_participant{
		Event_id:       e1id,
		Participant_id: u1id,
		Pet_id:         p2id,
		Status:         1,
	}
	ep2 = Event_participant{
		Event_id:       e1id,
		Participant_id: u2id,
		Pet_id:         p3id,
		Status:         1,
	}
	ep3 = Event_participant{
		Event_id:       e1id,
		Participant_id: u2id,
		Pet_id:         p4id,
		Status:         1,
	}

	uLoc1 = UserLocation{
		User_id: u1id,
		Position: Location{
			Lat:  1.23,
			Long: 4.56,
		},
		Country: "USA",
		State:   "CA",
		City:    "Los Angeles",
		Address: "1878 Greenfield Avenue",
	}
	uLoc2 = UserLocation{
		User_id: u2id,
		Position: Location{
			Lat:  1.23,
			Long: 4.56,
		},
		Country: "USA",
		State:   "CA",
		City:    "Los Angeles",
		Address: "1878 Greenfield Avenue",
	}
	uLoc3 = UserLocation{
		User_id: u3id,
		Position: Location{
			Lat:  1.232,
			Long: 4.56,
		},
		Country: "USA",
		State:   "CA",
		City:    "Los Angeles",
		Address: "1878 Greenfield Avenue",
	}

	e1Loc = EventLocation{
		Event_id: e1id,
		Position: Location{
			Lat:  13.33,
			Long: 24.432,
		},
		Country: "USA",
		State:   "CA",
		City:    "Los Angeles",
		Address: "fwfewfewfew",
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
	InputUserLocation(t, &g)
	InputPet(t, &g)
	InputPetOwner(t, &g)
	InputPetConnect(t, &g)
	InputPetRecommend(t, &g)
	InputEvent(t, &g)
	InputEventParticipant(t, &g)
	InputEventLocation(t, &g)

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
	result := g.gdb.Exec("INSERT INTO pet_recommend VALUES (?, ?, ?, ?)", pr1.Id1, pr1.Id2, pr1.Score, pr1.Status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO pet_recommend VALUES (?, ?, ?, ?)", pr2.Id1, pr2.Id2, pr2.Score, pr2.Status)
	assert.Nil(result.Error)
}

func InputEvent(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("event").Create(&e1)
	assert.Nil(result.Error)
}

func InputEventParticipant(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep1.Event_id, ep1.Participant_id, ep1.Pet_id, ep1.Status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep2.Event_id, ep2.Participant_id, ep2.Pet_id, ep2.Status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep3.Event_id, ep3.Participant_id, ep3.Pet_id, ep3.Status)
	assert.Nil(result.Error)

}

func InputUserLocation(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("user_location").Create(&uLoc1)
	assert.Nil(result.Error)
	result = g.gdb.Table("user_location").Create(&uLoc2)
	assert.Nil(result.Error)
	result = g.gdb.Table("user_location").Create(&uLoc3)
	assert.Nil(result.Error)
}

func InputEventLocation(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("event_location").Create(&e1Loc)
	assert.Nil(result.Error)
}

package gorm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	u1 = User{
		Id:       User_1_id,
		Name:     "User_1",
		Gender:   1,
		Birthday: time.Now(),
	}
	u2 = User{
		Id:       User_2_id,
		Name:     "User_2",
		Gender:   1,
		Birthday: time.Now(),
	}
	u3 = User{
		Id:       User_3_id,
		Name:     "User_3",
		Gender:   0,
		Birthday: time.Now(),
	}
	p1 = Pet{
		Id: Pet_1_id,
	}
	p2 = Pet{
		Id: Pet_2_id,
	}
	p3 = Pet{
		Id: Pet_3_id,
	}
	p4 = Pet{
		Id: Pet_4_id,
	}
	po1 = Pet_owner{
		User_id: User_1_id,
		Pet_id:  Pet_1_id,
	}
	po2 = Pet_owner{
		User_id: User_1_id,
		Pet_id:  Pet_2_id,
	}
	po3 = Pet_owner{
		User_id: User_2_id,
		Pet_id:  Pet_3_id,
	}
	po4 = Pet_owner{
		User_id: User_2_id,
		Pet_id:  Pet_4_id,
	}
	pc1 = Pet_connection{
		id1: Pet_1_id,
		id2: Pet_3_id,
	}
	pc2 = Pet_connection{
		id1: Pet_1_id,
		id2: Pet_4_id,
	}
	pr1 = Pet_recommend{
		Id1:    Pet_2_id,
		Id2:    Pet_3_id,
		Score:  0.03,
		Status: 0,
	}
	pr2 = Pet_recommend{
		Id1:    Pet_2_id,
		Id2:    Pet_4_id,
		Score:  0.05,
		Status: 0,
	}
	e1 = Event{
		Id:             Event_1_id,
		Holder_Id:      Pet_1_id,
		Start_date:     time.Now(),
		End_date:       time.Now(),
		Image:          "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/holiday-logo-202111?wid=71&hei=87&fmt=jpeg&qlt=95&.v=1636070054000",
		Limit_user_num: 0,
		Limit_pet_num:  0,
		Description:    "",
	}
	ep1 = Event_participant{
		Event_id:       Event_1_id,
		Participant_id: User_1_id,
		Pet_id:         Pet_2_id,
		Status:         1,
	}
	ep2 = Event_participant{
		Event_id:       Event_1_id,
		Participant_id: User_2_id,
		Pet_id:         Pet_3_id,
		Status:         1,
	}
	// ep3 = Event_participant{
	// 	Event_id:       Event_1_id,
	// 	Participant_id: User_2_id,
	// 	Pet_id:         Pet_4_id,
	// 	Status:         1,
	// }

	e1Loc = EventLocation{
		Event_id: Event_1_id,
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
	// command: "go test -timeout 60s -run ^TestInputValueToDatabse$ github.com/tingpo/pupgobackend/internal/gorm"
	// please no add in db connection lol.
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
	if pr1.Id1 > pr1.Id2 {
		pr1.Id1, pr1.Id2 = pr1.Id2, pr1.Id1
	}
	if pr2.Id1 > pr2.Id2 {
		pr2.Id1, pr2.Id2 = pr2.Id2, pr2.Id1
	}
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
	// result = g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep3.Event_id, ep3.Participant_id, ep3.Pet_id, ep3.Status)
	// assert.Nil(result.Error)

}

func InputUserLocation(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("user_location").Create(&UserLocation1)
	assert.Nil(result.Error)
	result = g.gdb.Table("user_location").Create(&UserLocation2)
	assert.Nil(result.Error)
	result = g.gdb.Table("user_location").Create(&UserLocation3)
	assert.Nil(result.Error)
}

func InputEventLocation(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("event_location").Create(&e1Loc)
	assert.Nil(result.Error)
}

package gorm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// bf448152-bf2c-421e-af4c-ae458737da0e	eKWNaSb5FkHmrRiJ_4SUg-:APA91bGjnf2hGD_jUkpwUIMgp0z-UZJnFc-7YvdiJ-eX0KiGhevbh7bw9mhoT9XYXhdxFkpnI0h2SZSS_VPv_q6-esp7xld1fx7Wil-NKhbahjUgEk1sA3yc1h5INKpNitd2sEHua5Tt
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
	u4 = User{
		Id:       User_4_id,
		Name:     "User_4",
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

	e2 = Event{
		Id:             Event_2_id,
		Holder_Id:      Pet_1_id,
		Start_date:     time.Now(),
		End_date:       time.Now(),
		Image:          "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/holiday-logo-202111?wid=71&hei=87&fmt=jpeg&qlt=95&.v=1636070054000",
		Limit_user_num: 0,
		Limit_pet_num:  0,
		Description:    "",
	}
	e3 = Event{
		Id:             Event_3_id,
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

	e1Loc = EventLocation{
		Event_id:  Event_1_id,
		Latitude:  101,
		Longitude: 101,
		Country:   "USA",
		State:     "CA",
		City:      "Los Angeles",
		Address:   "fwfewfewfew",
	}

	e2Loc = EventLocation{
		Event_id:  Event_2_id,
		Latitude:  301,
		Longitude: 301,
		Country:   "USA",
		State:     "CA",
		City:      "Los Angeles",
		Address:   "fwfewfewfew",
	}

	e3Loc = EventLocation{
		Event_id:  Event_3_id,
		Latitude:  201,
		Longitude: 201,
		Country:   "USA",
		State:     "CA",
		City:      "Los Angeles",
		Address:   "fwfewfewfew",
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
	InputUserDevice(t, &g)

}

func InputUser(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("users").Create(&u1) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("users").Create(&u2) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("users").Create(&u3) // pass pointer of data to Create
	assert.Nil(result.Error)
	result = g.gdb.Table("users").Create(&u4) // pass pointer of data to Create
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
	for i := 0; i < Petsize; i++ {
		result = g.gdb.Table("pet").Create(&Pet{
			Id:           Pet_ids[i],
			Name:         Pet_names[i],
			Image:        Pet_imgs[i],
			Gender:       Pet_genders[i],
			Breed:        Pet_breeds[i],
			IsCastration: Pet_isCastrations[i],
			Birthday:     Pet_Birthdays[i],
		})
		assert.Nil(result.Error)
	}
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
	for i := 0; i < Petsize; i++ {
		result = g.gdb.Table("petowner").Create(&Pet_owner{
			User_id: User_4_id,
			Pet_id:  Pet_ids[i],
		})
		assert.Nil(result.Error)
	}
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
	v := 0.1
	result := g.gdb.Exec("INSERT INTO pet_recommend VALUES (?, ?, ?, ?)", pr1.Id1, pr1.Id2, pr1.Score, pr1.Status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO pet_recommend VALUES (?, ?, ?, ?)", pr2.Id1, pr2.Id2, pr2.Score, pr2.Status)
	for i := 1; i < Petsize; i++ {
		pid1 := Pet_ids[0]
		pid2 := Pet_ids[i]
		if pid1 > pid2 {
			pid1, pid2 = pid2, pid1
		}
		result := g.gdb.Exec("INSERT INTO pet_recommend VALUES (?, ?, ?,?)", pid1, pid2, float64(i)*v, 0)
		assert.Nil(result.Error)
	}
	assert.Nil(result.Error)
}

func InputEvent(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Table("event").Create(&e1)
	assert.Nil(result.Error)
	result = g.gdb.Table("event").Create(&e2)
	assert.Nil(result.Error)
	result = g.gdb.Table("event").Create(&e3)
	assert.Nil(result.Error)
}

func InputEventParticipant(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep1.Event_id, ep1.Participant_id, ep1.Pet_id, ep1.Status)
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO event_participant VALUES (?, ?, ?, ?)", ep2.Event_id, ep2.Participant_id, ep2.Pet_id, ep2.Status)
	assert.Nil(result.Error)
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
	result = g.gdb.Table("event_location").Create(&e2Loc)
	assert.Nil(result.Error)
	result = g.gdb.Table("event_location").Create(&e3Loc)
	assert.Nil(result.Error)
}

func InputUserDevice(t *testing.T, g *gormTestor) {
	assert := assert.New(t)
	result := g.gdb.Exec("INSERT INTO user_device VALUES (?, ?)", User_1_id, "eKWNaSb5FkHmrRiJ_4SUg-:APA91bGjnf2hGD_jUkpwUIMgp0z-UZJnFc-7YvdiJ-eX0KiGhevbh7bw9mhoT9XYXhdxFkpnI0h2SZSS_VPv_q6-esp7xld1fx7Wil-NKhbahjUgEk1sA3yc1h5INKpNitd2sEHua5Tt")
	assert.Nil(result.Error)
	result = g.gdb.Exec("INSERT INTO user_device VALUES (?, ?)", User_2_id, "cyu9JfHduEfOtr6Y0Wyqz2:APA91bEC_NEGcK71QB7YclvMMXOv0DgzK7MJeCc99disvVQtrNkiPpunpjo1ILEoo1eEyBKi28ChFEMcX-gz0AU_6dbgTCJsXYDfHwRBDb89E_DSMRzcgxon7f_o3E33GYwq8oj_VG4b")
	assert.Nil(result.Error)
}

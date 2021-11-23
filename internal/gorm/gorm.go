package gorm

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewSQLCnter(gdb *gorm.DB) *SQLCnter {
	return &SQLCnter{
		gdb,
	}
}

type SQLCnter struct {
	gdb *gorm.DB
}

func (*SQLCnter) CreateUser() {
	id1 := uuid.NewString()
	id2 := uuid.NewString()
	// Id1 := []byte("abcd")
	// Id1, _ := uuid.New().MarshalBinary()
	// Id2, _ := uuid.New().MarshalBinary()
	user := User{
		Id:   id1,
		Name: "test",
	}
	pet := Pet{
		Id: id2,
	}
	petower := Pet_owner{
		User_id: id1,
		Pet_id:  id2,
	}

	gdb.Table("users").Create(&user)
	gdb.Table("pet").Create(&pet)
	gdb.Table("petowner").Create(&petower)

}

func (s *SQLCnter) findEventByUId(ctx context.Context, uid string) (event []string) {
	var eventParticipant []Event_participant
	eventMap := make(map[string]bool)

	s.gdb.Table("event_participant").Where("Participant_id = ?", uid).Find(&eventParticipant)
	for i := 0; i < len(eventParticipant); i++ {
		cur := eventParticipant[i]
		_, ok := eventMap[cur.Event_id]
		if !ok {
			event = append(event, cur.Event_id)
			eventMap[cur.Event_id] = true

		}

	}
	return event
}

func (s *SQLCnter) findEventLocationByIdList(ctx context.Context, id []string) (eventLocations []EventLocation) {
	(*s.gdb).Table("event_location").Where("event_id IN ? ", id).Find(&eventLocations)
	return eventLocations
}

func (s *SQLCnter) findEventByIdList(ctx context.Context, id []string) (events []Event) {
	(*s.gdb).Table("event").Where("id IN ? ", id).Find(&events)
	return events
}

func (s *SQLCnter) findEventParticipantById(ctx context.Context, id string) (pets []string, participants []string) {
	var eventParticipant []Event_participant
	s.gdb.Table("event_participant").Where("Event_id = ?", id).Find(&eventParticipant)
	for i := 0; i < len(eventParticipant); i++ {
		cur := eventParticipant[i]
		pets = append(pets, cur.Pet_id)
		participants = append(participants, cur.Participant_id)
	}

	return pets, participants
}

func (s *SQLCnter) findPetsByUId(ctx context.Context, uid string) (pets []string) {
	var pet_owners []Pet_owner

	s.gdb.Table("petowner").Where("user_id = ?", uid).Find(&pet_owners)
	for i := 0; i < len(pet_owners); i++ {
		cur := pet_owners[i]
		pets = append(pets, cur.Pet_id)
	}
	return pets
}

func (s *SQLCnter) findUserLocationByPetsIdList(ctx context.Context, pid []string) (userLocations []UserLocation) {
	err := (*s.gdb).Table("petowner").Select("user_location.user_id as user_id,user_location.position as position, user_location.country as country, user_location.state as state, user_location.address as address,user_location.city as city").Joins("left join user_location on user_location.user_id = petowner.user_id where pet_id in ?", pid).Scan(&userLocations)

	if err.Error != nil {
		fmt.Println(err.Error)
	}
	// err
	return userLocations
	// Where("pet_id in ?", pid)
}

func (s *SQLCnter) findUserLocationByIdList(ctx context.Context, uid []string) (userLocations []UserLocation, err error) {

	(*s.gdb).Table("user_location").Where("user_id IN ? ", uid).Find(&userLocations)
	return userLocations, nil
}

func (s *SQLCnter) findUserByIdList(ctx context.Context, uid []string) (users []User) {
	(*s.gdb).Table("users").Where("id IN ? ", uid).Find(&users)
	return users
}

func (s *SQLCnter) findPetByIdList(ctx context.Context, pid []string) (pets []Pet) {
	(*s.gdb).Table("pet").Where("id IN ? ", pid).Find(&pets)
	return pets
}

func (s *SQLCnter) findPetByOwner(ctx context.Context, uid string) (pets []Pet) {
	s.gdb.Joins("Company", s.gdb.Where(&Pet_owner{User_id: uid})).Find(&pets)
	return pets
}

func (s *SQLCnter) findPetRecommend(ctx context.Context, pid string) (petRecommend []Pet_recommend) {
	s.gdb.Table("pet_recommend").Where("Id1 = ? OR Id2 = ?", pid, pid).Find(&petRecommend)
	return petRecommend
}
func (s *SQLCnter) FindPetRecommendByID(ctx context.Context, pid1 string, pid2 string) (petRecommend Pet_recommend, err error) {
	if pid1 > pid2 {
		pid1, pid2 = pid2, pid1
	}
	s.gdb.Table("pet_recommend").Find(&petRecommend, "id1 = ? AND id2 = ?", pid1, pid2)
	return petRecommend, nil

}
func (s *SQLCnter) UpdatePetRecommendByID(ctx context.Context, p Pet_recommend) (err error) {

	if p.Id1 > p.Id2 {
		p.Id1, p.Id2 = p.Id2, p.Id1
	}
	s.gdb.Table("pet_recommend").Find(&p, "id1 = ? AND id2 = ?", p.Id1, p.Id2).Updates(&p)
	return nil

}

func (s *SQLCnter) CreatePets(ctx context.Context, pet Pet) error {
	result := s.gdb.Table("pet").Create(&pet)
	return result.Error
}
func (s *SQLCnter) CreateUserPetRelation(ctx context.Context, uid string, pid string) error {
	result := s.gdb.Table("petowner").Create(&Pet_owner{
		User_id: uid,
		Pet_id:  pid,
	})
	return result.Error
}
func (s *SQLCnter) CreateEvents(ctx context.Context, e Event) error {
	result := s.gdb.Table("event").Create(&e)
	return result.Error
}
func (s *SQLCnter) FindEvents(ctx context.Context, pid string, eid string) (e *Event_participant, err error) {
	result := s.gdb.Table("event_participant").Where("event_id = ? AND pet_id = ? ", eid, pid).Find(&e)
	return e, result.Error
}
func (s *SQLCnter) DeletePet(ctx context.Context, pid string) error {
	result := s.gdb.Table("pet").Delete(&Pet{}, pid)
	return result.Error
}

func (s *SQLCnter) CreatePetConnection(ctx context.Context, pid1 string, pid2 string) error {
	if pid1 > pid2 {
		pid1, pid2 = pid2, pid1
	}
	result := s.gdb.Table("pet_connection").Create(&Pet_connection{
		id1: pid1,
		id2: pid2,
	})
	return result.Error
}

func (s *SQLCnter) FindPetById(ctx context.Context, pid string) (pets []Pet) {
	(*s.gdb).Table("pet").Where("id IN ? ", pid).Find(&pets)
	return pets
}

func (s *SQLCnter) UpdatePet(ctx context.Context, pet Pet) error {
	result := (*s.gdb).Table("pet").Model(&pet).Updates(&pet)
	return result.Error
}
func (s *SQLCnter) DeleteFriend(ctx context.Context, id1 string, id2 string) error {
	result := s.gdb.Table("pet_connection").Delete(&Pet_connection{id1: id1, id2: id2})
	return result.Error
}
func (s *SQLCnter) GetUserIdbyPetId(ctx context.Context, pid string) (*string, error) {
	p := Pet_owner{}
	result := s.gdb.Table("pet_owner").First(&p, "Pet_id = ?", pid)
	return &p.User_id, result.Error
}
func (s *SQLCnter) CreateParticipants(ctx context.Context, e Event_participant) error {
	result := s.gdb.Table("event_participant").Create(e)
	return result.Error
}

// func (s *SQLCnter) CreateParticipants(ctx context.Context, e Event_participant) error {

// }
/*
func (s *SQLCnter) findUsersByEvents(ctx context.Context) (user uuid.UUID) {

// }

// func (s *SQLCnter) findPetsByEvents(ctx context.Context) (pets []uuid.UUID) {

// }

func (s *SQLCnter) findEventsByUser(ctx context.Context) {

}

func (s *SQLCnter) findEventsByPets(ctx context.Context) {

}

func (S *SQLCnter) findEventsNearBy(ctx context.Context) {

}

func (S *SQLCnter) findConnection(ctx context.Context) {

}

func (S *SQLCnter) findPetRecommend(ctx context.Context) {

}

func (S *SQLCnter) removePets(ctx context.Context) {
}

*/

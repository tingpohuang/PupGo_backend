package gorm

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

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

func (*SQLCnter) CreateUser(name string, email string) (user User) {
	id1 := uuid.NewString()
	user = User{
		Id:         id1,
		Name:       name,
		Cooldown:   time.Time{},
		Created_at: time.Time{},
		Gender:     0,
		Birthday:   time.Time{},
		Email:      email,
	}

	gdb.Table("users").Create(&user)
	return user
}

func (*SQLCnter) CreateUserToken(uid string, token string) (err error) {
	userToken := UserToken{
		User_id: uid,
		Token:   token,
	}

	result := gdb.Table("user_token").Create(&userToken)
	return result.Error
}

func (s *SQLCnter) FindTokenByID(uid string) (userToken UserToken, err error) {

	err = s.gdb.Table("user_token").Where("user_id = ?", uid).First(&userToken).Error
	return userToken, err

}
func (s *SQLCnter) UpdateTokenByID(uid string, userToken UserToken) (err error) {

	result := s.gdb.Table("user_token").Find(&userToken, "user_id = ?", uid).Updates(&userToken)
	return result.Error

}

func (s *SQLCnter) FindUserByEmail(email string) (user User, err error) {
	err = (*s.gdb).Table("users").Where("email = ? ", email).First(&user).Error
	return user, err
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

func (s *SQLCnter) FindRecommendEventByUId(ctx context.Context, uid string) (event []string) {

	var eventRaw []Event
	s.gdb.Table("event").Find(&eventRaw)

	userLocations, err := s.findUserLocationByIdList(ctx, []string{uid})
	if err != nil {
		return nil
	}

	userLocation := userLocations[0]
	distanceMap := make(map[float64][]string)
	for i := 0; i < len(eventRaw); i++ {
		cur := eventRaw[i]
		eventLocations := s.findEventLocationByIdList(ctx, []string{eventRaw[i].Id})
		eventLocation := eventLocations[0]
		distance := math.Sqrt(math.Pow(userLocation.Latitude-eventLocation.Latitude, 2) + math.Pow(userLocation.Longitude-eventLocation.Longitude, 2))
		distanceMap[distance] = append(distanceMap[distance], cur.Id)
	}

	keys := make([]float64, 0, len(distanceMap))
	for k := range distanceMap {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(distanceMap[keys[i]]); j++ {
			event = append(event, distanceMap[keys[i]][j])
		}
	}
	return event
}

func (s *SQLCnter) findEventLocationByIdList(ctx context.Context, id []string) (eventLocations []EventLocation) {
	for i := 0; i < len(id); i++ {
		var eventLocation EventLocation
		(*s.gdb).Table("event_location").Where("event_id = ? ", id[i]).Find(&eventLocation)
		eventLocations = append(eventLocations, eventLocation)

	}

	return eventLocations
}

func (s *SQLCnter) findEventByIdList(ctx context.Context, id []string) (events []Event) {

	for i := 0; i < len(id); i++ {
		var event Event
		(*s.gdb).Table("event").Where("id = ? ", id[i]).Find(&event)
		events = append(events, event)
	}

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
	err := (*s.gdb).Table("petowner").Select("user_location.user_id as user_id,user_location.latitude as latitude, user_location.longitude as longitude, user_location.country as country, user_location.state as state, user_location.address as address,user_location.city as city").Joins("left join user_location on user_location.user_id = petowner.user_id where pet_id in ?", pid).Scan(&userLocations)

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
	res := s.gdb.Table("pet_recommend").Where("id1 = ? AND id2 = ?", p.Id1, p.Id2).Update("status", p.Status)
	return res.Error
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
func (s *SQLCnter) FindEventsParticipantByPetID(ctx context.Context, pid string, eid string) (e *Event_participant, err error) {
	result := s.gdb.Table("event_participant").Where("event_id = ? AND pet_id = ?", eid, pid).First(&e)
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
	result := s.gdb.Exec("INSERT INTO pet_connection VALUES (?, ?)", pid1, pid2)
	// result := s.gdb.Table("pet_connection").Create(&p)
	return result.Error
}

func (s *SQLCnter) FindPetById(ctx context.Context, pid string) (pets []Pet) {
	(*s.gdb).Table("pet").Where("id IN ? ", pid).Find(&pets)
	return pets
}
func (s *SQLCnter) FindHolderIdByEventId(ctx context.Context, eid string) (str string) {
	(*s.gdb).Table("event").Select("holder_id").Where("id = ? ", eid).Find(&str)
	return str
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
	result := s.gdb.Table("petowner").First(&p, "Pet_id = ?", pid)
	return &p.User_id, result.Error
}
func (s *SQLCnter) CreateParticipants(ctx context.Context, e Event_participant) error {
	result := s.gdb.Table("event_participant").Create(e)
	return result.Error
}
func (s *SQLCnter) UpdateParticipants(ctx context.Context, e Event_participant) error {
	result := s.gdb.Table("event_participant").Where("event_id = ? AND pet_id = ?", e.Event_id, e.Pet_id).Updates(&e)
	return result.Error
}

func (s *SQLCnter) FindDeviceByUserID(ctx context.Context, uid string) (devices []string, err error) {
	result := s.gdb.Table("user_device").Select("device_id").Where("user_id = ?", uid).Scan(devices)
	return devices, result.Error
}
func (s *SQLCnter) FindDeviceByUserIDs(ctx context.Context, uids []string) (devices []string, err error) {
	result := s.gdb.Table("user_device").Select("device_id").Where("user_id IN ?", uids).Scan(devices)
	return devices, result.Error
}

func (s *SQLCnter) FindDeviceByPetID(ctx context.Context, pid string) (devices []string, err error) {
	result := s.gdb.Table("user_device").Select("device_id").Joins("left join petowner on user_device.user_id = petowner.user_id where petowner.pet_id = ?", pid).Scan(&devices)
	return devices, result.Error
}

func (s *SQLCnter) FindDeviceByPetIDs(ctx context.Context, pid []string) (devices []string, err error) {
	result := s.gdb.Table("user_device").Select("device_id").Joins("left join petowner on user_device.user_id = petowner.user_id where petowner.pet_id IN ?", pid).Scan(&devices)
	return devices, result.Error
}

func (s *SQLCnter) FindDeviceByAllParticipant(ctx context.Context, eid string) (devices []string, err error) {
	result := s.gdb.Table("user_device").Select("device_id").Joins("left join event_participant on user_device.user_id = event_participant.user_id wherer event_participant.event_id = ?", eid).Scan(&devices)
	return devices, result.Error
}
func (s *SQLCnter) FindPetProfileByPetID(ctx context.Context, pid string) (pet *Pet, err error) {
	result := s.gdb.Table("pet").Where("id = ?", pid).First(pet)
	return pet, result.Error
}

func (s *SQLCnter) CreateUserDeviceID(ctx context.Context, uid string, device_id string) (err error) {
	value := &User_device{
		User_id:   uid,
		Device_id: device_id,
	}
	result := s.gdb.Table("user_device").Create(&value)
	return result.Error
}

func (s *SQLCnter) FindEventParticipantById(ctx context.Context, id string) (pets []string, participants []string) {
	var eventParticipant []Event_participant
	s.gdb.Table("event_participant").Where("Event_id = ?", id).Find(&eventParticipant)
	for i := 0; i < len(eventParticipant); i++ {
		cur := eventParticipant[i]
		pets = append(pets, cur.Pet_id)
		participants = append(participants, cur.Participant_id)
	}
	return pets, participants
}

func (s *SQLCnter) CreateNotification(ctx context.Context, n *Notification) error {
	result := s.gdb.Table("user_notification").Create(n)
	return result.Error
}

func (s *SQLCnter) FindUserIdListByEventId(ctx context.Context, eventId string) (res []string, err error) {
	result := s.gdb.Table("event_id").Select("participant_id").Where("event_id = ?", eventId).Find(res)
	return res, result.Error
}

func (s *SQLCnter) GetUserIdsbyPetIds(ctx context.Context, pid []string) (uid []string, err error) {
	result := s.gdb.Table("petowner").Distinct("user_id").Select("user_Id").Where("Pet_id IN ?", pid).Find(&uid)
	return uid, result.Error
}

func (s *SQLCnter) GetFriendsPetIdByPetId(ctx context.Context, pid string) (freindPids []string, err error) {
	var friendPids1 []string
	var friendPids2 []string
	result := s.gdb.Table("pet_connection").Where("Id1 = ?", pid).Select("Id2").Find(&friendPids1)
	if result.Error != nil {
		return nil, result.Error
	}
	result = s.gdb.Table("pet_connection").Where("Id2 = ?", pid).Select("Id1").Find(&friendPids2)
	freindPids = append(friendPids1, friendPids2...)
	return freindPids, result.Error
}

func (s *SQLCnter) UpdatesEvents(ctx context.Context, e Event) error {
	result := s.gdb.Table("event").Updates(&e)
	return result.Error
}

func (s *SQLCnter) GetEventByEventId(ctx context.Context, eid string) (event Event, err error) {

	result := s.gdb.Table("event").First(&event)
	return event, result.Error
}

func (s *SQLCnter) UpdateEventLocations(ctx context.Context, e EventLocation) (err error) {
	result := s.gdb.Table("event_location").Updates(&e)
	return result.Error
}

func (s *SQLCnter) GetNotificationByUserId(ctx context.Context, uid string) (n []Notification, err error) {
	result := s.gdb.Table("user_notification").Order("created_at DESC").
		Where("user_id = ? AND has_read = 0", uid).Find(&n)
	return n, result.Error
}

func (s *SQLCnter) UpdateNotificationHasReadStatus(ctx context.Context, n string) (err error) {
	result := s.gdb.Table("user_notification").Where("notification_id = ?", n).Update("has_read", true)
	return result.Error
}

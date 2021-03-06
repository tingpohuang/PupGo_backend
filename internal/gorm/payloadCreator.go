package gorm

import (
	"context"
	"fmt"
	"strconv"

	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
)

const (
	PET_RECOMMEND_LIST_LENGTH = 5
)

type PayloadCreator struct {
	sql *SQLCnter
}

func NewPayloadCreator(sqlCnter *SQLCnter) *PayloadCreator {
	return &PayloadCreator{
		sql: sqlCnter,
	}
}

func (p *PayloadCreator) GetUserProfileById(ctx context.Context, uid []string) (userProfiles []*model1.UserProfile) {
	users := p.sql.findUserByIdList(ctx, uid)
	usersLocation := p.createUserLocationById(ctx, uid)
	usersProfile := make([]*model1.UserProfile, len(users))
	for i := 0; i < len(users); i++ {
		user := users[i]
		userGender := model1.UserGender(strconv.Itoa(user.Gender))
		birthday := user.Birthday.String()
		usersProfile[i] = &model1.UserProfile{
			ID:       &user.Id,
			Name:     &user.Name,
			Gender:   &userGender,
			Birthday: &birthday,
			Email:    nil,
			Location: &usersLocation[i],
		}
	}
	return usersProfile
}

func (p *PayloadCreator) GetPetProfileById(ctx context.Context, pid []string) (petProfiles []*model1.PetProfile) {
	pets, petHobbies := p.sql.findPetByIdList(ctx, pid)
	petsProfile := make([]*model1.PetProfile, len(pets))
	usersLocations := p.sql.findUserLocationByPetsIdList(ctx, pid)
	for i := 0; i < len(pets); i++ {
		pet := pets[i]
		petGender := model1.PetGender(strconv.Itoa(pet.Gender))
		petLocation := p.createPetLocation(ctx, usersLocations[i])
		birthday := pet.Birthday.String()
		hobbies := petHobbies[i].hobbies
		petsProfile[i] = &model1.PetProfile{
			ID:           &pet.Id,
			Name:         &pet.Name,
			Image:        &pet.Image,
			Gender:       &petGender,
			Breed:        &pet.Breed,
			IsCastration: pet.IsCastration,
			Birthday:     &birthday,
			Location:     &petLocation,
			Description:  &pet.Description,
			Hobby:        hobbies,
		}
	}
	return petsProfile
}

// 0
// 1 first
// 2 second
// 3
// -1

func (p *PayloadCreator) GetPetRecommendationById(ctx context.Context, pid string) (petRecommendations []*model1.Recommendation) {
	petConnections := p.sql.findPetRecommend(ctx, pid)

	for i := 0; i < len(petConnections) && len(petRecommendations) < PET_RECOMMEND_LIST_LENGTH; i++ {
		petConnection := petConnections[i]
		var friendId = petConnection.Id1

		if pid == petConnection.Id1 {
			friendId = petConnection.Id2
			if petConnection.Status == 1 || petConnection.Status == 3 || petConnection.Status == -1 {
				continue
			}
		} else {
			if petConnection.Status == 2 || petConnection.Status == 3 || petConnection.Status == -1 {
				continue
			}
		}

		status := model1.RecommendationStatus(strconv.Itoa(petConnection.Status))
		petProfiles := p.GetPetProfileById(ctx, []string{friendId})

		petRecommendations = append(petRecommendations, &model1.Recommendation{
			ID:     friendId,
			Pet:    petProfiles[0],
			Status: &status,
		})
	}

	return petRecommendations

}

func (p *PayloadCreator) GetEventsById(ctx context.Context, id []string) (events []*model1.Event) {
	eventsRaw := p.sql.findEventByIdList(ctx, id)
	eventLocations := p.createEventLocationById(ctx, id)
	events = make([]*model1.Event, len(eventsRaw))
	for i := 0; i < len(eventsRaw); i++ {
		event := eventsRaw[i]
		eventLocation := eventLocations[i]
		eventLimit := model1.EventsLimits{
			LimitOfPet:  &event.Limit_pet_num,
			LimitOfUser: &event.Limit_user_num,
		}
		holderProfile := p.GetPetProfileById(ctx, []string{event.Holder_Id})
		startTime := event.Start_date.String()
		endTime := event.End_date.String()
		timeRange := model1.TimeRange{
			StartTime: &startTime,
			EndTime:   &endTime,
		}
		events[i] = &model1.Event{
			ID:          event.Id,
			Location:    &eventLocation,
			TimeRange:   &timeRange,
			Limit:       &eventLimit,
			Image:       &event.Image,
			Description: &event.Description,
			Holder:      holderProfile[0],
			Type:        &event.Type,
		}

	}
	return events
}

func (p *PayloadCreator) GetEventsByUId(ctx context.Context, uid string) (events []*model1.Event) {
	eventId := p.sql.findEventByUId(ctx, uid)
	eventsRaw := p.sql.findEventByIdList(ctx, eventId)
	eventLocations := p.createEventLocationById(ctx, eventId)
	events = make([]*model1.Event, len(eventsRaw))
	for i := 0; i < len(eventsRaw); i++ {
		event := eventsRaw[i]
		eventLocation := eventLocations[i]
		eventLimit := model1.EventsLimits{
			LimitOfPet:  &event.Limit_pet_num,
			LimitOfUser: &event.Limit_user_num,
		}
		holderProfile := p.GetPetProfileById(ctx, []string{event.Holder_Id})
		startTime := event.Start_date.String()
		endTime := event.End_date.String()
		timeRange := model1.TimeRange{
			StartTime: &startTime,
			EndTime:   &endTime,
		}
		events[i] = &model1.Event{
			ID:          event.Id,
			Location:    &eventLocation,
			TimeRange:   &timeRange,
			Limit:       &eventLimit,
			Image:       &event.Image,
			Description: &event.Description,
			Holder:      holderProfile[0],
			Type:        &event.Type,
		}

	}
	return events
}

func (p *PayloadCreator) GetRecommendEventsByUId(ctx context.Context, uid string) (events []*model1.Event) {
	eventId := p.sql.FindRecommendEventByUId(ctx, uid)
	eventsRaw := p.sql.findEventByIdList(ctx, eventId)
	eventLocations := p.createEventLocationById(ctx, eventId)
	events = make([]*model1.Event, len(eventsRaw))
	for i := 0; i < len(eventsRaw); i++ {
		event := eventsRaw[i]
		eventLocation := eventLocations[i]
		eventLimit := model1.EventsLimits{
			LimitOfPet:  &event.Limit_pet_num,
			LimitOfUser: &event.Limit_user_num,
		}
		holderProfile := p.GetPetProfileById(ctx, []string{event.Holder_Id})
		startTime := event.Start_date.String()
		endTime := event.End_date.String()
		timeRange := model1.TimeRange{
			StartTime: &startTime,
			EndTime:   &endTime,
		}
		events[i] = &model1.Event{
			ID:          event.Id,
			Location:    &eventLocation,
			TimeRange:   &timeRange,
			Limit:       &eventLimit,
			Image:       &event.Image,
			Description: &event.Description,
			Holder:      holderProfile[0],
			Type:        &event.Type,
		}

	}
	return events
}

func (p *PayloadCreator) createUserLocationById(ctx context.Context, uid []string) (userLocations []model1.Location) {
	locations, _ := p.sql.findUserLocationByIdList(ctx, uid)
	userLocations = make([]model1.Location, len(locations))
	for i := 0; i < len(locations); i++ {
		userLocation := locations[i]
		lat := fmt.Sprintf("%f", userLocation.Latitude)
		long := fmt.Sprintf("%f", userLocation.Longitude)
		userLocations[i] = model1.Location{
			Country:   &userLocation.Country,
			State:     &userLocation.State,
			City:      &userLocation.City,
			Address:   &userLocation.Address,
			Latitude:  &lat,
			Longitude: &long,
		}
	}

	return userLocations
}

func (p *PayloadCreator) createPetLocation(ctx context.Context, userLocation UserLocation) (petLocation model1.Location) {

	lat := fmt.Sprintf("%f", userLocation.Latitude)
	long := fmt.Sprintf("%f", userLocation.Longitude)
	petLocation = model1.Location{
		Country:   &userLocation.Country,
		State:     &userLocation.State,
		City:      &userLocation.City,
		Address:   &userLocation.Address,
		Latitude:  &lat,
		Longitude: &long,
	}

	return petLocation
}

func (p *PayloadCreator) createEventLocationById(ctx context.Context, id []string) (eventLocations []model1.Location) {
	locations := p.sql.findEventLocationByIdList(ctx, id)
	eventLocations = make([]model1.Location, len(locations))
	for i := 0; i < len(locations); i++ {
		eventLocation := locations[i]
		lat := fmt.Sprintf("%f", eventLocation.Latitude)
		long := fmt.Sprintf("%f", eventLocation.Longitude)
		eventLocations[i] = model1.Location{
			Country:   &eventLocation.Country,
			State:     &eventLocation.State,
			City:      &eventLocation.City,
			Address:   &eventLocation.Address,
			Latitude:  &lat,
			Longitude: &long,
		}
	}

	return eventLocations
}

func (p *PayloadCreator) GetPetListByUId(ctx context.Context, uid string) (petProfiles []*model1.PetProfile) {
	pets := p.sql.findPetsByUId(ctx, uid)
	petsProfile := p.GetPetProfileById(ctx, pets)
	return petsProfile
}

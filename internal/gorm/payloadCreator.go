package gorm

import (
	"context"
	"fmt"
	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
	"strconv"
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
	pets := p.sql.findPetByIdList(ctx, pid)
	petsProfile := make([]*model1.PetProfile, len(pets))
	for i := 0; i < len(pets); i++ {
		pet := pets[i]
		petGender := model1.PetGender(strconv.Itoa(pet.Gender))
		birthday := pet.Birthday.String()
		petsProfile[i] = &model1.PetProfile{
			ID:           &pet.Id,
			Name:         &pet.Name,
			Image:        &pet.Image,
			Gender:       &petGender,
			Breed:        &pet.Breed,
			IsCastration: pet.IsCastration,
			Birthday:     &birthday,
			Location:     nil,
		}
	}
	return petsProfile
}

func (p *PayloadCreator) GetPetRecommendationById(ctx context.Context, pid string) (petRecommendations []*model1.Recommendation) {
	petConnections := p.sql.findPetRecommend(ctx, pid)
	recommendations := make([]*model1.Recommendation, len(petConnections))
	for i := 0; i < len(petConnections); i++ {
		petConnection := petConnections[i]
		var friendId = petConnection.Id1
		if pid == petConnection.Id1 {
			friendId = petConnection.Id2
		}

		status := model1.RecommendationStatus(strconv.Itoa(petConnection.Status))
		petProfiles := p.GetPetProfileById(ctx, []string{friendId})

		recommendations[i] = &model1.Recommendation{
			ID:     friendId,
			Pet:    petProfiles[0],
			Status: &status,
		}
	}

	return recommendations

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
		pets, participants := p.sql.findEventParticipantById(ctx, event.Id)
		petsProfile := p.GetPetProfileById(ctx, pets)
		participantsProfile := p.GetUserProfileById(ctx, participants)
		startTime := event.Start_date.String()
		endTime := event.End_date.String()
		timeRange := model1.TimeRange{
			StartTime: &startTime,
			EndTime:   &endTime,
		}
		events[i] = &model1.Event{
			ID:           event.Id,
			Location:     &eventLocation,
			TimeRange:    &timeRange,
			Limit:        &eventLimit,
			Image:        &event.Image,
			Description:  []string{event.Description},
			Holder:       holderProfile[0],
			Pets:         petsProfile,
			Participants: participantsProfile,
		}

	}
	return events
}

func (p *PayloadCreator) createUserLocationById(ctx context.Context, uid []string) (userLocations []model1.Location) {
	locations := p.sql.findUserLocationByIdList(ctx, uid)
	userLocations = make([]model1.Location, len(locations))
	for i := 0; i < len(locations); i++ {
		userLocation := locations[i]
		lat := fmt.Sprintf("%f", userLocation.Position.Lat)
		long := fmt.Sprintf("%f", userLocation.Position.Long)
		userLocations[i] = model1.Location{
			Country: &userLocation.Country,
			State:   &userLocation.State,
			City:    &userLocation.City,
			Address: &userLocation.Address,
			Coor: &model1.Coordinate{
				IsBlur:    false,
				Latitude:  &lat,
				Longitude: &long,
			},
		}
	}

	return userLocations
}

func (p *PayloadCreator) createEventLocationById(ctx context.Context, id []string) (eventLocations []model1.Location) {
	locations := p.sql.findEventLocationByIdList(ctx, id)
	eventLocations = make([]model1.Location, len(locations))
	for i := 0; i < len(locations); i++ {
		eventLocation := locations[i]
		lat := fmt.Sprintf("%f", eventLocation.Position.Lat)
		long := fmt.Sprintf("%f", eventLocation.Position.Long)
		eventLocations[i] = model1.Location{
			Country: &eventLocation.Country,
			State:   &eventLocation.State,
			City:    &eventLocation.City,
			Address: &eventLocation.Address,
			Coor: &model1.Coordinate{
				IsBlur:    false,
				Latitude:  &lat,
				Longitude: &long,
			},
		}
	}

	return eventLocations
}

func (p *PayloadCreator) GetPetListByUId(ctx context.Context, uid string) (petProfiles []*model1.PetProfile) {
	pets := p.sql.findPetsByUId(ctx, uid)
	petsProfile := p.GetPetProfileById(ctx, pets)
	return petsProfile
}

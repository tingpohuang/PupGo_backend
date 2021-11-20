package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"
	generated1 "github.com/tingpo/pupgobackend/internal/graph/generated"
	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
)

func (r *queryResolver) EventsListGet(ctx context.Context, eventsListGetInput model1.EventsListGetInput) (*model1.EventsListGetPayload, error) {
	timestamp := time.Now().String()
	newPayload := &model1.EventsListGetPayload{
		Error:     nil,
		Result:    []*model1.Event{},
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) NotifiactionsGet(ctx context.Context, notifiactionsGetInput model1.NotifiactionsGetInput) (*model1.NotifiactionsGetPayload, error) {
	timestamp := time.Now().String()
	newPayload := &model1.NotifiactionsGetPayload{
		Error:     nil,
		Result:    []*model1.Notification{},
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) RecommendationGet(ctx context.Context, recommendationGetInput model1.RecommendationGetInput) (*model1.RecommendationGetPayload, error) {
	timestamp := time.Now().String()
	newPayload := &model1.RecommendationGetPayload{
		Error:     nil,
		Result:    []*model1.Recommendation{},
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) FriendsListGet(ctx context.Context, friendsListGetInput model1.FriendsListGetInput) (*model1.FriendsListGetPayload, error) {
	// var Errors []Error
	newPet := &model1.Pet{
		ID:              uuid.NewString(),
		Owner:           nil,
		PetProfile:      nil,
		PetRelationShip: nil,
	}
	timestamp := time.Now().String()
	newPayload := &model1.FriendsListGetPayload{
		Error:     nil,
		Result:    []*model1.Pet{newPet},
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) PetProfileListGet(ctx context.Context, petProfileListGetInput model1.PetProfileListGetInput) (*model1.PetProfileListGetPayload, error) {
	timestamp := time.Now().String()
	pets := sqlCnter.FindPetByIdList(ctx, petProfileListGetInput.Pid)
	petsProfile := make([]*model1.PetProfile, len(pets))
	for i := 0; i < len(pets); i++ {
		pet := pets[i]
		petGender := model1.PetGender(strconv.Itoa(pet.Gender))
		petsProfile[i] = &model1.PetProfile{
			ID:           &pet.Id,
			Name:         &pet.Name,
			Image:        &pet.Image,
			Gender:       &petGender,
			Breed:        &pet.Breed,
			IsCastration: pet.IsCastration,
			Birthday:     nil,
			Location:     nil,
		}
	}
	newPayload := &model1.PetProfileListGetPayload{
		Error:     nil,
		Result:    petsProfile,
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) UserProfileListGet(ctx context.Context, userProfileListGetInput model1.UserProfileListGetInput) (*model1.UserProfileListGetPayload, error) {
	timestamp := time.Now().String()
	users := sqlCnter.FindUserByIdList(ctx, userProfileListGetInput.UID)
	usersProfile := make([]*model1.UserProfile, len(users))
	for i := 0; i < len(users); i++ {
		user := users[i]
		userGender := model1.UserGender(strconv.Itoa(user.Gender))
		usersProfile[i] = &model1.UserProfile{
			ID:       &user.Id,
			Name:     &user.Name,
			Gender:   &userGender,
			Age:      nil,
			Email:    nil,
			Location: nil,
		}
	}

	newPayload := &model1.UserProfileListGetPayload{
		Error:     nil,
		Result:    usersProfile,
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) ProfileListGet(ctx context.Context, profileListGetInput model1.ProfileListGetInput) (*model1.ProfileListGetPayload, error) {
	timestamp := time.Now().String()
	newPayload := &model1.ProfileListGetPayload{
		Error: nil,
		//Result:    []*model1.UserProfile{},
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	TIMEOUT int = 30
)

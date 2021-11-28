package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	generated1 "github.com/tingpo/pupgobackend/internal/graph/generated"
	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
)

func (r *queryResolver) EventsListGet(ctx context.Context, eventsListGetInput model1.EventsListGetInput) (*model1.EventsListGetPayload, error) {
	timestamp := time.Now().String()
	events := payloadCreator.GetEventsByUId(ctx, eventsListGetInput.UID)
	newPayload := &model1.EventsListGetPayload{
		Error:     nil,
		Result:    events,
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) NotificationsGet(ctx context.Context, notificationsGetInput model1.NotificationsGetInput) (*model1.NotificationsGetPayload, error) {
	timestamp := time.Now().String()
	res, err := sqlCnter.GetNotificationByUserId(ctx, notificationsGetInput.UID)
	if err != nil {
		return nil, err
	}
	// fmt.Print(res)
	m := make([]*model1.Notification, len(res))
	for i := 0; i < len(res); i++ {
		m[i] = &model1.Notification{
			NotificationID:   res[i].Notification_id,
			NotificationType: &res[i].Notification_type,
			UserID:           &res[i].User_id,
			EventID:          &res[i].Event_id,
			PetID:            &res[i].Pet_id,
			HasRead:          &res[i].Has_read,
		}
	}

	payload := &model1.NotificationsGetPayload{
		Error:     nil,
		Result:    m,
		Timestamp: &timestamp,
	}
	return payload, nil
}

func (r *queryResolver) RecommendationGet(ctx context.Context, recommendationGetInput model1.RecommendationGetInput) (*model1.RecommendationGetPayload, error) {
	timestamp := time.Now().String()
	petConnections := payloadCreator.GetPetRecommendationById(ctx, recommendationGetInput.Pid)
	newPayload := &model1.RecommendationGetPayload{
		Error:     nil,
		Result:    petConnections,
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
	petsProfile := payloadCreator.GetPetProfileById(ctx, petProfileListGetInput.Pid)
	newPayload := &model1.PetProfileListGetPayload{
		Error:     nil,
		Result:    petsProfile,
		Timestamp: &timestamp,
	}
	return newPayload, nil
}

func (r *queryResolver) UserProfileListGet(ctx context.Context, userProfileListGetInput model1.UserProfileListGetInput) (*model1.UserProfileListGetPayload, error) {
	timestamp := time.Now().String()
	usersProfile := payloadCreator.GetUserProfileById(ctx, userProfileListGetInput.UID)
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

func (r *queryResolver) PetsListGet(ctx context.Context, petsListGetInput model1.PetsListGetInput) (*model1.PetsListGetPayload, error) {
	timestamp := time.Now().String()
	petsList := payloadCreator.GetPetListByUId(ctx, petsListGetInput.UID)
	newPayload := &model1.PetsListGetPayload{
		Error:     nil,
		Timestamp: &timestamp,
		Result:    petsList,
	}
	return newPayload, nil
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

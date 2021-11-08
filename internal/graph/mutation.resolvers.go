package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/tingpo/pupgobackend/internal/graph/generated"
	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
)

func (r *mutationResolver) UserCreateByID(ctx context.Context, userCreateByIDInput model1.UserCreateByIDInput) (*model1.UserCreateByIDPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EventsCreate(ctx context.Context, eventsCreateInput model1.EventsCreateInput) (*model1.EventsCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EventsJoin(ctx context.Context, eventsJoinInput model1.EventsJoinInput) (*model1.EventsJoinPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) NotificationRemove(ctx context.Context, notificationRemoveInput model1.NotificationRemoveInput) (*model1.NotificationRemovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RecommendationResponse(ctx context.Context, recommendationResponseInput model1.RecommendationResponseInput) (*model1.RecommendationResponsePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FriendRemove(ctx context.Context, friendRemoveInput model1.FriendRemoveInput) (*model1.FriendRemovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PetProfileUpdates(ctx context.Context, petProfileUpdatesInput model1.PetProfileUpdatesInput) (*model1.PetProfileUpdatesPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PetCreate(ctx context.Context, petCreateInput model1.PetCreateInput) (*model1.PetCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PetDelete(ctx context.Context, petDeleteInput model1.PetDeleteInput) (*model1.PetDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatesNotificationSettings(ctx context.Context, updatesNotificationSettingsInput model1.UpdatesNotificationSettingsInput) (*model1.UpdatesNotificationSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

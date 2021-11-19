package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/tingpo/pupgobackend/internal/graph/generated"
	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
)

func (r *queryResolver) EventsListGet(ctx context.Context, eventsListGetInput model1.EventsListGetInput) (*model1.EventsListGetPayload, error) {

	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NotifiactionsGet(ctx context.Context, notifiactionsGetInput model1.NotifiactionsGetInput) (*model1.NotifiactionsGetPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RecommendationGet(ctx context.Context, recommendationGetInput model1.RecommendationGetInput) (*model1.RecommendationGetPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FriendsListGet(ctx context.Context, friendsListGetInput model1.FriendsListGetInput) (*model1.FriendsListGetPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PetProfileListGet(ctx context.Context, petProfileListGetInput model1.PetProfileListGetInput) (*model1.PetProfileListGetPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserProfileListGet(ctx context.Context, userProfileListGetInput model1.UserProfileListGetInput) (*model1.UserProfileListGetPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProfileListGet(ctx context.Context, profileListGetInput model1.ProfileListGetInput) (*model1.ProfileListGetPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

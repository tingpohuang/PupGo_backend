package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	generated1 "github.com/tingpo/pupgobackend/internal/graph/generated"
	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
)

func (r *mutationResolver) UserCreateByID(ctx context.Context, userCreateByIDInput model1.UserCreateByIDInput) (*model1.UserCreateByIDPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EventsCreate(ctx context.Context, eventsCreateInput model1.EventsCreateInput) (*model1.EventsCreatePayload, error) {
	// if eci.ID != "" {
	// 	p := new(EmptyValueError)
	// }
	loc := eventsCreateInput.Location
	tr := eventsCreateInput.TimeRange
	lmt := eventsCreateInput.Limit
	img := eventsCreateInput.Image
	desc := eventsCreateInput.Description
	errmsg := ""
	if loc == nil {
		errmsg += "location cannot be empty value."
	}
	if tr == nil {
		errmsg += "time range cannot be empty value"
	}
	if lmt == nil {
		// dftv := 5
		lmt = &model1.EventsLimitsInput{
			LimitOfDog:   &defaultEventLimitPet,
			LimitOfHuman: &defaultEventLimitHuman,
		}
	}
	if img == nil {
		dftvImg := "test.img"
		img = &dftvImg
	}
	if desc == nil {
		desc = new(String)
	}
	if errmsg != "" {
		return nil, errors.New(errmsg)
	}
	ret := new(model1.EventsCreatePayload)
	tstmp := time.Now().String()
	ret.Timestamp = &tstmp

	return ret, nil
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
	name := petCreateInput.Name
	img := petCreateInput.Image
	gender := petCreateInput.Gender
	breed := petCreateInput.Breed
	isCastration := petCreateInput.IsCastration
	birthday := petCreateInput.Birthday
	uid := petCreateInput.UID
	errmsg := ""

	if name == nil || *name == "" {
		errmsg += "name cannot be empty string."
	}
	if img == nil || *img == "" {
		img = &defaultPetImageUrl
	}
	if gender == nil || *gender == "" {
		*gender = "unknown"
	}
	if breed == nil || *breed == "" {
		*breed = "unknown"
	}
	if birthday == nil || *birthday == "" {
		errmsg += "birthday cannot be empty string."
	}
	if uid == "" {
		errmsg += "user should not be empty."
	}

	if errmsg != "" {
		return nil, errors.New(errmsg)
	}
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	defaultPetImageUrl     = "hahaha.jpg"
	defaultEventLimitPet   = 5
	defaultEventLimitHuman = 5
)

type EmptyValueError struct {
}

func (EmptyValueError) IsError() {}

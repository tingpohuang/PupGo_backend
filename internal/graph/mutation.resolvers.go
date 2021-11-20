package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	gorm "github.com/tingpo/pupgobackend/internal/gorm"
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
	// desc := eventsCreateInput.Description
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
	// if desc == nil {
	// 	desc = new(String)
	// }
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
	pid := friendRemoveInput.PetID
	friendId := friendRemoveInput.FriendID
	err := removeFriend(ctx, pid, friendId)
	if err != nil {
		return nil, err
	} else {
		return &model1.FriendRemovePayload{
			Timestamp: getNowTimestamp(),
		}, nil
	}
	// panic(fmt.Errorf("not implemented"))
}
func getNowTimestamp() *string {
	tstmp := time.Now().String()
	return &tstmp
}
func removeFriend(ctx context.Context, id1 string, id2 string) error {
	if id1 > id2 {
		id1, id2 = id2, id1
	}
	err := sqlCnter.DeleteFriend(ctx, id1, id2)
	return err
}

func (r *mutationResolver) PetProfileUpdates(ctx context.Context, petProfileUpdatesInput model1.PetProfileUpdatesInput) (*model1.PetProfileUpdatesPayload, error) {
	id := petProfileUpdatesInput.ID
	if id == "" {
		return nil, errors.New("cannot updates null id")
	}
	petProfileList := sqlCnter.FindPetById(ctx, id)
	if len(petProfileList) < 1 {
		return nil, errors.New("pet id doesn't exist")
	}
	petProfile := petProfileList[0]
	if petProfileUpdatesInput.Name != nil {
		petProfile.Name = *petProfileUpdatesInput.Name
	}
	if petProfileUpdatesInput.Image != nil {
		petProfile.Image = *petProfileUpdatesInput.Image
	}
	if petProfileUpdatesInput.Gender != nil {
		if *petProfileUpdatesInput.Gender == "Male" {
			petProfile.Gender = 1
		}
	}
	if petProfileUpdatesInput.Image != nil {
		petProfile.Breed = *petProfileUpdatesInput.Breed
	}
	if petProfileUpdatesInput.IsCastration {
		petProfile.IsCastration = petProfileUpdatesInput.IsCastration
	}
	// TODO: location, birthday
	err := sqlCnter.UpdatePet(ctx, petProfile)
	if err != nil {
		return nil, err
	}
	gender := petGenderIntToString(petProfile.Gender)
	pgender := model1.PetGender(*gender)
	tstmp := time.Now().String()
	m := &model1.PetProfileUpdatesPayload{
		Error:     nil,
		Timestamp: &tstmp,
		Result: &model1.PetProfile{
			ID:           &petProfile.Id,
			Name:         &petProfile.Name,
			Image:        &petProfile.Image,
			Gender:       &pgender,
			Breed:        &petProfile.Breed,
			IsCastration: petProfile.IsCastration,
		},
	}
	return m, nil
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
	gender_num := 0
	if name == nil || *name == "" {
		errmsg += "name cannot be empty string."
	}
	if img == nil || *img == "" {
		img = &defaultPetImageUrl
	}
	if *gender == "Male" {
		gender_num = 1
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
	pid := uuid.NewString()
	res := sqlCnter.CreatePets(ctx, gorm.Pet{
		Id:           pid,
		Name:         *name,
		Image:        *img,
		Gender:       gender_num,
		Breed:        *breed,
		IsCastration: isCastration,
	})
	if res != nil {
		return nil, errors.New("internal error in SQL")
	}
	res = sqlCnter.CreateUserPetRelation(ctx, uid, pid)
	if res != nil {
		return nil, errors.New("internal error in SQL")
	}
	return &model1.PetCreatePayload{
		Error: nil,
		Result: &model1.PetProfile{
			ID:           &pid,
			Name:         name,
			Image:        img,
			Gender:       petCreateInput.Gender,
			Breed:        petCreateInput.Breed,
			IsCastration: petCreateInput.IsCastration,
		}}, nil
}

func (r *mutationResolver) PetDelete(ctx context.Context, petDeleteInput model1.PetDeleteInput) (*model1.PetDeletePayload, error) {
	if petDeleteInput.Pid == "" {
		return nil, errors.New("pid cannot be null")
	}
	sqlCnter.DeletePet(ctx, petDeleteInput.Pid)
	return &model1.PetDeletePayload{
		Error:  nil,
		Result: true}, nil
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

func petGenderIntToString(i int) *string {
	var s string
	if i == 1 {
		s = "Male"
	} else if i == 2 {
		s = "Female"
	} else {
		s = "Unknown"
	}
	return &s
}
func petGenderStringToInt(s string) int {
	if s == "Male" {
		return 1
	} else if s == "Female" {
		return 2
	}
	return 0
}

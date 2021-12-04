package graph

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tingpo/pupgobackend/internal/gorm"
	"github.com/tingpo/pupgobackend/internal/graph/model"
)

const ()

var (
	s      = "0"
	trange = model.TimeRangeInput{
		StartTime: &s,
		EndTime:   &s,
	}
	c = "US"
	l = model.LocationInput{
		Country: &c,
	}
)

func TestEventsCreate(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()

	_, err := r.EventsCreate(ctx, model.EventsCreateInput{
		Pid:       gorm.Pet_1_id,
		TimeRange: &trange,
		Location:  &l,
	})
	assert.Nil(err)
}

func TestEventsUpdate(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	image := "url"
	place := "0.0001"
	country := "US"
	state := "CA"
	city := "LA"
	address := "1111"
	_, err := r.EventsUpdate(ctx, model.EventsUpdateInput{
		Eid:   gorm.Event_1_id,
		Pid:   gorm.Pet_1_id,
		Image: &image,
		Location: &model.LocationInput{
			Country: &country,
			City:    &city,
			Address: &address,
			State:   &state,
			Coordinate: &model.CoordinateInput{
				Latitude:  &place,
				Longitude: &place,
			},
		},
	})
	assert.Nil(err)
}

func TestEventsJoin(t *testing.T) {
	ctx := context.Background()
	// assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	d := "123321"
	_, _ = r.EventsJoin(ctx, model.EventsJoinInput{
		Pid:         gorm.Pet_1_id,
		EventID:     gorm.Event_1_id,
		Description: &d,
	})
	// assert.Nil(err)
}

func TestEventsJoin2(t *testing.T) {
	ctx := context.Background()
	// assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	d := "123321"
	_, _ = r.EventsJoin(ctx, model.EventsJoinInput{
		Pid:         gorm.Pet_2_id,
		EventID:     gorm.Event_1_id,
		Description: &d,
	})
	// assert.Nil(err)
}

func TestNotificationRead(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	_, err := r.NotificationRead(ctx, model.NotificationReadInput{
		Nid: "0",
	})
	assert.Nil(err)
}

func TestRecommendationResponse(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	_, err := r.RecommendationResponse(ctx, model.RecommendationResponseInput{
		Pid:         gorm.Pet_2_id,
		RecommendID: gorm.Pet_3_id,
		Result:      true,
	})
	assert.Nil(err)
	_, _ = r.RecommendationResponse(ctx, model.RecommendationResponseInput{
		Pid:         gorm.Pet_3_id,
		RecommendID: gorm.Pet_2_id,
		Result:      true,
	})
	// assert.Nil(err)
	// _, _ = r.RecommendationResponse(ctx, model.RecommendationResponseInput{
	// 	Pid:         gorm.Pet_4_id,
	// 	RecommendID: gorm.Pet_2_id,
	// 	Result:      true,
	// })
	// _, _ = r.RecommendationResponse(ctx, model.RecommendationResponseInput{
	// 	Pid:         gorm.Pet_2_id,
	// 	RecommendID: gorm.Pet_4_id,
	// 	Result:      true,
	// })
	// _, _ = r.RecommendationResponse(ctx, model.RecommendationResponseInput{
	// 	Pid:         gorm.Pet_2_id,
	// 	RecommendID: gorm.Pet_4_id,
	// 	Result:      false,
	// })
	// _, err = r.RecommendationResponse(ctx, model.RecommendationResponseInput{
	// 	Pid:         gorm.Pet_4_id,
	// 	RecommendID: gorm.Pet_2_id,
	// 	Result:      true,
	// })
	// assert.Nil(err)
}

func TestFriendsListGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.FriendsListGet(ctx, model.FriendsListGetInput{
		Pid: gorm.Pet_1_id,
	})
	assert.Nil(err)
}

func TestPetCreateDelete(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	name := "name"
	image := "a.jpg"
	pg := model.PetGender("Male")
	breed := "husky"
	res, err := r.PetCreate(ctx, model.PetCreateInput{
		Name:     &name,
		Image:    &image,
		Gender:   &pg,
		Breed:    &breed,
		Birthday: &s,
		Location: &l,
		UID:      gorm.User_1_id,
	})
	assert.Nil(err)
	_, err = r.PetDelete(ctx, model.PetDeleteInput{
		Pid: *res.Result.ID,
	})
	assert.Nil(err)
}

func TestPetProfileUpdates(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	name := "gu"
	image := "iiig"
	gender := "Female"
	breed := "haahahusky"
	_, err := r.PetProfileUpdates(ctx, model.PetProfileUpdatesInput{
		ID:           gorm.Pet_2_id,
		Name:         &name,
		Image:        &image,
		Gender:       (*model.PetGender)(&gender),
		Breed:        &breed,
		IsCastration: false,
	})
	assert.Nil(err)
}

func TestFriendRemove(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	_, err := r.FriendRemove(ctx, model.FriendRemoveInput{
		PetID:    gorm.Pet_1_id,
		FriendID: gorm.Pet_3_id,
	})
	assert.Nil(err)
}

func TestEventsAccept(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Mutation()
	_, err := r.EventsAccept(ctx, model.EventsAcceptInput{
		Pid:     gorm.Pet_1_id,
		EventID: gorm.Event_1_id,
		Accept:  true,
	})
	assert.Nil(err)
}

func TestPetsListGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.PetsListGet(ctx, model.PetsListGetInput{
		UID: gorm.User_1_id,
	})
	assert.Nil(err)
}

func TestProfileListGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.ProfileListGet(ctx, model.ProfileListGetInput{
		ID: []string{gorm.User_1_id, gorm.User_2_id},
	})
	assert.Nil(err)
}
func TestUserProfileListGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.UserProfileListGet(ctx, model.UserProfileListGetInput{
		UID: []string{gorm.User_1_id, gorm.User_2_id},
	})
	assert.Nil(err)
}

func TestPetProfileListGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.PetProfileListGet(ctx, model.PetProfileListGetInput{
		Pid: []string{gorm.Pet_1_id, gorm.Pet_2_id},
	})
	assert.Nil(err)
}

func TestRecommendationGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.RecommendationGet(ctx, model.RecommendationGetInput{
		Pid: gorm.Pet_1_id,
	})
	assert.Nil(err)
}

func TestNotificationsGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.NotificationsGet(ctx, model.NotificationsGetInput{
		UID: gorm.User_1_id,
	})
	assert.Nil(err)
}

func TestRecommendEventsListGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.RecommendEventsListGet(ctx, model.EventsListGetInput{
		UID: gorm.User_1_id,
	})
	assert.Nil(err)
}

func TestEventsListGet(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	a := &Resolver{}
	r := a.Query()
	_, err := r.EventsListGet(ctx, model.EventsListGetInput{
		UID: gorm.User_1_id,
	})
	assert.Nil(err)
}

func TestPetGenderStringToInt(t *testing.T) {
	assert := assert.New(t)
	v := PetGenderStringToInt("Male")
	assert.Equal(1, v)
	v = PetGenderStringToInt("Female")
	assert.Equal(2, v)
	v = PetGenderStringToInt("F")
	assert.Equal(0, v)
}

func TestToString(t *testing.T) {
	assert := assert.New(t)
	v := PetGenderIntToString(1)
	assert.Equal("Male", *v)
	v = PetGenderIntToString(2)
	assert.Equal("Female", *v)
}

func TestGetUidbyPid(t *testing.T) {
	GetUidbyPid(context.Background(), "")
	GetUidbyPid(context.Background(), gorm.Pet_ids[0])
}

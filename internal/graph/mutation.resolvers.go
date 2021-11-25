package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tingpo/pupgobackend/internal/gorm"
	generated1 "github.com/tingpo/pupgobackend/internal/graph/generated"
	model1 "github.com/tingpo/pupgobackend/internal/graph/model"
	"github.com/tingpo/pupgobackend/internal/notification"
	gormio "gorm.io/gorm"
)

func (r *mutationResolver) UserCreateByID(ctx context.Context, userCreateByIDInput model1.UserCreateByIDInput) (*model1.UserCreateByIDPayload, error) {
	panic(fmt.Errorf("this function havs closed"))
}

func (r *mutationResolver) EventsCreate(ctx context.Context, eventsCreateInput model1.EventsCreateInput) (*model1.EventsCreatePayload, error) {
	pid := eventsCreateInput.Pid
	loc := eventsCreateInput.Location
	tr := eventsCreateInput.TimeRange
	lmt := eventsCreateInput.Limit
	img := eventsCreateInput.Image
	errmsg := ""
	if eventsCreateInput.Pid == "" {
		errmsg += "pid cannot be empty"
	}
	if loc == nil {
		errmsg += "location cannot be empty value."
	}
	if tr == nil {
		errmsg += "time range cannot be empty value"
	}
	if lmt == nil {
		lmt = &model1.EventsLimitsInput{
			LimitOfDog:   &defaultEventLimitPet,
			LimitOfHuman: &defaultEventLimitHuman,
		}
	}
	if img == nil {
		dftvImg := "test.img"
		img = &dftvImg
	}
	if errmsg != "" {
		return nil, errors.New(errmsg)
	}
	// trStart := tr.StartTime.String()
	// trEnd := tr.EndTime.String()
	data := gorm.Event{
		Id:        uuid.NewString(),
		Holder_Id: pid,
		// Start_date: &trStart,
		Start_date:     time.Now(),
		End_date:       time.Now(),
		Image:          *img,
		Limit_user_num: 5,
		Limit_pet_num:  5,
		Description:    "",
	}
	err := sqlCnter.CreateEvents(ctx, data)
	if err != nil {
		return nil, err
	}
	tstmp := time.Now().String()
	ret := &model1.EventsCreatePayload{
		Timestamp: &tstmp,
		Result: &model1.Event{
			ID:           data.Id,
			TimeRange:    &model1.TimeRange{},
			Holder:       nil,
			Participants: nil,
			Image:        &data.Image,
		},
	}
	return ret, nil
}

func (r *mutationResolver) EventsJoin(ctx context.Context, eventsJoinInput model1.EventsJoinInput) (*model1.EventsJoinPayload, error) {
	pid := eventsJoinInput.Pid
	eid := eventsJoinInput.EventID
	if pid == "" {
		return nil, errors.New("pid should not be empty")
	}
	if eid == "" {
		return nil, errors.New("event id should not be empty")
	}
	_, err := sqlCnter.FindEventsParticipantByPetID(ctx, pid, eid)
	if !errors.Is(err, gormio.ErrRecordNotFound) {
		return nil, errors.New("already exists participant log")
	} else {
		uid, err := sqlCnter.GetUserIdbyPetId(ctx, pid)
		if err != nil {
			return nil, err
		}
		err = sqlCnter.CreateParticipants(ctx, gorm.Event_participant{
			Event_id:       eid,
			Participant_id: *uid,
			Pet_id:         pid,
			Status:         0,
		})
		n := notification.Notification{}
		// bug over here
		go n.SendNewParticipantsMessage(context.Background(), pid, pid, sqlCnter)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (r *mutationResolver) EventsAccept(ctx context.Context, eventsAcceptInput model1.EventsAcceptInput) (*model1.EventsAcceptPayload, error) {
	res := &model1.EventsAcceptPayload{
		Error:     nil,
		Timestamp: GetNowTimestamp(),
		Result:    nil,
	}
	pid := eventsAcceptInput.Pid
	eid := eventsAcceptInput.EventID
	// if pid == "" {
	// 	return nil, errors.New("pid should not be empty")
	// }
	// if eid == "" {
	// 	return nil, errors.New("event id should not be empty")
	// }
	var status EventStatus = EventStatusNoAnswer
	if !eventsAcceptInput.Accept {
		status = EventStatusDecline
	} else {
		status = EventStatusAccept
	}
	participants, err := sqlCnter.FindEventsParticipantByPetID(ctx, pid, eid)
	if err != nil {
		return nil, err
	}
	participants.Status = int(status)
	err = sqlCnter.UpdateParticipants(ctx, *participants)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) NotificationRemove(ctx context.Context, notificationRemoveInput model1.NotificationRemoveInput) (*model1.NotificationRemovePayload, error) {
	panic(fmt.Errorf("this function no long support"))
}

func (r *mutationResolver) RecommendationResponse(ctx context.Context, recommendationResponseInput model1.RecommendationResponseInput) (*model1.RecommendationResponsePayload, error) {
	payload := &model1.RecommendationResponsePayload{
		Timestamp: GetNowTimestamp(),
	}
	petId := recommendationResponseInput.Pid
	recommendId := recommendationResponseInput.RecommendID
	result := recommendationResponseInput.Result
	res, err := sqlCnter.FindPetRecommendByID(ctx, petId, recommendId)
	if err != nil {
		return nil, errors.New("pet recommend not exist")
	}
	if !result {
		res.Status = int(RecommendationStatusDecline)
		err := sqlCnter.UpdatePetRecommendByID(ctx, res)
		return payload, err
	}
	if res.Status == int(RecommendationStatusLowAgree) && petId > recommendId { // first pet agree
		err := sqlCnter.CreatePetConnection(ctx, petId, recommendId)
		if err != nil {
			return nil, err
		}
		res.Status = int(RecommendationStatusBothAgree) // means all agree
		err = sqlCnter.UpdatePetRecommendByID(ctx, res)
		if err != nil {
			return nil, err
		}
		n := notification.Notification{}
		go n.SendNewFriendMessage(context.Background(), petId, recommendId, sqlCnter)
	} else if res.Status == int(RecommendationStatusHighAgree) && petId < recommendId { // second pet agree
		err := sqlCnter.CreatePetConnection(ctx, petId, recommendId)
		if err != nil {
			return nil, err
		}
		res.Status = int(RecommendationStatusBothAgree) // means all agree
		err = sqlCnter.UpdatePetRecommendByID(ctx, res)
		if err != nil {
			return nil, err
		}
		n := notification.Notification{}
		go n.SendNewFriendMessage(context.Background(), petId, recommendId, sqlCnter)
	} else if res.Status == int(RecommendationStatusNoAnswer) && petId < recommendId { // no pet agree
		res.Status = int(RecommendationStatusLowAgree)
		err := sqlCnter.UpdatePetRecommendByID(ctx, res)
		if err != nil {
			return nil, err
		}
		n := notification.Notification{}
		//check for go
		go n.SendFriendsInviteMessage(context.Background(), petId, recommendId, sqlCnter)
	} else if res.Status == int(RecommendationStatusNoAnswer) && petId > recommendId { // no pet agree
		res.Status = int(RecommendationStatusHighAgree)
		err := sqlCnter.UpdatePetRecommendByID(ctx, res)
		if err != nil {
			return nil, err
		}
		n := notification.Notification{}
		//check for go
		go n.SendFriendsInviteMessage(context.Background(), petId, recommendId, sqlCnter)
	}
	payload.Result = &model1.PetProfile{}
	return payload, nil
}

func (r *mutationResolver) FriendRemove(ctx context.Context, friendRemoveInput model1.FriendRemoveInput) (*model1.FriendRemovePayload, error) {
	pid := friendRemoveInput.PetID
	friendId := friendRemoveInput.FriendID
	err := RemoveFriend(ctx, pid, friendId)
	if err != nil {
		return nil, err
	} else {
		return &model1.FriendRemovePayload{
			Timestamp: GetNowTimestamp(),
		}, nil
	}
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
			petProfile.Gender = int(PetGenderMale)
		} else if *petProfileUpdatesInput.Gender == "Female" {
			petProfile.Gender = int(PetGenderFemale)
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
	gender := PetGenderIntToString(petProfile.Gender)
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
	gender_num := PetGenderMale
	if name == nil || *name == "" {
		errmsg += "name cannot be empty string."
	}
	if img == nil || *img == "" {
		img = &defaultPetImageUrl
	}
	if *gender == "Male" {
		gender_num = PetGenderFemale
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
		Gender:       int(gender_num),
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

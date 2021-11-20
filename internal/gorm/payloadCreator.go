package gorm

import (
	"context"
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
	usersProfile := make([]*model1.UserProfile, len(users))
	for i := 0; i < len(users); i++ {
		user := users[i]
		userGender := model1.UserGender(strconv.Itoa(user.Gender))
		usersProfile[i] = &model1.UserProfile{
			ID:       &user.Id,
			Name:     &user.Name,
			Gender:   &userGender,
			Birthday: nil,
			Email:    nil,
			Location: nil,
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
	return petsProfile
}

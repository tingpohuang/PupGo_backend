package gorm

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SQLCnter struct {
	gdb *gorm.DB
}

func (*SQLCnter) createUser() {
	user := User{
		Name: "test",
	}
	gdb.Create(&user)
}

func (s *SQLCnter) findUserByIdList(ctx context.Context, uid []uuid.UUID) (users User) {
	(*s.gdb).Where("id IN ? ", uid).Find(&users)
	return users
}

func (s *SQLCnter) findPetByOwner(ctx context.Context, uid uuid.UUID) (pets []Pet) {
	s.gdb.Joins("Company", s.gdb.Where(&Pet_owner{user_id: uid})).Find(&pets)
	return pets
}

// func (s *SQLCnter) findUsersByEvents(ctx context.Context) (user uuid.UUID) {

// }

// func (s *SQLCnter) findPetsByEvents(ctx context.Context) (pets []uuid.UUID) {

// }

func (s *SQLCnter) findEventsByUser(ctx context.Context) {

}

func (s *SQLCnter) findEventsByPets(ctx context.Context) {

}

func (S *SQLCnter) findEventsNearBy(ctx context.Context) {

}

func (S *SQLCnter) findConnection(ctx context.Context) {

}

func (S *SQLCnter) findPetRecommend(ctx context.Context) {

}

func (S *SQLCnter) removePets(ctx context.Context) {
}

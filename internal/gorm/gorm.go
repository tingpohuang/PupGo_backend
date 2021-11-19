package gorm

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewSQLCnter(gdb *gorm.DB) *SQLCnter {
	return &SQLCnter{
		gdb,
	}
}

type SQLCnter struct {
	gdb *gorm.DB
}

func (*SQLCnter) CreateUser() {
	id1, _ := uuid.New().MarshalBinary()
	id2, _ := uuid.New().MarshalBinary()
	user := User{
		Id:   id1,
		Name: "test",
	}
	pet := Pet{
		Id: id2,
	}
	petower := Pet_owner{
		User_id: id1,
		Pet_id:  id2,
	}

	gdb.Table("users").Create(&user)
	gdb.Table("pet").Create(&pet)
	gdb.Table("petowner").Create(&petower)

}

func (s *SQLCnter) findUserByIdList(ctx context.Context, uid []uuid.UUID) (users User) {
	(*s.gdb).Where("id IN ? ", uid).Find(&users)
	return users
}

func (s *SQLCnter) findPetByOwner(ctx context.Context, uid []byte) (pets []Pet) {
	s.gdb.Joins("Company", s.gdb.Where(&Pet_owner{User_id: uid})).Find(&pets)
	return pets
}

/*
func (s *SQLCnter) findUsersByEvents(ctx context.Context) (user uuid.UUID) {

}

func (s *SQLCnter) findPetsByEvents(ctx context.Context) (pets []uuid.UUID) {

}

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

*/

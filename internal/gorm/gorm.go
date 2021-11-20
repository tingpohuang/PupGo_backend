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
	id1 := uuid.NewString()
	id2 := uuid.NewString()
	// id1 := []byte("abcd")
	// id1, _ := uuid.New().MarshalBinary()
	// id2, _ := uuid.New().MarshalBinary()
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

func (s *SQLCnter) findUserByIdList(ctx context.Context, uid []string) (users []User) {
	(*s.gdb).Table("users").Where("id IN ? ", uid).Find(&users)
	return users
}

func (s *SQLCnter) findPetByIdList(ctx context.Context, pid []string) (pets []Pet) {
	(*s.gdb).Table("pet").Where("id IN ? ", pid).Find(&pets)
	return pets
}

func (s *SQLCnter) findPetByOwner(ctx context.Context, uid string) (pets []Pet) {
	s.gdb.Joins("Company", s.gdb.Where(&Pet_owner{User_id: uid})).Find(&pets)
	return pets
}

func (s *SQLCnter) FindPetRecommend(ctx context.Context, pid string) (petRecommend []Pet_recommend) {
	s.gdb.Table("pet_recommend").Where("id1 = ? OR id2 = ?", pid, pid).Find(&petRecommend)
	return petRecommend

}

func (s *SQLCnter) CreatePets(ctx context.Context, pet Pet) error {
	result := s.gdb.Table("pet").Create(&pet)
	return result.Error
}
func (s *SQLCnter) CreateUserPetRelation(ctx context.Context, uid string, pid string) error {
	result := s.gdb.Table("petowner").Create(&Pet_owner{
		User_id: uid,
		Pet_id:  pid,
	})
	return result.Error
}
func (s *SQLCnter) DeletePet(ctx context.Context, pid string) error {
	result := s.gdb.Table("pet").Delete(&Pet{}, pid)
	return result.Error
}

func (s *SQLCnter) CreatePetConnection(pid1 string, pid2 string) error {
	return nil
}

func (s *SQLCnter) FindPetById(ctx context.Context, pid string) (pets []Pet) {
	(*s.gdb).Table("pet").Where("id IN ? ", pid).Find(&pets)
	return pets
}

func (s *SQLCnter) UpdatePet(ctx context.Context, pet Pet) error {
	result := (*s.gdb).Table("pet").Model(&pet).Updates(&pet)
	return result.Error
}
func (s *SQLCnter) DeleteFriend(ctx context.Context, id1 string, id2 string) error {
	result := s.gdb.Table("pet_connection").Delete(&Pet_connection{id1: id1, id2: id2})
	return result.Error
}
func (s *SQLCnter) GetUserIdbyPetId(ctx context.Context, pid string) (*string, error) {
	p := Pet_owner{}
	result := s.gdb.Table("pet_owner").First(&p, "Pet_id = ?", pid)
	return &p.User_id, result.Error
}

/*
func (s *SQLCnter) findUsersByEvents(ctx context.Context) (user uuid.UUID) {

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

*/

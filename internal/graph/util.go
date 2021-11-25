package graph

import (
	"context"
	"fmt"
	"time"

	gorm "github.com/tingpo/pupgobackend/internal/gorm"
)

var (
	TIMEOUT int = 30
)

var sqlCnter *gorm.SQLCnter
var payloadCreator *gorm.PayloadCreator

func init() {
	mysqlConnector, err := gorm.GetConnectorFactory("mySQL")
	if err != nil {
		panic(fmt.Errorf("Connect to DB failed: %w \n", err))
	}

	db := mysqlConnector.NewDBConnection()
	sqlCnter = gorm.NewSQLCnter(db)
	payloadCreator = gorm.NewPayloadCreator(sqlCnter)

	//a := []string{"64c2d025-61d0-484f-ad0f-8df773d41a82"}
	//a := "170ed548-ebe4-46f1-8800-7f153e50e69f"
	//payloadCreator.GetPetListByUId(nil, a[0])
	//b, c := sqlCnter.FindOwnerByPIdList(nil, a)
	//println(len(b))
	//println(len(c))
	//	println(b[0].Id)
	//	println(c[0].Address)
	//println(b[0].Name)
	// println(b[0].Birthday.String())

}

func GetUidbyPid(ctx context.Context, pid string) (string, error) {
	ret, err := sqlCnter.GetUserIdbyPetId(ctx, pid)
	if err != nil {
		return "", err
	} else {
		return *ret, err
	}
}
func GetNowTimestamp() *string {
	tstmp := time.Now().String()
	return &tstmp
}
func RemoveFriend(ctx context.Context, id1 string, id2 string) error {
	if id1 > id2 {
		id1, id2 = id2, id1
	}
	err := sqlCnter.DeleteFriend(ctx, id1, id2)
	return err
}

var (
	defaultPetImageUrl     = "hahaha.jpg"
	defaultEventLimitPet   = 5
	defaultEventLimitHuman = 5
)

type EmptyValueError struct {
}

func (EmptyValueError) IsError() {}
func PetGenderIntToString(i int) *string {
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
func PetGenderStringToInt(s string) int {
	if s == "Male" {
		return 1
	} else if s == "Female" {
		return 2
	}
	return 0
}

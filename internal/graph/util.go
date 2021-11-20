package graph

import (
	"fmt"
	gorm "github.com/tingpo/pupgobackend/internal/gorm"
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

	//a := []string{"d080f320-a537-49b5-b0d5-b343d475caee", "d0afcc75-d24b-40e3-b8ab-2a68d9da26b1"}
	//b := sqlCnter.findUserByIdList(nil, a)
	//println(b[0].Id)
	//println(b[1].Id)

}

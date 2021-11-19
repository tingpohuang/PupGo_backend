package gorm

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

var (
	gdb *gorm.DB
	// m   = mysql.Config{}
)

func GetConnectorFactory(dbType string) (DBConnector, error) {
	if dbType == "mySQL" {
		return new(MySQLDBConnector), nil
	}
	return nil, errors.New("not vaild db type name")
}

type DBConnector interface {
	NewDBConnection() *gorm.DB
	// never closed connection until server down.
}
type MySQLDBConnector struct {
}

func (m MySQLDBConnector) NewDBConnection() *gorm.DB {

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	once.Do(func() { // <-- atomic, does not allow repeating
		// format dsn
		// DB_RESOURCE_NAME := os.Getenv("DB_RESOURCE_NAME")
		DB_USERNAME := os.Getenv("DB_USERNAME")
		DB_PASSWORD := os.Getenv("DB_PASSWORD")
		DB_URL := os.Getenv("DB_URL")
		DB_PORT := os.Getenv("DB_PORT")
		DB_TABLENAME := os.Getenv("DB_TABLENAME")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASSWORD, DB_URL, DB_PORT, DB_TABLENAME)
		dsn = "tim:greenfield204@tcp(192.168.1.129:3306)/pupgo?charset=utf8mb4&parseTime=True&loc=Local"
		tmpgdb, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{}) // <-- thread safe
		gdb = tmpgdb
		if err != nil {
			panic(err)
		}

	})
	return gdb
}

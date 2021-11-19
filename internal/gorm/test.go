package gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputValueToDatabse(t *testing.T) {
	// timestamp := time.Now()
	connector, err := GetConnectorFactory("mySQL")
	db := connector.NewDBConnection()
	assert := assert.New(t)
	assert.Nil(t, err)
	assert.NotNil(t, db)
	user := User{
		Name: "test",
		// Cooldown: time.Now(),
		gender: 1,
	}

	result := db.Create(&user) // pass pointer of data to Create
	assert.Nil(result.Error)
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tingpo/pupgobackend/internal/handler"
	"net/http"
)

func init() {
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.SetConfigType("json")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs/") // path to look for the config file in
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func main() {
	server := gin.Default()

	/*
		mysqlDBConnector, err := gorm.GetConnectorFactory("mySQL")

		if err != nil {
			panic(fmt.Errorf("Connect to DB failed: %w \n", err))
		}

		gdb := mysqlDBConnector.NewDBConnection()
		sql := gorm.NewSQLCnter(gdb)
		sql.CreateUser()


	*/

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	server.GET("/playground", handler.PlaygroundHandler())
	server.POST("/query", handler.GraphQLHandler())

	server.Run(viper.GetString("server.address"))

}

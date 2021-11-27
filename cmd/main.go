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
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})

	})
	server.GET("/playground", handler.PlaygroundHandler())
	server.POST("/query", handler.AuthRequired, handler.GraphQLHandler())
	server.GET("/users", handler.SignInHandler())
	server.POST("/users", handler.SignUpHandler())
	server.Run(viper.GetString("server.address"))

}

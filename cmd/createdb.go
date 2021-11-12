package createdb

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tingpo/pupgobackend/internal/handler"
	db "github.com/tingpo/pupgobackend/internal/db"
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
	db.createDB("testing")
}

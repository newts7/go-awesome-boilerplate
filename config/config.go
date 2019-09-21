package config

import (
	"fmt"
	"github.com/spf13/viper"
)
type Confuguration struct {
	Server struct {
		Port string
	}
}
var config Confuguration

func initDevConfig(){
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil{
		fmt.Println("Error in reading config file", err)
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil{
		fmt.Println("Error in unmarshalling config file", err)
		return
	}
}

func Init(environment string){
	if environment == "development" {
		initDevConfig()
	}
}

func GetConfig() Confuguration {
 	return config
}
package main

import (
	"fmt"
	"goBoilerPlate/config"
	"goBoilerPlate/db"
	"goBoilerPlate/logger"
	"goBoilerPlate/server"
	"os"
	
)

var environment string

func setEnviroment()  {
	environment = os.Getenv("environment")
	if environment == "" {
		fmt.Println("Error in getting environment variable")
		os.Exit(1)
	}
}
// @title Swagger Example API
func main(){
	setEnviroment()
	config.Init(environment)
	logger.Init()
	fmt.Println(logger.LogH.GetLevel())
	logger.LogH.WithField("environment", environment).Info("Setting Environment")
	db.Init()
	server.Init()
}

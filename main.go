package main

import (
	"fmt"
	"goBoilerPlate/config"
	"goBoilerPlate/db"
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

func main(){
	setEnviroment()
	fmt.Println("Setting Environment to ", environment)
	config.Init(environment)
	db.Init()
	server.Init()
}

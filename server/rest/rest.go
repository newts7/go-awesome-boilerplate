package rest

import (
	"fmt"
	"goBoilerPlate/config"
)

func Init(){
	r := NewRouter()
	err := r.Run(":"+config.GetConfig().Server.Port)
	if err != nil{
		fmt.Println("Error in initialising server", err)
	}
}
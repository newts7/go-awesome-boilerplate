package rest

import (
	"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"
	//"github.com/swaggo/gin-swagger"
	"goBoilerPlate/controllers"
)

func NewRouter()*gin.Engine{

	user := new (controllers.UserController)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/hello", user.Hello)
	return router
}
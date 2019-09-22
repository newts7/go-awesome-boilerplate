package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goBoilerPlate/config"
	"goBoilerPlate/controllers"
)

func NewRouter()*gin.Engine{

	user := new (controllers.UserController)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/* Swaager Integration */
	swaggerUrl := ginSwagger.URL("http://localhost:"+config.GetConfig().Server.Port+"/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	router.GET("/hello", user.Hello)
	return router
}
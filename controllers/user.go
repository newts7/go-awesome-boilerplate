package controllers

import (
	"github.com/gin-gonic/gin"
	_ "goBoilerPlate/docs"
	"net/http"
)

type UserController struct {}

// SayHello godoc
// @Summary Check if swagger working
// @Description Check if Swaager Working
// @Router /hello [get]
// @Success 200
func(h UserController)Hello(c *gin.Context){
	c.JSON(http.StatusOK, map[string]interface{}{
		"hello": "gopher",
	})
}
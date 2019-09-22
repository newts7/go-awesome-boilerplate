package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {

}


func(h UserController)Hello(c *gin.Context){
	c.String(http.StatusOK, "Hello World")
}
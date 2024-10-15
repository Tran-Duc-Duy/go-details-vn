package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {}

func NewUserController(router *gin.Engine) *UserController {
	return &UserController{}
}

//uc := &UserController{} lấy hai chữ cái đầu 
func (uc *UserController) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Duy": "code",
	})
}
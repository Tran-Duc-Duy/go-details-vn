package controller

import (
	"go-tip/internal/service"
	"go-tip/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(router *gin.Engine) *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

//uc := &UserController{} lấy hai chữ cái đầu 
func (uc *UserController) GetUser(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"Duy": uc.userService.GetUser(),
	// })

	// response.SuccessResponse(c, http.StatusOK, uc.userService.GetUser())
	response.ErrorResponse(c, http.StatusNotFound, uc.userService.GetUser())
}
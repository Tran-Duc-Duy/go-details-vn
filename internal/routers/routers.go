package routers

import (
	c "go-tip/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
		v1.GET("/user", c.NewUserController(r).GetUser)
	}
	return r
}

func Pong (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"users":[]string{
			"1",
			"2",
			"3",
		 },
	})
}
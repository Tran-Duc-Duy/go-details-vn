package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
		v1.GET("/hello", Pong)
		v1.POST("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
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
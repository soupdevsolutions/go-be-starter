package main

import (
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/hello", Hello)
	}

	return router
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

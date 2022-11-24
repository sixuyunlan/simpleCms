package main

import (
	"simpleCms/app/content"
	"simpleCms/app/user"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {

	group := router.Group("/", user.Fuck)

	group.GET("/content/:id", content.Get)

	group.DELETE("/content/:id", content.Delete)
	group.PUT("/content/:id", content.Put)

	group.POST("/content/", content.Post)

	router.POST("/user/login", user.Login)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

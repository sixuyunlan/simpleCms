package simpleCms

import (
	"simpleCms/app/content"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	router.GET("/content:id",	 content.Get)
	router.POST("/content", content.Post)
	router.PUT("/content:id", content.Put)
	router.DELETE("/content:id", conten	t.Delete)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
}

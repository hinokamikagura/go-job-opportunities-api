package router

import (
	// "os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := "8080"
	router.Run("0.0.0.0:" + port)
}

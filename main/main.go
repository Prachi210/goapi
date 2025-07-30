package main

import (
	"goapi/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Please proceed with a number for processing"})
	})
	router.GET("/:number", handler.ChunkHandler)
	router.NoRoute(func(c *gin.Context) {
		if c.Request.RequestURI == "/favicon.ico" {
			c.Status(204)
			return
		}
		c.JSON(404, gin.H{"error": "Not found"})
	})

	router.Run(":8080")
}

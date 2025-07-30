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

	router.Run(":8080")
}

package main

import (
	"fmt"
	"goapi/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Please proceed with a number for processing"})
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path != "/favicon.ico" {
			c.Next()
			fmt.Println("Path:", c.Request.URL.Path, "Status:", c.Writer.Status())
		} else {
			c.AbortWithStatus(http.StatusNoContent)
		}
	})
	router.GET("/:number", handler.ChunkHandler)
	router.Run(":8080")
}

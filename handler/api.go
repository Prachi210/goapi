package handler

import (
	"goapi/model"
	"goapi/service"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func ChunkHandler(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
		return
	}
	if number < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Give a Positive number"})
		return
	}
	chunks := service.ChunkNumbers(number)

	var wg sync.WaitGroup
	resultChan := make(chan int, len(chunks))

	startTime := time.Now()

	service.ProcessChunks(chunks, resultChan, &wg)
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	total := service.SumOfChunks(resultChan)
	elapsed := time.Since(startTime)
	c.JSON(http.StatusOK, gin.H{
		"message":    "Chunks processed successfully",
		"total_sum":  total,
		"number":     number,
		"chunks":     len(chunks),
		"chunk_size": model.ChunkSize,
		"time_taken": elapsed.String(),
	})
}

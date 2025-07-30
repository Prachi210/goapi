package handler

import (
	"encoding/json"
	"fmt"
	"goapi/model"
	"goapi/service"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockChunkNumbers(n int) [][]int {
	var chunks [][]int
	for i := 0; i < n; i += model.ChunkSize {
		end := i + model.ChunkSize
		if end > n {
			end = n
		}
		chunk := make([]int, end-i)
		for j := range chunk {
			chunk[j] = i + j + 1
		}
		chunks = append(chunks, chunk)
	}
	return chunks
}

func mockProcessChunks(chunks [][]int, resultChan chan int, wg *sync.WaitGroup) {
	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			sum := 0
			for _, v := range chunk {
				sum += v
			}
			resultChan <- sum
		}(chunk)
	}
}

func mockSumOfChunks(resultChan chan int) int {
	total := 0
	for v := range resultChan {
		total += v
	}
	return total
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/:number", ChunkHandler)
	return r
}

func TestChunkHandler_InvalidNumber(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid number", resp["error"])
}

func TestChunkHandler_NegativeNumber(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/-5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Give a Positive number", resp["error"])
}

func TestChunkHandler_ValidNumber(t *testing.T) {
	// Patch service functions for isolation
	origChunkNumbers := service.ChunkNumbers
	origProcessChunks := service.ProcessChunks
	origSumOfChunks := service.SumOfChunks
	service.ChunkNumbers = mockChunkNumbers
	service.ProcessChunks = mockProcessChunks
	service.SumOfChunks = mockSumOfChunks
	defer func() {
		service.ChunkNumbers = origChunkNumbers
		service.ProcessChunks = origProcessChunks
		service.SumOfChunks = origSumOfChunks
	}()

	router := setupRouter()
	number := 10
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%d", number), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Chunks processed successfully", resp["message"])
	assert.Equal(t, float64(number), resp["number"])
	assert.Equal(t, float64(len(mockChunkNumbers(number))), resp["chunks"])
	assert.Equal(t, float64(model.ChunkSize), resp["chunk_size"])
	assert.NotEmpty(t, resp["time_taken"])
	assert.NotNil(t, resp["total_sum"])
}

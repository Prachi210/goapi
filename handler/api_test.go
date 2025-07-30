package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/:number", ChunkHandler)
	return r
}

func TestChunkHandler_ValidNumber(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/12", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"total_sum":78`)
	assert.Contains(t, w.Body.String(), `"message":"Chunks processed successfully"`)
	assert.Contains(t, w.Body.String(), `"chunks":4`)
	assert.Contains(t, w.Body.String(), `"chunk_size":3`)
}

func TestChunkHandler_InvalidNumber(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/abc", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"Invalid number"`)
}

func TestChunkHandler_NegativeNumber(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/-5", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"Give a Positive number"`)
}

func TestChunkHandler_Zero(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/0", nil)
	w := httptest.NewRecorder()

	start := time.Now()
	router.ServeHTTP(w, req)
	elapsed := time.Since(start)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"total_sum":0`)
	assert.Contains(t, w.Body.String(), `"chunks":0`)
	assert.LessOrEqual(t, elapsed.Milliseconds(), int64(100))
}

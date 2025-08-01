package handler

import (
	"encoding/json"
	"goapi/model"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:number", ChunkHandler)
	return router
}

func TestChunkHandler_ValidInput(t *testing.T) {
	router := setupRouter()

	number := 250
	req, _ := http.NewRequest(http.MethodGet, "/"+strconv.Itoa(number), nil)
	rec := httptest.NewRecorder()

	start := time.Now()
	router.ServeHTTP(rec, req)
	duration := time.Since(start)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedSum := (number * (number + 1)) / 2
	assert.Equal(t, "Chunks processed successfully", response["message"])
	assert.Equal(t, float64(expectedSum), response["total_sum"])
	assert.Equal(t, float64(number), response["number"])
	assert.Equal(t, float64(3), response["chunks"])
	assert.Equal(t, float64(model.ChunkSize), response["chunk_size"])
	assert.NotEmpty(t, response["time_taken"])
	assert.LessOrEqual(t, parseDuration(t, response["time_taken"].(string)), duration)
}

func TestChunkHandler_InvalidInput(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/abc", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid number", response["error"])
}

func TestChunkHandler_NegativeInput(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/-10", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Give a Positive number", response["error"])
}

func parseDuration(t *testing.T, s string) time.Duration {
	d, err := time.ParseDuration(s)
	assert.NoError(t, err)
	return d
}

func TestChunkHandler_InvalidNumber(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "number", Value: "abc"}}

	ChunkHandler(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid number", resp["error"])
}

func TestChunkHandler_NegativeNumber(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "number", Value: "-5"}}

	ChunkHandler(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Give a Positive number", resp["error"])
}

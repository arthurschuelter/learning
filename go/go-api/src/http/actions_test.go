package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetHealth(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "The api is healthy")
}

func TestGetLinks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	Configure()
	router.GET("/links", GetLinks)

	req, err := http.NewRequest(http.MethodGet, "/links", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	expectedBody := `1`
	assert.JSONEq(t, expectedBody, rr.Body.String())
}

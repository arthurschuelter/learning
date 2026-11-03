package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	SetRoutes(router)
	Configure() // sets `test = 1`

	tests := []struct {
		path       string
		statusCode int
		expected   string
	}{
		{"/", http.StatusOK, "healthy"},
		{"/links", http.StatusOK, "1"},
	}

	for _, tc := range tests {
		req, _ := http.NewRequest(http.MethodGet, tc.path, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, tc.statusCode, w.Code, "Unexpected status for %s", tc.path)
		assert.Contains(t, w.Body.String(), tc.expected, "Unexpected body for %s", tc.path)
	}
}

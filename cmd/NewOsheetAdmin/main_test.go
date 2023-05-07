package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// End-to-End test
func TestHcRoute(t *testing.T) {
	router := setupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hc", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"OK\"}", w.Body.String())
}

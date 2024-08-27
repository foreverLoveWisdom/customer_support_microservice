package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/foreverLoveWisdom/customer_support_microservice.git/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String())
}

func TestCreateTicket(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter()

	// Create a request body
	requestBody := `{"title":"New Ticket","description":"Ticket description"}`
	req, _ := http.NewRequest(http.MethodPost, "/tickets", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Ticket created"}`, w.Body.String())
}

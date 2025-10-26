package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	r := gin.Default()
	r.GET("/user/:id", GetUser)

	req, _ := http.NewRequest(http.MethodGet, "/user/123", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]*User
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "123", response["user"].ID)
	assert.Equal(t, "John Doe", response["user"].Name)
	assert.Equal(t, "john.doe@example.com", response["user"].Email)
}

func TestCreateUser(t *testing.T) {
	r := gin.Default()
	r.POST("/user", CreateUser)

	user := User{ID: "123", Name: "John Doe", Email: "john.doe@example.com"}
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]*User
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "123", response["user"].ID)
	assert.Equal(t, "John Doe", response["user"].Name)
	assert.Equal(t, "john.doe@example.com", response["user"].Email)
}

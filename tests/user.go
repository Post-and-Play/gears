package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	infra.DevDatabaseConnect()
	UserMock()
	defer DeleteUserMock()

	r := RoutesSetup()
	r.GET("/users")

	path := fmt.Sprintf("/users?id=%s", strconv.Itoa(ID))
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var userMock models.User
	json.Unmarshal(resp.Body.Bytes(), &userMock)

	assert.Equal(t, "Jonas", userMock.Name, "Names should be the same")
	assert.Equal(t, "jonas_brothers", userMock.Mail, "Mails should be the same")
}

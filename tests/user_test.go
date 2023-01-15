package tests

import (
	"api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Valid(t *testing.T) {
	r := setupTest()

	input := models.CreateUserInput{
		Email:     "person@notexist.tld",
		FirstName: "Person",
		LastName:  "Notexist",
		Password:  "notasecurepassword",
	}
	jsonStr, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	var res models.ResData[models.BearerTokenResponse]
	json.Unmarshal(w.Body.Bytes(), &res)

	models.DB.Where("email = ?", input.Email).Delete(&models.User{})

	assert.Contains(t, res.Data.BearerToken, "Bearer ")
	assert.Equal(t, 200, w.Code)
}

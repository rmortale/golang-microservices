package domain

import (
	"net/http"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)
	assert.Nil(t, user, "not expeting user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}

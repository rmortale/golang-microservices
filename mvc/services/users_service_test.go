package services

import (
	"github.com/rmortale/golang-microservices/mvc/domain"
	"github.com/rmortale/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userid int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userid)
}

func TestGetUserNotFoundInDb(t *testing.T) {
	getUserFunction = func(userid int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)

}

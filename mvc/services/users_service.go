package services

import (
	"github.com/rmortale/golang-microservices/mvc/domain"
	"github.com/rmortale/golang-microservices/mvc/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}

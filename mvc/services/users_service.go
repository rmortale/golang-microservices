package services

import (
	"github.com/rmortale/golang-microservices/mvc/domain"
	"github.com/rmortale/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}

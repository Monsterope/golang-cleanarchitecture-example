package repositories

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/modules/requests"
)

type UserRepoInterface interface {
	FindById(userid int64) (*entities.User, error)
	FindUsername(reqdata *requests.LoginRequest) (*entities.User, error)
	Save(user *entities.User) error
	Edit(userid int64, reqdata *requests.UpdateUserRequest) error
}

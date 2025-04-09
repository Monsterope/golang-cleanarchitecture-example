package usecases

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/modules/requests"
)

type UserUsecaseInterface interface {
	FindAccountWithClaim(userid int64) (*entities.User, error)
	FindAccount(reqdata *requests.LoginRequest) (*entities.User, error)
	CreateAccount(reqdata *requests.CreateUserRequest) error
	UpdateAccount(userid int64, reqdata *requests.UpdateUserRequest) (*entities.User, error)
}

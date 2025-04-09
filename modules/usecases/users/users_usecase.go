package usecases

import (
	"cleanarchitecture-example/modules/entities"
	repositories "cleanarchitecture-example/modules/repositories/users"
	"cleanarchitecture-example/modules/requests"
	"cleanarchitecture-example/pkg/utils"
)

type UserUsecase struct {
	Repo repositories.UserRepoInterface
}

func NewUserUsecase(repo repositories.UserRepoInterface) UserUsecaseInterface {
	return &UserUsecase{Repo: repo}
}

func (uc *UserUsecase) FindAccountWithClaim(userid int64) (*entities.User, error) {
	user, err := uc.Repo.FindById(userid)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (uc *UserUsecase) FindAccount(reqdata *requests.LoginRequest) (*entities.User, error) {
	user, err := uc.Repo.FindUsername(reqdata)

	if err != nil {
		return nil, err
	}
	if checkPass := utils.CompareHasPassword(user.Password, reqdata.Password); checkPass != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) CreateAccount(reqdata *requests.CreateUserRequest) error {
	user := entities.User{
		Username: reqdata.Username,
		Password: utils.CreateHashPassword(reqdata.Password),
		Name:     reqdata.Name,
		UserType: "cust",
		Status:   1,
	}
	err := uc.Repo.Save(&user)
	return err
}

func (uc *UserUsecase) UpdateAccount(userid int64, reqdata *requests.UpdateUserRequest) (*entities.User, error) {
	err := uc.Repo.Edit(userid, reqdata)
	if err != nil {
		return nil, err
	}

	user, err := uc.Repo.FindById(userid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

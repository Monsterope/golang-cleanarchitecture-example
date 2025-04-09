package repositories

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/modules/requests"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/utils"
)

type UserRepo struct {
	Db *databases.DatabaseConfig
}

func NewUserRepo(db *databases.DatabaseConfig) UserRepoInterface {
	return &UserRepo{Db: db}
}

func (u *UserRepo) FindById(userid int64) (*entities.User, error) {
	dbuser := new(entities.User)
	result := u.Db.DB.Where("id = ?", userid).First(dbuser)
	return dbuser, result.Error

}

func (u *UserRepo) FindUsername(reqdata *requests.LoginRequest) (*entities.User, error) {
	dbuser := new(entities.User)
	result := u.Db.DB.Where("Username = ?", reqdata.Username).First(dbuser)
	return dbuser, result.Error
}

func (u *UserRepo) Save(user *entities.User) error {
	result := u.Db.DB.Create(&user)
	return result.Error
}

func (u *UserRepo) Edit(userid int64, reqdata *requests.UpdateUserRequest) error {
	updData := utils.CheckKeyIsHave(reqdata)
	updResult := u.Db.DB.Model(&entities.User{}).Where("id = ?", userid).Updates(updData)
	return updResult.Error
}

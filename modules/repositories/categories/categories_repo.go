package repositories

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/modules/requests"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/utils"
)

type CategoryRepo struct {
	Db *databases.DatabaseConfig
}

func NewCategoryRepo(db *databases.DatabaseConfig) CategoryRepoInterface {
	return &CategoryRepo{
		Db: db,
	}
}

func (repo *CategoryRepo) FindAll() ([]entities.Category, error) {
	categories := new([]entities.Category)
	result := repo.Db.DB.Find(categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return *categories, nil
}

func (repo *CategoryRepo) FindById(cateid int64) (entities.Category, error) {
	category := new(entities.Category)
	result := repo.Db.DB.Where("id = ?", cateid).First(category)
	return *category, result.Error

}

func (repo *CategoryRepo) Save(reqdata *requests.CategoryCreateRequest) error {
	result := repo.Db.DB.Create(reqdata.Item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CategoryRepo) Edit(cateid int64, reqdata *requests.CategoryUpdateRequest) error {
	updData := utils.CheckKeyIsHave(reqdata)
	updResult := repo.Db.DB.Model(&entities.Category{}).Where("id = ?", cateid).Updates(updData)
	return updResult.Error
}

package usecases

import (
	"cleanarchitecture-example/modules/entities"
	repositories "cleanarchitecture-example/modules/repositories/categories"
	"cleanarchitecture-example/modules/requests"
)

type CategoryUsecase struct {
	Repo repositories.CategoryRepoInterface
}

func NewCategoryUsecase(repo repositories.CategoryRepoInterface) CategoryUsecaseInterface {
	return &CategoryUsecase{Repo: repo}
}

func (repo CategoryUsecase) CreateCate(reqdata requests.CategoryCreateRequest) error {
	return repo.Repo.Save(&reqdata)
}

func (repo CategoryUsecase) GetCateAll() ([]entities.Category, error) {
	return repo.Repo.FindAll()
}

func (repo CategoryUsecase) UpdateCate(cateid int64, reqdata *requests.CategoryUpdateRequest) (*entities.Category, error) {
	result := repo.Repo.Edit(cateid, reqdata)

	if result != nil {
		return nil, result
	}

	category, err := repo.Repo.FindById(cateid)
	if err != nil {
		return nil, err
	}

	return &category, nil

}

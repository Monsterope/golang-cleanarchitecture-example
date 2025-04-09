package repositories

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/modules/requests"
)

type CategoryRepoInterface interface {
	FindAll() ([]entities.Category, error)
	FindById(cateid int64) (entities.Category, error)
	Save(reqdata *requests.CategoryCreateRequest) error
	Edit(cateid int64, reqdata *requests.CategoryUpdateRequest) error
}

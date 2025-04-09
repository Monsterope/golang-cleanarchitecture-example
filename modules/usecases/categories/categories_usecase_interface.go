package usecases

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/modules/requests"
)

type CategoryUsecaseInterface interface {
	CreateCate(reqdata requests.CategoryCreateRequest) error
	GetCateAll() ([]entities.Category, error)
	UpdateCate(cateid int64, reqdata *requests.CategoryUpdateRequest) (*entities.Category, error)
}

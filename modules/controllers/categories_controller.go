package controllers

import (
	"cleanarchitecture-example/modules/requests"
	"cleanarchitecture-example/modules/responses"
	usecases "cleanarchitecture-example/modules/usecases/categories"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	Usecase    usecases.CategoryUsecaseInterface
	DBConfig   *databases.DatabaseConfig
	RedisStore *utils.RedisAuthStore
	Validator  *validator.Validate
}

func NewCategoryController(uc usecases.CategoryUsecaseInterface, db *databases.DatabaseConfig, redis *utils.RedisAuthStore) *CategoryController {
	return &CategoryController{
		Usecase:    uc,
		DBConfig:   db,
		RedisStore: redis,
		Validator:  validator.New(),
	}
}

func (ctr *CategoryController) CreateCategory(c *fiber.Ctx) error {

	requestCre := new(requests.CategoryCreateRequest)
	if err := c.BodyParser(requestCre); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "bad request"))
	}
	if validate := ctr.Validator.Struct(requestCre); validate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", validate.Error()))
	}

	result := ctr.DBConfig.DB.Create(requestCre.Item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseFailureData("error", "Server error, please try again."))
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseSuccessData("success", "Created success."))
}

func (ctr *CategoryController) GetCategoryAll(c *fiber.Ctx) error {

	categories, err := ctr.Usecase.GetCateAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseFailureData("error", "Server error."))
	}

	var categoryResource []responses.CategoryResource
	for _, category := range categories {
		categoryResource = append(categoryResource, responses.GetCategoryResource(&category))
	}

	return c.JSON(ResponseSuccessData("success", categoryResource))
}

func (ctr *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	requestUpd := new(requests.CategoryUpdateRequest)
	if err := c.BodyParser(requestUpd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "bad request."))
	}
	if validate := ctr.Validator.Struct(requestUpd); validate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "bad request."))
	}

	cateIdParam := c.Params("cateid")

	cateid, err := strconv.ParseInt(cateIdParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", err.Error()))
	}

	category, err := ctr.Usecase.UpdateCate(cateid, requestUpd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseFailureData("error", "Server error."))
	}

	return c.JSON(ResponseSuccessData("success", category))

}

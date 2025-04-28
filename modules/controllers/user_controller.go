package controllers

import (
	"cleanarchitecture-example/modules/requests"
	"cleanarchitecture-example/modules/responses"
	usecases "cleanarchitecture-example/modules/usecases/users"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/middlewares"
	"cleanarchitecture-example/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Usecase    usecases.UserUsecaseInterface
	DBConfig   *databases.DatabaseConfig
	RedisStore *utils.RedisAuthStore
	Validator  *validator.Validate
}

func NewUserController(uc usecases.UserUsecaseInterface, db *databases.DatabaseConfig, redis *utils.RedisAuthStore) *UserController {
	return &UserController{
		Usecase:    uc,
		DBConfig:   db,
		RedisStore: redis,
		Validator:  validator.New(),
	}
}

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body requests.LoginRequest true "Login Data"
// @Success 200 {object} controllers.ResponseSuccessLogin
// @Failure 400 {object} controllers.ResponseFailure
// @Router /api/login [post]
func (ctr *UserController) Login(c *fiber.Ctx) error {
	requestUser := new(requests.LoginRequest)
	if err := c.BodyParser(requestUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "bad request"))
	}
	if validate := ctr.Validator.Struct(requestUser); validate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", validate.Error()))
	}

	user, err := ctr.Usecase.FindAccount(requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "User not found."))
	}

	resultToken := middlewares.Login(*requestUser, *user, ctr.RedisStore)

	if resultToken.Status != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", resultToken.Message))
	}

	responseData := ResponseSuccessLoginData("success", resultToken.Message, resultToken.Message2, responses.SafeModelCustomer(user))
	return c.JSON(responseData)
}

// @Summary Refresh token
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body middlewares.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} controllers.ResponseSuccessRefresh
// @Failure 400 {object} controllers.ResponseFailure
// @Router /api/refresh [post]
func (ctr *UserController) RefreshToken(c *fiber.Ctx) error {

	reqToken := new(middlewares.RefreshTokenRequest)

	if err := c.BodyParser(reqToken); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "bad request"))
	}

	resultToken := middlewares.RefreshToken(reqToken.RefreshToken, ctr.RedisStore)

	if resultToken.Status != 0 {
		return c.Status(resultToken.Status).JSON(ResponseFailureData("failure", resultToken.Message))
	}

	return c.JSON(ResponseSuccessRefreshData(resultToken.Message))

}

// @Summary Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body requests.CreateUserRequest true "Register Data"
// @Success 200 {object} controllers.ResponseSuccess
// @Failure 400 {object} controllers.ResponseFailure
// @Router /api/register [post]
func (ctr *UserController) Register(c *fiber.Ctx) error {
	requestRegister := new(requests.CreateUserRequest)

	if err := c.BodyParser(requestRegister); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "bad request"))
	}
	if validate := ctr.Validator.Struct(requestRegister); validate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", validate.Error()))
	}

	err := ctr.Usecase.CreateAccount(requestRegister)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseFailureData("error", "Server error, please try again."))
	}
	return c.Status(fiber.StatusCreated).JSON(ResponseSuccessData("success", fiber.Map{"message": "created success."}))
}

// @Summary Get users info
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} controllers.ResponseSuccess
// @Failure 401 {object} controllers.ResponseFailure
// @Failure 404 {object} controllers.ResponseFailure
// @Router /api/cust/user [get]
func (ctr *UserController) UserInfo(c *fiber.Ctx) error {
	claim := middlewares.GetClaim(c)
	if claim == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseFailureData("failure", "Unauthorzation"))
	}
	user, err := ctr.Usecase.FindAccountWithClaim(claim.ID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseFailureData("failure", err.Error()))
	}

	return c.JSON(ResponseSuccessData("success", responses.ModelUser(user)))
}

// @Summary Update user info
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Param user body requests.UpdateUserRequest false "User Data"
// @Success 200 {object} controllers.ResponseSuccess
// @Failure 400 {object} controllers.ResponseFailure
// @Failure 401 {object} controllers.ResponseFailure
// @Failure 500 {object} controllers.ResponseFailure
// @Router /api/cust/user/{userid} [put]
func (ctr *UserController) UpdateUser(c *fiber.Ctx) error {
	claim := middlewares.GetClaim(c)
	if claim == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseFailureData("failure", "Unauthorzation"))
	}

	requestUpd := new(requests.UpdateUserRequest)
	if err := c.BodyParser(requestUpd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", "bad request"))
	}
	if validate := ctr.Validator.Struct(requestUpd); validate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", validate.Error()))
	}
	userIdParam := c.Params("userid")
	userid, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFailureData("failure", err.Error()))
	}

	user, err := ctr.Usecase.UpdateAccount(userid, requestUpd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseFailureData("error", "Server error, please try again."))
	}

	return c.JSON(ResponseSuccessData("success", responses.ModelUser(user)))

}

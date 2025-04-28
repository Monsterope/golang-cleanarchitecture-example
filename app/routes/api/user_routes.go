package routes

import (
	"cleanarchitecture-example/modules/controllers"
	"cleanarchitecture-example/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	App         *fiber.App
	Controllers *controllers.UserController
	Middlewares *middlewares.RedisAuthMiddleware
}

func NewUserRoute(app *fiber.App, ctr *controllers.UserController, middleware *middlewares.RedisAuthMiddleware) *UserRoute {
	return &UserRoute{
		App:         app,
		Controllers: ctr,
		Middlewares: middleware,
	}
}

func (route *UserRoute) RouteUser() {
	api := route.App.Group("/api")

	controllers := route.Controllers
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)
	api.Post("/refresh", controllers.RefreshToken)

	customer := api.Group("cust", route.Middlewares.AuthIsCustomer)
	customer.Get("/user", controllers.UserInfo)
	customer.Put("/user/:userid", controllers.UpdateUser)

}

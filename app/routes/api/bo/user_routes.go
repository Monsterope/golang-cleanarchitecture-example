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
	middleware := route.Middlewares
	bo := api.Group("bo", middleware.AuthIsAdmin)

	bo.Get("/user", controllers.UserInfo)
	bo.Put("/user/:userid", controllers.UpdateUser)

}

package routes

import (
	"cleanarchitecture-example/modules/controllers"
	"cleanarchitecture-example/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

type CategoryRoute struct {
	App            *fiber.App
	Controllers    *controllers.CategoryController
	UserController *controllers.UserController
	Middlewares    *middlewares.RedisAuthMiddleware
}

func NewCategoryRoute(app *fiber.App, ctr *controllers.CategoryController, ctrUser *controllers.UserController, middleware *middlewares.RedisAuthMiddleware) *CategoryRoute {
	return &CategoryRoute{
		App:            app,
		Controllers:    ctr,
		UserController: ctrUser,
		Middlewares:    middleware,
	}
}

func (route *CategoryRoute) RouteBo() {
	api := route.App.Group("/api")
	controller := route.Controllers
	userController := route.UserController
	middleware := route.Middlewares

	bo := api.Group("bo", middleware.AuthIsAdmin)
	bo.Get("/user", userController.UserInfo)
	bo.Put("/user/:userid", userController.UpdateUser)

	bo.Get("/category", controller.GetCategoryAll)
	bo.Post("/category", controller.CreateCategory)
	bo.Put("/category/:cateid", controller.UpdateCategory)

}

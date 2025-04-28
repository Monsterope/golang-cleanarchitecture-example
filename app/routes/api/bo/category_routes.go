package routes

import (
	"cleanarchitecture-example/modules/controllers"
	"cleanarchitecture-example/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

type CategoryRoute struct {
	App         *fiber.App
	Controllers *controllers.CategoryController
	Middlewares *middlewares.RedisAuthMiddleware
}

func NewCategoryRoute(app *fiber.App, ctr *controllers.CategoryController, middleware *middlewares.RedisAuthMiddleware) *CategoryRoute {
	return &CategoryRoute{
		App:         app,
		Controllers: ctr,
		Middlewares: middleware,
	}
}

func (route *CategoryRoute) RouteCategory() {
	api := route.App.Group("/api")
	controller := route.Controllers
	middleware := route.Middlewares
	bo := api.Group("bo", middleware.AuthIsAdmin)

	bo.Get("/category", controller.GetCategoryAll)
	bo.Get("/category/:cateid", controller.GetCategory)
	bo.Post("/category", controller.CreateCategory)
	bo.Put("/category/:cateid", controller.UpdateCategory)

}

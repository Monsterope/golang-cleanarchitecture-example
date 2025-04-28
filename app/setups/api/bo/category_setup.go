package bo

import (
	routes "cleanarchitecture-example/app/routes/api/bo"
	"cleanarchitecture-example/modules/controllers"
	repositories "cleanarchitecture-example/modules/repositories/categories"
	usecases "cleanarchitecture-example/modules/usecases/categories"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/middlewares"
	"cleanarchitecture-example/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CategorySetupRoute(fiberConApp *fiber.App, dbConfig *databases.DatabaseConfig, authStoreInstance *utils.RedisAuthStore, middleware *middlewares.RedisAuthMiddleware) {
	category_repo := repositories.NewCategoryRepo(dbConfig)
	category_usecase := usecases.NewCategoryUsecase(category_repo)
	category_controller := controllers.NewCategoryController(category_usecase, dbConfig, authStoreInstance)
	routeCategory := routes.NewCategoryRoute(fiberConApp, category_controller, middleware)
	routeCategory.RouteCategory()
}

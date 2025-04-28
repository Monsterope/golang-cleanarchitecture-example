package bo

import (
	routes "cleanarchitecture-example/app/routes/api/bo"
	"cleanarchitecture-example/modules/controllers"
	repositories "cleanarchitecture-example/modules/repositories/users"
	usecases "cleanarchitecture-example/modules/usecases/users"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/middlewares"
	"cleanarchitecture-example/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func UserSetupRoute(fiberConApp *fiber.App, dbConfig *databases.DatabaseConfig, authStoreInstance *utils.RedisAuthStore, middleware *middlewares.RedisAuthMiddleware) {
	user_repo := repositories.NewUserRepo(dbConfig)
	user_usecase := usecases.NewUserUsecase(user_repo)
	user_controller := controllers.NewUserController(user_usecase, dbConfig, authStoreInstance)
	routeUser := routes.NewUserRoute(fiberConApp, user_controller, middleware)
	routeUser.RouteUser()
}

package setups

import (
	"cleanarchitecture-example/app/setups/api"
	"cleanarchitecture-example/app/setups/api/bo"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/middlewares"
	"cleanarchitecture-example/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteSetup(fiberConApp *fiber.App, dbConfig *databases.DatabaseConfig, authStoreInstance *utils.RedisAuthStore, middleware *middlewares.RedisAuthMiddleware) {
	api.UserSetupRoute(fiberConApp, dbConfig, authStoreInstance, middleware)
	bo.UserSetupRoute(fiberConApp, dbConfig, authStoreInstance, middleware)
	bo.CategorySetupRoute(fiberConApp, dbConfig, authStoreInstance, middleware)
}

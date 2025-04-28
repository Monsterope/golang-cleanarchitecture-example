package app

import (
	"cleanarchitecture-example/app/setups"
	"cleanarchitecture-example/configs"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/middlewares"
	"cleanarchitecture-example/pkg/utils"
	"fmt"

	_ "cleanarchitecture-example/docs"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	App *fiber.App
}

func NewApp() *App {
	fiberConApp := &App{
		App: fiber.New(),
	}

	fiberConApp.App.Use(cors.New())

	return fiberConApp

}

// @Title Clean architecture GO
// @Version 1.0
// @description This is example api docs project
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func (fiberConApp *App) Start(address string) {

	configs.Load()

	dbConfig := &databases.DatabaseConfig{}
	dbConfig.MysqlConnect()

	authStoreInstance := utils.NewRedisAuthStore(configs.GetEnv("redis.dns"))
	if authStoreInstance == nil {
		fmt.Println("Failed for connect Redis Auth")
	}
	middleware := middlewares.NewMiddlewareAuthRedis(authStoreInstance)

	// Swagger config
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}
	fiberConApp.App.Use(swagger.New(cfg))

	// register route
	setups.RouteSetup(fiberConApp.App, dbConfig, authStoreInstance, middleware)

	fiberConApp.App.Listen(address)
}

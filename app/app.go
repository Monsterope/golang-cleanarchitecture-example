package app

import (
	"cleanarchitecture-example/app/routes"
	"cleanarchitecture-example/configs"
	"cleanarchitecture-example/modules/controllers"
	repositoriesCate "cleanarchitecture-example/modules/repositories/categories"
	repositories "cleanarchitecture-example/modules/repositories/users"
	usecasesCate "cleanarchitecture-example/modules/usecases/categories"
	usecases "cleanarchitecture-example/modules/usecases/users"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/middlewares"
	"cleanarchitecture-example/pkg/utils"
	"fmt"

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

func (fiberConApp *App) Start(address string) {

	configs.Load()

	dbConfig := &databases.DatabaseConfig{}
	dbConfig.MysqlConnect()

	authStoreInstance := utils.NewRedisAuthStore(configs.GetEnv("redis.dns"))
	if authStoreInstance == nil {
		fmt.Println("Failed for connect Redis Auth")
	}
	middleware := middlewares.NewMiddlewareAuthRedis(authStoreInstance)

	user_repo := repositories.NewUserRepo(dbConfig)
	user_usecase := usecases.NewUserUsecase(user_repo)
	user_controller := controllers.NewUserController(user_usecase, dbConfig, authStoreInstance)
	routeUser := routes.NewUserRoute(fiberConApp.App, user_controller, middleware)
	routeUser.RouteApi()

	cate_repo := repositoriesCate.NewCategoryRepo(dbConfig)
	cate_usecase := usecasesCate.NewCategoryUsecase(cate_repo)
	cate_controller := controllers.NewCategoryController(cate_usecase, dbConfig, authStoreInstance)
	routeCate := routes.NewCategoryRoute(fiberConApp.App, cate_controller, user_controller, middleware)
	routeCate.RouteBo()

	fiberConApp.App.Listen(address)
}

package main

import (
	"cleanarchitecture-example/configs"
	"cleanarchitecture-example/modules/migrations"
	"cleanarchitecture-example/pkg/databases"
)

func main() {
	configs.Load()

	dbConfig := &databases.DatabaseConfig{}
	dbConfig.MysqlConnect()
	migrations.CreateUserTable(dbConfig)
	migrations.CreateCategoryTable(dbConfig)
	migrations.CreateProductTable(dbConfig)
}

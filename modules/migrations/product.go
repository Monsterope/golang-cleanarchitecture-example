package migrations

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/pkg/databases"
)

func CreateProductTable(db *databases.DatabaseConfig) {
	if !db.DB.Migrator().HasTable(entities.Product{}) {
		db.DB.AutoMigrate(entities.Product{})
	}
}

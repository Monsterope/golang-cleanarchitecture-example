package migrations

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/pkg/databases"
)

func CreateCategoryTable(db *databases.DatabaseConfig) {
	if !db.DB.Migrator().HasTable(entities.Category{}) {
		db.DB.AutoMigrate(entities.Category{})
	}
}

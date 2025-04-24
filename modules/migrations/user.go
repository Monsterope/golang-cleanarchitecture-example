package migrations

import (
	"cleanarchitecture-example/modules/entities"
	"cleanarchitecture-example/pkg/databases"
)

func CreateUserTable(db *databases.DatabaseConfig) {
	if !db.DB.Migrator().HasTable(entities.User{}) {
		db.DB.AutoMigrate(entities.User{})
	}
}

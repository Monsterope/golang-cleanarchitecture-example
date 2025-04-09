package databases

import "gorm.io/gorm"

type DatabaseConfig struct {
	DB *gorm.DB
}

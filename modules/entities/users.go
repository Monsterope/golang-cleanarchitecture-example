package entities

import (
	"time"
)

type User struct {
	Id        int64     `gorm:"primary_key;auto_increment"`
	Username  string    `gorm:"size256"`
	Password  string    `gorm:"size256"`
	Name      string    `gorm:"size256"`
	UserType  string    `gorm:"size8"`
	Status    int       `gorm:"size1;default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

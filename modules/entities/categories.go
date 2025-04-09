package entities

import "time"

type Category struct {
	Id        int64     `gorm:"primary_key;auto_increment"`
	Name      string    `gorm:"size256"`
	Product   []Product `gorm:"foreignkey:Id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

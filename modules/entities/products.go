package entities

import "time"

type Product struct {
	Id        int64     `gorm:"primary_key;auto_increment"`
	Code      string    `gorm:"type:varchar(255);unique;index"`
	Name      string    `gorm:"type:varchar(255);index"`
	Stock     int       `gorm:"size10"`
	Price     float64   `gorm:"type:decimal(10,2);default:0"`
	CateId    int64     `gorm:"not null"`
	Category  Category  `gorm:"foreignKey:CateId;references:Id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

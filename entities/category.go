package entities

import "time"

type Category struct {
	CategoryId		int `gorm:"autoIncrement;primaryKey"`
	CategoryName	string `gorm:"column:product_name"`
	Description		string `gorm:"column:description"`
	CreatedAt		time.Time `gorm:"column:created_at"`
	UpdatedAt		time.Time `gorm:"column:updated_at"`
}

func (Category) TableName() string {
	return "categories"
}
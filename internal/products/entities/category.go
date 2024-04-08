package entities

import "time"

type Category struct {
	Id          int `gorm:"autoIncrement;primaryKey"`
	CategoryId  int `gorm:"column:category_id"`
	CategoryName        string `gorm:"column:product_name"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt	time.Time `gorm:"column:updated_at"`
}

func (Category) TableName() string {
	return "category"
}
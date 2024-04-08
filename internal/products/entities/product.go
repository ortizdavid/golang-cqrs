package entities

import "time"

type Product struct {
	Id          int `gorm:"autoIncrement;primaryKey"`
	CategoryId  int `gorm:"column:category_id"`
	ProductName        string `gorm:"column:product_name"`
	Code        string `gorm:"column:code"`
	Price       float32 `gorm:"column:price"`
	Description string `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt	time.Time `gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "product"
}

type ProductData struct {
}
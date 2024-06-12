package entities

import "time"

type Product struct {
	ProductId       int64 `gorm:"autoIncrement;primaryKey"`
	CategoryId  	int `gorm:"column:category_id"`
	ProductName		string `gorm:"column:product_name"`
	Code			string `gorm:"column:code"`
	UnitPrice       float64 `gorm:"column:unit_price"`
	Description 	string `gorm:"column:description"`
	UniqueId   		time.Time `gorm:"column:unique_id"`
	CreatedAt   	time.Time `gorm:"column:created_at"`
	UpdatedAt		time.Time `gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "products"
}

type ProductData struct {
	ProductId		int64
	UniqueId		string
	ProductName		string
	Code			string
	UnitPrice		float64
	Image			string
	Description		string
	CreatedAt		string
	UpdatedAt		string
	CategoryId		int
	CategoryName	string	
}

package repositories

import (
	"github.com/ortizdavid/golang-cqrs/config"
	"github.com/ortizdavid/golang-cqrs/entities"
)

type ProductRepository struct {
}

// --------CREATE --------------------
func (repo *ProductRepository) Create(product entities.Product) error {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// --------UPDATE --------------------
func (repo *ProductRepository) Update(product entities.Product) error {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// --------GET All--------------------
func (repo *ProductRepository) GetAll() ([]entities.Product, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var products []entities.Product
	result := db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// ------- Get by id and uniqueid
func (repo *ProductRepository) GetById(id int) (entities.Product, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var product entities.Product
	result := db.Find(&product)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetByUniqueId(uniqueId string) (entities.Product, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var product entities.Product
	result := db.First(&product, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

// --------DELETE --------------------
func (repo *ProductRepository) Delete(product entities.Product) error {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Delete(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// -------Exists ------------------------
func (repo *ProductRepository) Exists(productName string) bool {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var product entities.Product
	result := db.Where("product_name=?", productName).Find(&product)
	if result.Error != nil {
		return false
	}
	return product.ProductId != 0
}

//----------Others-----------------------
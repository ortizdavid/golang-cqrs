package repositories

import (
	"github.com/ortizdavid/golang-cqrs/config"
	"github.com/ortizdavid/golang-cqrs/entities"
)

type CategoryRepository struct {
}

// --------CREATE --------------------
func (repo *CategoryRepository) Create(category entities.Category) error {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// --------UPDATE --------------------
func (repo *CategoryRepository) Update(category entities.Category) error {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// --------GET All--------------------
func (repo *CategoryRepository) GetAll() ([]entities.Category, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var categories []entities.Category
	result := db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// ------- Get by id and uniqueid
func (repo *CategoryRepository) GetById(id int) (entities.Category, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var category entities.Category
	result := db.Find(&category)
	if result.Error != nil {
		return entities.Category{}, result.Error
	}
	return category, nil
}

func (repo *CategoryRepository) GetByUniqueId(uniqueId string) (entities.Category, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var category entities.Category
	result := db.First(&category, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.Category{}, result.Error
	}
	return category, nil
}

// --------DELETE --------------------
func (repo *CategoryRepository) Delete(category entities.Category) error {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Delete(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// -------Exists ------------------------
func (repo *CategoryRepository) Exists(categoryName string) bool {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var category entities.Category
	result := db.Where("category_name=?", categoryName).Find(&category)
	if result.Error != nil {
		return false
	}
	return category.CategoryId != 0
}

//----------Others-----------------------
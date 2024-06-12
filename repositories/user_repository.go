package repositories

import (
	"github.com/ortizdavid/golang-cqrs/config"
	"github.com/ortizdavid/golang-cqrs/entities"
)

type UserRepository struct {
}

//--------CREATE --------------------
func (repo *UserRepository) Create(user entities.User) error {
	db, _:= config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//--------UPDATE --------------------
func (repo *UserRepository) Update(user entities.User) error {
	db, _:= config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//--------GET All--------------------
func (repo *UserRepository) GetAll() ([]entities.User, error) {
	db, _:= config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

//------- Get by id and uniqueid
func (repo *UserRepository) GetById(id int) (entities.User, error) {
	db, _:= config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.Find(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetByUniqueId(uniqueId string) (entities.User, error) {
	db, _:= config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

//--------DELETE --------------------
func (repo *UserRepository) Delete(user entities.User) error {
	db, _:= config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//-------Exists ------------------------
func (repo *UserRepository) Exists(userName string) bool {
	db, _:= config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.Where("user_name=?", userName).Find(&user)
	if result.Error != nil {
		return false
	}
	return user.UserId != 0
}

//----------Others-----------------------
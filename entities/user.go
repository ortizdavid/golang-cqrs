package entities

import "time"

type User struct {
	UserId    	int `gorm:"autoIncrement;primarykey"`
	RoleId  	int `gorm:"column:role_id"`
	UserName  	string `gorm:"column:user_name"`
	Password  	string `gorm:"column:password"`
	IsActive  	bool `gorm:"column:is_active"`
	Image  		string `gorm:"column:image"`
	UniqueId  	string `gorm:"column:unique_id"`
	Token  		string `gorm:"column:token"`
	CreatedAt 	time.Time `gorm:"column:created_at"`
	UpdatedAt  	time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

type UserData struct {
	UserId			int 
	UniqueId 		string
	Token 			string
	UserName 		string
	Password 		string
	Active   		bool
	Image   		string
	CreatedAt 		string
	UpdatedAt  		string
	RoleId 			int
	RoleName 		string
	RoleCode 		string
}

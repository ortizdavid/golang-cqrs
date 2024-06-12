package entities

type Role struct {
	RoleId		int `gorm:"primaryKey;autoIncrement"`
	RoleName	string `gorm:"column:role_name"`
	Code		string `gorm:"column:code"`
}

func (Role) TableName() string {
	return "roles"
}
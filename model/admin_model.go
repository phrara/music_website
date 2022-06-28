package model

type Admin struct {
	Id int `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
}

func (a Admin) TableName() string{
	return "admin"
}

func NewAdmin(name, password string) *Admin {
	return &Admin{
		Name: name,
		Password: password,
	}
}
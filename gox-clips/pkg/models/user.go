package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string 
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

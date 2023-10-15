package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"varchar(100);unique" json:"username"`
	Email    string `gorm:"varchar(100);unique" json:"email"`
	Password string `gorm:"varchar(300)" json:"password"`
}

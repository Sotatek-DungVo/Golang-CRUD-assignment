package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Posts    []Post `json:"posts" gorm:"foreignKey:UserID"`
}

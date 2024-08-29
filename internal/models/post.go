package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

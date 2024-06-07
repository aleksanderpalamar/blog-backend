package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Password string    `json:"password"`
	Comments []Comment `json:"comments" gorm:"foreignKey:UserID"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID   uint      `json:"post_id"`
	UserID   uint      `json:"user_id"`
	Content  string    `json:"content"`
	ParentID *uint     `json:"parent_id,omitempty"`
	Replies  []Comment `json:"replies" gorm:"foreignKey:ParentID"`
}

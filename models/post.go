package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  []string  `json:"content" gorm:"type:json"`
	Likes    int       `json:"likes"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

// Implement serializer interface for serializing to JSON
func (p *Post) BeforeSave(tx *gorm.DB) (err error) {
	if p.Content != nil {
		content, err := json.Marshal(p.Content)
		if err != nil {
			return err
		}
		p.Content = []string{string(content)}
	}
	return
}

func (p *Post) AfterFind(tx *gorm.DB) (err error) {
	if len(p.Content) > 0 {
		var content []string
		if err := json.Unmarshal([]byte(p.Content[0]), &content); err != nil {
			return err
		}
		p.Content = content
	}
	return
}

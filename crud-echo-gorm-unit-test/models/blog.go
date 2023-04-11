package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  uint   `json:"userid" form:"userid" gorm:"column:user_id"`
	User    User
}

func (blog *Blog) TableName() string {
	return "blog"
}

func (blog Blog) ToString() string {
	return fmt.Sprintf("id: %d\ntitle: %s\ncontent: %s", blog.ID, blog.Title, blog.Content)
}

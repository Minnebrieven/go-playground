package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID  uint   `json:"userid" form:"userid"`
}

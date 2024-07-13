package models

import (
    "gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    Content  string `gorm:"type:text;not null" json:"content"`
	UserID   uint  `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"user"`
	ThreadID uint  `json:"thread_id"`
	Upvotes int `gorm:"default:0" json:"upvotes"`
}

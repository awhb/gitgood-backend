package models

import (
    "gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    Content  string `gorm:"type:text;not null" json:"content"`
	UserID   uint  
	ThreadID uint  
	Upvotes int `gorm:"default:0"`
}

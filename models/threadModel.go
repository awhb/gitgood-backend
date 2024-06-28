package models

import (
    "gorm.io/gorm"
)

type Thread struct {
    gorm.Model
    Title   string	`gorm:"unique;not null"`
    Content string	`gorm:"type:text;not null"`
    UserID  uint
    User    User
	Comments []Comment
    Tags    []Tag `gorm:"many2many:thread_tags;"`
	Upvotes int `gorm:"default:0"`
}

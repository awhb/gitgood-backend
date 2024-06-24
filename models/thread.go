package models

import (
    "gorm.io/gorm"
)

type Thread struct {
    gorm.Model
    Title   string
    Content string
    UserID  uint
    User    User
    Tags    []Tag `gorm:"many2many:thread_tags;"`
}

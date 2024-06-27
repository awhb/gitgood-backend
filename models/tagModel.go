package models

import (
    "gorm.io/gorm"
)

type Tag struct {
    gorm.Model
    Name    string `gorm:"unique;not null"`
    Threads []Thread `gorm:"many2many:thread_tags;"`
}

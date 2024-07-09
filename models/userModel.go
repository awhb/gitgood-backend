package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username    string `gorm:"unique;not null" json:"username"`
    Password string   `gorm:"not null" json:"-"`
	Threads  []Thread   `gorm: "foreignKey:UserID, constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments []Comment   `gorm: "foreignKey:UserID, constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

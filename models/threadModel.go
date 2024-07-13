package models

import (
    "gorm.io/gorm"
	"github.com/lib/pq"
)

type Thread struct {
    gorm.Model
    Title     string	      `gorm:"unique;not null" json:"title"`
    Content   string		  `gorm:"type:text;not null" json:"content"`
	UserID    uint            `json:"user_id"`
	User      User            `gorm:"foreignKey:UserID" json:"user"`
    Tags      pq.StringArray  `gorm:"type:text[]" json:"tags"`
    Comments  []Comment       `gorm:"foreignKey:ThreadID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comments"`
	Upvotes   int             `gorm:"default:0" json:"upvotes"`
}

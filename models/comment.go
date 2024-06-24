// models/comment.go
package models

import (
    "gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    Content  string
    ThreadID uint
    Thread   Thread
    UserID   uint
    User     User
}

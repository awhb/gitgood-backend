package models

import (
    "gorm.io/gorm"
	"time"
)

type User struct {
    gorm.Model
    Username    string `gorm:"unique;not null"`
    Password string   `gorm:"not null"`
	Threads  []Thread  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Comments []Comment  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type UserResponse struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func MapUserToResponse(user User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

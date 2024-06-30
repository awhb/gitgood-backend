package models

import (
	"time"
    "gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    Content  string `gorm:"type:text;not null"`
    UserID   uint
    User     User    `gorm:"foreignKey:UserID"`
    ThreadID uint
    Thread   Thread   `gorm:"foreignKey:ThreadID"`
	Upvotes int `gorm:"default:0"`
}

type CommentResponse struct {
    Content   string      `json:"content"`
    User      UserResponse `json:"user"`
    CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func MapCommenttoResponse(comment Comment) CommentResponse {
	return CommentResponse{
		Content:   comment.Content,
		User:      MapUserToResponse(comment.User),
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func MapCommentsToResponse(comments []Comment) []CommentResponse {
	var commentResponses []CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, MapCommenttoResponse(comment))
	}
	return commentResponses
}

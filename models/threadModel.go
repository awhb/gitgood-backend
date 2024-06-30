package models

import (
	"time"
    "gorm.io/gorm"
)

type Thread struct {
    gorm.Model
    Title   string	`gorm:"unique;not null"`
    Content string	`gorm:"type:text;not null"`
    UserID    uint
    User      User       `gorm:"foreignKey:UserID"`
    Tags      []Tag      `gorm:"many2many:thread_tags;"`
    Comments  []Comment
	Upvotes int `gorm:"default:0"`
}

type ThreadResponse struct {
    Title     string          `json:"title"`
    Content   string          `json:"content"`
    User      UserResponse    `json:"user"`
    Tags      []TagResponse   `json:"tags"`
    Comments  []CommentResponse `json:"comments"`
    CreatedAt time.Time       `json:"created_at"`
    UpdatedAt time.Time       `json:"updated_at"`
}

func MapThreadToResponse(thread Thread) ThreadResponse {
	return ThreadResponse{
		Title:     thread.Title,
		Content:   thread.Content,
		User:      MapUserToResponse(thread.User),
		Tags:      MapTagsToResponse(thread.Tags),
		Comments:  MapCommentsToResponse(thread.Comments),
		CreatedAt: thread.CreatedAt,
		UpdatedAt: thread.UpdatedAt,
	}
}

func MapThreadsToResponse(threads []Thread) []ThreadResponse {
	var threadResponses []ThreadResponse
	for _, thread := range threads {
		threadResponses = append(threadResponses, MapThreadToResponse(thread))
	}
	return threadResponses
}

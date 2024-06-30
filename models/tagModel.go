package models

type Tag struct {
	ID        uint `gorm:"primarykey"`
    Name    string `gorm:"unique;not null"`
    Threads []Thread `gorm:"many2many:thread_tags;"`
}

type TagResponse struct {
	Name string `json:"name"`
}

func MapTagToResponse(tag Tag) TagResponse {
	return TagResponse{
		Name: tag.Name,
	}
}

func MapTagsToResponse(tags []Tag) []TagResponse {
	var tagResponses []TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses, MapTagToResponse(tag))
	}
	return tagResponses
}

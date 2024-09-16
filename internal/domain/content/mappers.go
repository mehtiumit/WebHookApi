package content

import (
	"github.com/google/uuid"
	"time"
	"webhook/internal/domain/entities"
)

func ToCreateContentRequestDto(request CreateContentRequest) CreateContentRequestDto {
	return CreateContentRequestDto{
		Title: request.Title,
		Body:  request.Body,
	}
}

func ToContentEntity(request CreateContentRequestDto) entities.Content {
	return entities.Content{
		ID:        uuid.NewString(),
		Title:     request.Title,
		Body:      request.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToContentDto(entity *entities.Content) ContentDto {
	return ContentDto{
		ID:    entity.ID,
		Title: entity.Title,
		Body:  entity.Body,
	}
}

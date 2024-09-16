package contentApi

import (
	"context"
	"webhook/internal/domain/content"
	"webhook/internal/mongodb"
	"webhook/pkg/log"
	"webhook/pkg/models"
)

type ContentService interface {
	CreateContent(ctx context.Context, content content.CreateContentRequestDto) (content.CreateContentResponse, error)
	GetContent(ctx context.Context, id string) (content.ContentDto, error)
}

type contentService struct {
	contentRepo mongodb.ContentRepository
	logger      *log.Logrus
}

func (c contentService) CreateContent(ctx context.Context, contentReq content.CreateContentRequestDto) (content.CreateContentResponse, error) {
	contentIsExist, err := c.contentRepo.IsContentExist(ctx, contentReq.Title)
	if err != nil {
		return content.CreateContentResponse{}, err
	}
	if contentIsExist {
		return content.CreateContentResponse{}, models.CustomError{
			Code:        409001,
			ErrorDetail: "content with this title already exists",
		}
	}
	contentEntity := content.ToContentEntity(contentReq)
	err = c.contentRepo.Create(ctx, contentEntity)
	if err != nil {
		return content.CreateContentResponse{}, err
	}
	return content.CreateContentResponse{ID: contentEntity.ID}, nil
}

func (c contentService) GetContent(ctx context.Context, id string) (content.ContentDto, error) {
	contentEntity, err := c.contentRepo.GetContent(ctx, id)
	if err != nil {
		return content.ContentDto{}, err
	}
	return content.ToContentDto(contentEntity), nil
}

func NewContentService(contentRepo mongodb.ContentRepository, logger *log.Logrus) ContentService {
	return &contentService{contentRepo: contentRepo, logger: logger}
}

package hookApi

import (
	"context"
	"webhook/internal/domain/entities"
	"webhook/internal/domain/hook"
	"webhook/internal/mongodb"
	"webhook/pkg/log"
	"webhook/pkg/models"
)

type HookService interface {
	CreateHook(ctx context.Context, hookDto hook.CreateHookRequestDto) (hook.CreateHookResponse, error)
}

type hookService struct {
	hookRepository    mongodb.HookRepository
	contentRepository mongodb.ContentRepository
	logger            *log.Logrus
	// hookRepo is the repository that the service will use to interact with the database
}

func (s hookService) CreateHook(ctx context.Context, hookDto hook.CreateHookRequestDto) (hook.CreateHookResponse, error) {
	// Check if the hook already exists
	content, validationErr := s.CheckContents(ctx, hookDto)
	if validationErr != nil {
		return hook.CreateHookResponse{}, validationErr
	}
	// Create a new hook from the DTO
	newHook := hook.ToHookEntity(hookDto, *content)
	// Insert the new hook into the database
	saveErr := s.hookRepository.CreateHook(ctx, newHook)
	if saveErr != nil {
		return hook.CreateHookResponse{}, saveErr
	}
	// Return the response
	return hook.CreateHookResponse{
		Message:   "Hook created successfully",
		MessageId: newHook.ID,
	}, nil
}

func (s hookService) CheckContents(ctx context.Context, hookDto hook.CreateHookRequestDto) (*entities.Content, error) {
	hookOnDb, err := s.contentRepository.GetContent(ctx, hookDto.ContentId)
	if err != nil {
		return nil, models.CustomError{
			Code:        404001,
			ErrorDetail: "content not found",
		}
	}
	if hookOnDb == nil {
		return nil, models.CustomError{
			Code:        404001,
			ErrorDetail: "content not found",
		}
	}

	hookIsExist, err := s.hookRepository.CheckIsSendBefore(ctx, hookDto.ContentId, hookDto.To)
	if err != nil {
		return nil, err
	}
	if hookIsExist {
		return nil, models.CustomError{
			Code:        409001,
			ErrorDetail: "hook with this contentId and to already exists",
		}
	}
	return hookOnDb, nil
}

func NewHookService(hookRepository mongodb.HookRepository, contentRepository mongodb.ContentRepository, logger *log.Logrus) HookService {
	return &hookService{
		hookRepository:    hookRepository,
		logger:            logger,
		contentRepository: contentRepository,
	}
}

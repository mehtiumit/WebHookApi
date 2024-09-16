package hook

import (
	"github.com/google/uuid"
	"time"
	"webhook/internal/domain/entities"
	"webhook/internal/domain/models"
)

func ToCreateHookRequestDto(createHookRequest CreateHookRequest) CreateHookRequestDto {
	return CreateHookRequestDto{
		To:        createHookRequest.To,
		ContentId: createHookRequest.ContentId,
	}
}

func ToHookEntity(createHookRequest CreateHookRequestDto, content entities.Content) entities.Hook {
	return entities.Hook{
		ID:        uuid.NewString(),
		To:        createHookRequest.To,
		ContentId: createHookRequest.ContentId,
		Body:      content.Body,
		Status:    models.HookStatusInitial,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func HookEtityToHookDto(hook entities.Hook) HookDto {
	return HookDto{
		To:      hook.To,
		Content: hook.Body,
	}
}

func HookEntityToHookRedisModel(hook entities.Hook) entities.HookRedisModel {
	return entities.HookRedisModel{
		HookId: hook.ID,
		SentAt: time.Now(),
	}
}

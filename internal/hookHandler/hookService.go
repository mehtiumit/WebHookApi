package hookHandler

import (
	"context"
	"webhook/internal/clients"
	hook2 "webhook/internal/domain/hook"
	"webhook/internal/mongodb"
	"webhook/internal/redis"
	"webhook/pkg/log"
)

type HookHandlerService interface {
	SendHookMessages() error
}

type hookService struct {
	hookServiceRepository mongodb.HookRepository
	hookApiClient         clients.HookApiClient
	redisClient           redis.RedisClient
	logger                *log.Logrus
}

func (s hookService) SendHookMessages() error {
	ctx := context.Background()
	hooks, err := s.hookServiceRepository.GetInitialHooks(ctx)
	if err != nil {
		return err
	}
	if len(hooks) == 0 {
		s.logger.Info("No hooks to send. waiting for new hooks...")
		return nil
	}
	for _, hook := range hooks {
		hookDto := hook2.HookEtityToHookDto(hook)
		sendErr := s.hookApiClient.SendHookMessage(hookDto)
		if sendErr != nil {
			s.logger.Errorf("Hook with id %s failed to send. | Error: %s", hook.ID, sendErr.Error())
		}
		statusErr := s.hookServiceRepository.SetHookStatus(ctx, hook.ID, "sent")
		if statusErr != nil {
			s.logger.Errorf("Hook with id %s failed to update status. | Error: %s", hook.ID, statusErr.Error())
		}
		redisSetErr := s.redisClient.Set(ctx, hook.ID, hook2.HookEntityToHookRedisModel(hook))
		if redisSetErr != nil {
			s.logger.Errorf("Hook with id %s failed to set redis. | Error: %s", hook.ID, redisSetErr.Error())
		}
	}
	return nil
}

func NewHookHandlerService(hookServiceRepository mongodb.HookRepository,
	hookApiClient clients.HookApiClient, redisClient redis.RedisClient,
	logger *log.Logrus) HookHandlerService {
	return &hookService{
		hookServiceRepository: hookServiceRepository,
		hookApiClient:         hookApiClient,
		redisClient:           redisClient,
		logger:                logger,
	}
}

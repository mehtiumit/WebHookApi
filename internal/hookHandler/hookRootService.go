package hookHandler

import (
	"time"
	"webhook/pkg/log"
)

type HookRootService interface {
	Start() error
}

type hookRootService struct {
	hookService HookHandlerService
	logger      *log.Logrus
}

func (h hookRootService) Start() error {
	h.logger.Info("Hook root service is starting...")
	for {
		err := h.hookService.SendHookMessages()
		if err != nil {
			return err
		}
		h.logger.Info("Hook root service is sleeping... for 2 minutes")
		time.Sleep(2 * time.Second)
	}
}

func NewHookRootService(logger *log.Logrus, hookService HookHandlerService) HookRootService {
	return &hookRootService{logger: logger, hookService: hookService}
}

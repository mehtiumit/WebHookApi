package hookApi

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"webhook/internal/domain/hook"
	"webhook/pkg/log"
)

type HookHandler struct {
	service HookService
	logger  *log.Logrus
}

func NewHookHandler(e *echo.Echo, service HookService, logger *log.Logrus) {
	h := &HookHandler{
		service: service,
		logger:  logger,
	}
	g := e.Group("/webhook/api/v1/hook")
	g.POST("", h.CreateHook)
}

// CreateHook
// @Summary      Create hook
// @Description  Create hook
// @Tags         Hook
// @Accept       json
// @Produce      json
// @Param        createHookRequest body hook.CreateHookRequest true "Create new hook"
// @Success      202                    {int}	int
// @Failure      400                    {object}  error
// @Failure      500                    {object}  error
// @Router       /v1/hook [post]
func (h HookHandler) CreateHook(c echo.Context) error {
	var createHookRequest = new(hook.CreateHookRequest)
	var ctx = c.Request().Context()
	if err := c.Bind(createHookRequest); err != nil {
		return c.JSON(http.StatusBadRequest, createHookRequest)
	}

	if validationErr := createHookRequest.Validate(); validationErr != nil {
		return c.JSON(http.StatusBadRequest, validationErr)
	}
	createHookResponse, serviceErr := h.service.CreateHook(ctx, hook.ToCreateHookRequestDto(*createHookRequest))
	if serviceErr != nil {
		return c.JSON(http.StatusBadRequest, serviceErr)
	}
	return c.JSON(http.StatusOK, createHookResponse)
}

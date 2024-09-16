package hook

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"webhook/internal/domain/models"
	models2 "webhook/pkg/models"
)

func (receiver CreateHookRequest) Validate() error {
	if receiver.To == "" {
		return echo.NewHTTPError(400, models2.CustomError{
			Code:        400001,
			ErrorDetail: "to field is required",
		})
	}
	if receiver.ContentId == "" {
		return echo.NewHTTPError(400, models2.CustomError{
			Code:        400002,
			ErrorDetail: "content field is required",
		})
	}
	if len(receiver.ContentId) > 1000 {
		return echo.NewHTTPError(400, models2.CustomError{
			Code:        400003,
			ErrorDetail: "content field is too long",
		})
	}
	if receiver.Action == "" {
		return echo.NewHTTPError(400, models2.CustomError{
			Code:        400004,
			ErrorDetail: "action field is required",
		})
	}
	if !models.IsActionValid(receiver.Action) {
		return echo.NewHTTPError(400, models2.CustomError{
			Code:        400005,
			ErrorDetail: fmt.Sprintf("action field is invalid value it must be one of %s or %s", models.ActionStart, models.ActionEnd),
		})
	}
	return nil
}

package echoExtension

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"webhook/pkg/models"
)

func ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			code := http.StatusInternalServerError
			var message interface{}

			if httpErr, ok := err.(*echo.HTTPError); ok {
				code = httpErr.Code
				message = httpErr.Message
			} else if customErr, ok := err.(models.CustomError); ok {
				code = http.StatusConflict
				message = customErr
			} else {
				message = models.CustomError{
					Code:        5000,
					ErrorDetail: "Internal Server Error",
				}
			}
			return c.JSON(code, message)
		}
		return nil
	}
}

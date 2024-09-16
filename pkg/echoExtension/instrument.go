package echoExtension

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

// RegisterGlobalMiddlewares Root Level (After router)
// The following built-in middleware should be registered at this level:
// BodyLimit
// Logger
// Gzip
// Recover
// ServerHeader middleware adds a `Server` header to the response.
func RegisterGlobalMiddlewares(e *echo.Echo, routePrefix string, statusCodes map[string]string) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: func(context echo.Context) bool {
			return MySkipper(context, routePrefix)
		},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodTrace, http.MethodConnect},
	}))
	e.Use(middleware.BodyLimit("1M"))
	addCollerationId(e)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(context echo.Context) bool {
			return MySkipper(context, routePrefix)
		},
		Level: -1,
	}))

	e.Use(middleware.Recover())
	e.Use(ErrorMiddleware)
}

func MySkipper(context echo.Context, routePrefix string) bool {
	if strings.HasPrefix(context.Path(), routePrefix+"/status") ||
		strings.HasPrefix(context.Path(), routePrefix+"/swagger") ||
		strings.HasPrefix(context.Path(), routePrefix+"/metrics") {
		return true
	}
	return false
}

func addCollerationId(e *echo.Echo) {
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			uid, _ := uuid.NewRandom()
			return uid.String()
		},
	}))
}

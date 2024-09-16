package contentApi

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"webhook/internal/domain/content"
	"webhook/pkg/log"
)

type ContentHandler struct {
	ContentService ContentService
	logger         *log.Logrus
}

func NewContentHandler(e *echo.Echo, contentService ContentService, logger *log.Logrus) ContentHandler {
	handler := ContentHandler{ContentService: contentService, logger: logger}
	g := e.Group("/webhook/api/v1/content")
	g.POST("", handler.CreateContent)
	g.GET("/:id", handler.GetContent)
	return handler
}

// CreateContent
// @Summary Create content
// @Description Create content
// @Tags Content
// @Accept json
// @Produce json
// @Param content body content.CreateContentRequestDto true "Content"
// @Success 200 {object} content.CreateContentResponse
// @Failure 400 {object} models.CustomError
// @Router /v1/content [post]
func (c ContentHandler) CreateContent(ctx echo.Context) error {
	request := new(content.CreateContentRequest)
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	response, err := c.ContentService.CreateContent(ctx.Request().Context(), content.ToCreateContentRequestDto(*request))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, response)
}

// GetContent
// @Summary Get content
// @Description Get content
// @Tags Content
// @Accept json
// @Produce json
// @Param id path string true "Content ID"
// @Success 200 {object} content.ContentDto
// @Failure 400 {object} models.CustomError
// @Router /v1/content/{id} [get]
func (c ContentHandler) GetContent(ctx echo.Context) error {
	id := ctx.Param("id")
	response, err := c.ContentService.GetContent(ctx.Request().Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(200, response)
}

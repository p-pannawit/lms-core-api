package health

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Checker interface {
	Check(context.Context) error
}

type Handler struct {
	service Checker
}

func NewHandler(service Checker) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) Check(c *gin.Context) {
	if err := handler.service.Check(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "unavailable",
			"database": "unavailable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"database": "ok",
	})
}

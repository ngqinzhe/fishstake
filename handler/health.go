package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/fishstake/dal/model"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Handle(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Health{
			Server:   "online",
			Metrics:  "online",
			Database: "online",
		})
	}
}

package handler

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/fishstake/consts"
	"github.com/ngqinzhe/fishstake/dal/model"
)

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (r *RootHandler) Handle(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, model.Root{
			Version:    consts.Version,
			Date:       time.Now().Unix(),
			Kubernetes: false,
		})
	}
}

package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/fishstake/dal/db"
	"github.com/ngqinzhe/fishstake/dal/model"
	"github.com/ngqinzhe/fishstake/util"
)

type ValidateHandler struct {
	DBClient db.MongoDBClient
}

func NewValidateHandler(db db.MongoDBClient) *ValidateHandler {
	return &ValidateHandler{
		DBClient: db,
	}
}

func (v *ValidateHandler) Handle(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.ValidateIPRequest{}
		if err := c.BindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, model.HttpError{
				Message: "request body invalid",
			})
			return
		}
		c.JSON(http.StatusOK, model.ValidateIPResponse{
			Status: util.IsValidIPV4(req.Ip),
		})
	}
}

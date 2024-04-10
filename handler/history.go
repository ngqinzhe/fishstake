package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/fishstake/dal/db"
	"github.com/ngqinzhe/fishstake/dal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HistoryHandler struct {
	DBClient db.MongoDBClient
}

func NewHistoryHandler(dbClient db.MongoDBClient) *HistoryHandler {
	return &HistoryHandler{
		DBClient: dbClient,
	}
}

func (h *HistoryHandler) Handle(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		options := options.Find()
		options.SetSort(bson.D{{"created_at", -1}})
		options.SetLimit(20)
		queries, err := h.DBClient.FindQueries(ctx, options)
		if err != nil {
			// TODO: log
			c.JSON(http.StatusBadRequest, model.HttpError{
				Message: "server error",
			})
			return
		}
		c.JSON(http.StatusOK, queries)
	}
}

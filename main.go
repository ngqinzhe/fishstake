package main

import (
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/fishstake/consts"
	"github.com/ngqinzhe/fishstake/dal/db"
	"github.com/ngqinzhe/fishstake/handler"
	"github.com/ngqinzhe/fishstake/metrics"
)

func main() {
	ctx := context.Background()
	db := db.Init()
	defer db.Close(ctx)

	server := newServer(ctx, gin.Default(), db)
	server.Init()
}

type server struct {
	ctx    context.Context
	router *gin.Engine
	db     db.MongoDBClient
}

func newServer(ctx context.Context, router *gin.Engine, db db.MongoDBClient) *server {
	return &server{
		ctx:    ctx,
		router: router,
		db:     db,
	}
}

func (s *server) Init() {
	s.router.Use(cors.Default())

	metrics.Init(s.router)
	s.initRoutes()
}

func (s *server) initRoutes() {
	// get
	s.router.GET("/", handler.NewRootHandler().Handle(s.ctx))
	s.router.GET("/health", handler.NewHealthHandler().Handle(s.ctx))
	s.router.GET(fmt.Sprintf("/%s/history", consts.Version), handler.NewHistoryHandler(s.db).Handle(s.ctx))
	s.router.GET(fmt.Sprintf("/%s/tools/lookup", consts.Version), handler.NewLookupHandler(s.db).Handle(s.ctx))

	// post
	s.router.POST(fmt.Sprintf("/%s/tools/validate", consts.Version), handler.NewValidateHandler(s.db).Handle(s.ctx))

	s.router.Run("localhost:3000")
}

package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(router *gin.Engine) {
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPingRouter(router *gin.Engine) {
	group := router.Group("")
	group.GET("/ping", pingHandler)
	group.GET("/", pingHandler)
}

func pingHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

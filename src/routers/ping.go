package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qsoyq/go-proxy-api/src/constants"
)

var RUN_AT_TS = time.Now().Unix()
var RUN_AT = (time.Now()).Format(time.DateTime)

type PingOutputScheme struct {
	Message   string `json:"message" form:"message" binding:"required" example:"pong"`
	Timestamp int64  `json:"timestamp" form:"timestamp" binding:"required" example:"1736303120"`
	Current   string `json:"current" form:"current" binding:"required" example:"2025-01-08 10:25:20"`
	RunAtTs   int64  `json:"run_at_ts" form:"run_at_ts" binding:"required" example:"1736303120"`
	RunAt     string `json:"run_at" form:"run_at" binding:"required" example:"2025-01-08 10:25:20"`
	Version   string `json:"version" form:"version" binding:"required" example:"0.1.0"`
}

func AddPingRouter(router *gin.Engine) {
	group := router.Group("")
	group.GET("/ping", pingHandler)
	group.GET("/", pingHandler)
}

func pingHandler(ctx *gin.Context) {
	current := time.Now()
	timestamp := current.Unix()
	formattedTime := current.Format(time.DateTime)
	output := PingOutputScheme{Message: "pong", Timestamp: timestamp, Current: formattedTime, RunAtTs: RUN_AT_TS, RunAt: RUN_AT, Version: constants.VERSION}
	ctx.JSON(http.StatusOK, output)
}

// @id		    ping.get
// @Summary		Ping
// @Description	Ping
// @Tags
// @Accept			json
// @Produce 		json
// @Router			/ping [get]
// @Success      200  {object}  PingOutputScheme
func _(c *gin.Context) {}

// @id		    /.get
// @Summary		Ping
// @Description	Ping
// @Tags
// @Accept			json
// @Produce 		json
// @Router			/ [get]
// @Success      200  {object}  PingOutputScheme
func _(c *gin.Context) {}

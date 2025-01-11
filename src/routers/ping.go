package routers

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qsoyq/go-proxy-api/src/constants"
)

var RUN_AT_TS = time.Now().Unix()
var RUN_AT = (time.Now()).Format(time.DateTime)

type PingDocsScheme struct {
	// 指向 Redoc 文档格式的页面
	Redoc string `json:"redoc" binding:"required" example:"/redoc"`
	// 指向 Swagger UI 文档格式的页面
	Docs string `json:"docs" binding:"required" example:"/docs/index.html"`
	// 域名信息
}

type DomainResultScheme struct {
	Domain string   `json:"domain"`
	Addrs  []string `json:"addrs"`
}

type PingOutputScheme struct {
	// 保留字段
	Message string `json:"message" form:"message" binding:"required" example:"pong"`
	// 当前时间戳
	Timestamp int64 `json:"timestamp" form:"timestamp" binding:"required" example:"1736303120"`
	// 当前日期时间字符串
	Current string `json:"current" form:"current" binding:"required" example:"2025-01-08 10:25:20"`
	// 服务启动时的时间戳
	RunAtTs int64 `json:"run_at_ts" form:"run_at_ts" binding:"required" example:"1736303120"`
	// 服务启动时的日期时间字符串
	RunAt string `json:"run_at" form:"run_at" binding:"required" example:"2025-01-08 10:25:20"`
	// 版本号
	Version string `json:"version" form:"version" binding:"required" example:"0.1.0"`
	// 接口文档
	Docs PingDocsScheme `json:"docs" form:"docs" binding:"required"`
	// 域名信息
	Domains []DomainResultScheme `json:"domains" binding:"required"`
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
	docs := PingDocsScheme{Redoc: "/redoc", Docs: "/docs/index.html"}
	output := PingOutputScheme{Docs: docs, Message: "pong", Timestamp: timestamp, Current: formattedTime, RunAtTs: RUN_AT_TS, RunAt: RUN_AT, Version: constants.VERSION}
	domainEnv := os.Getenv("PingDomains")
	if domainEnv != "" {
		domainList := strings.Split(domainEnv, ",")
		for _, domain := range domainList {
			if addrList, err := net.LookupHost(domain); err == nil {
				output.Domains = append(output.Domains, DomainResultScheme{Domain: domain, Addrs: addrList})
			} else {
				fmt.Printf("[Ping] can't resolve domain: %s, err: %s\n", domain, err.Error())
			}
		}
	}

	ctx.JSON(http.StatusOK, output)
}

// @id			ping.get
// @Summary		Ping
// @Description	Ping
// @Tags
// @Accept		json
// @Produce	json
// @Router		/ping [get]
// @Success	200	{object}	PingOutputScheme
func _(c *gin.Context) {}

// @id				/.get
// @Summary		Ping
// @Description	Ping
// @Tags
// @Accept		json
// @Produce	json
// @Router		/ [get]
// @Success	200	{object}	PingOutputScheme
func _(c *gin.Context) {}

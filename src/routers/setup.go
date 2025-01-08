package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qsoyq/go-proxy-api/src/routers/convert/svg"
	"github.com/qsoyq/go-proxy-api/src/routers/convert/xml"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	AddPingRouter(r)
	svg.AddSvgRouter(r)
	xml.AddXMLRouter(r)
	return r
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qsoyq/go-proxy-api/src/routers/apple/location"
	"github.com/qsoyq/go-proxy-api/src/routers/convert/svg"
	"github.com/qsoyq/go-proxy-api/src/routers/convert/xml"
	"github.com/qsoyq/go-proxy-api/src/routers/webhook/twilio/sms"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	AddPingRouter(r)
	svg.AddSvgRouter(r)
	xml.AddXMLRouter(r)
	sms.AddTwilioSmsRouter(r)
	location.AddLocaltionRouter(r)
	return r
}

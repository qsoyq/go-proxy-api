package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	_ "github.com/mvrilo/go-redoc/_examples/gen/docs"

	"github.com/qsoyq/go-proxy-api/src/routers"
	"github.com/qsoyq/go-proxy-api/src/routers/convert/xml"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func redocHandler(doc redoc.Redoc) gin.HandlerFunc {
	handler := doc.Handler()
	return func(ctx *gin.Context) {
		handler(ctx.Writer, ctx.Request)
		if ctx.Writer.Status() == http.StatusOK {
			if strings.HasSuffix(ctx.Request.URL.Path, doc.SpecPath) {
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}

func setupOpenAPI(r *gin.Engine) {
	doc := redoc.Redoc{
		Title:       "go-proxy-api",
		Description: "",
		SpecFile:    "./docs/swagger.json", // "./openapi.yaml"
		SpecPath:    "/openapi.json",       // "/openapi.yaml"
		DocsPath:    "/redoc",
	}
	r.Use(redocHandler(doc))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
		c.URL = "openapi.json"
	}))
}

func setup() *gin.Engine {
	r := routers.SetupRouter()
	setupOpenAPI(r)
	xml.AddXMLRouter(r)
	return r
}

// @title	go-proxy-api
// @version	1.0
// @description
// @BasePath	/api
func main() {
	r := setup()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Listening on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	_ "github.com/mvrilo/go-redoc/_examples/gen/docs"

	"github.com/qsoyq/go-proxy-api/src/middleware"
	"github.com/qsoyq/go-proxy-api/src/routers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CommandLine struct {
	SwaggerPath        string
	RedirectIfNotFound bool
	RedirectScheme     string
	RedirectHost       string
}

var cmd CommandLine

func setupOpenAPI(r *gin.Engine) {
	doc := redoc.Redoc{
		Title:       "go-proxy-api",
		Description: "",
		SpecFile:    cmd.SwaggerPath, // "./openapi.yaml"
		SpecPath:    "/openapi.json", // "/openapi.yaml"
		DocsPath:    "/redoc",
	}
	r.Use(middleware.RedocHandler(doc))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
		c.URL = "openapi.json"
	}))
}

func setup() *gin.Engine {
	r := routers.SetupRouter()
	setupOpenAPI(r)
	if cmd.RedirectIfNotFound {
		r.Use(middleware.HTTPNotFoundRedirectHandler(cmd.RedirectScheme, cmd.RedirectHost))
	}
	return r
}

//	@title		go-proxy-api
//	@version	0.1.2
//	@description
//	@BasePath	/api
func main() {
	flag.StringVar(&cmd.SwaggerPath, "swagger", "./src/docs/swagger.json", "swagger json path")
	flag.BoolVar(&cmd.RedirectIfNotFound, "redirect", false, "是否启用 404 下自动重定向")
	flag.StringVar(&cmd.RedirectHost, "redirect-host", "", "http or https")
	flag.StringVar(&cmd.RedirectScheme, "redirect-scheme", "", "domain or ip address")

	flag.Parse()

	r := setup()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Listening on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}

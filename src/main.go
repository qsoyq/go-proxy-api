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

var swaggerPath string

func setupOpenAPI(r *gin.Engine) {
	fmt.Println("swagger path: ", swaggerPath)
	doc := redoc.Redoc{
		Title:       "go-proxy-api",
		Description: "",
		SpecFile:    swaggerPath,     // "./openapi.yaml"
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
	return r
}

// @title	go-proxy-api
// @version	1.0
// @description
// @BasePath	/api
func main() {
	_swaggerPath := flag.String("swagger", "./src/docs/swagger.json", "swagger json path")
	flag.Parse()
	swaggerPath = *_swaggerPath
	r := setup()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Listening on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}

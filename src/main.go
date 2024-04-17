package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/qsoyq/go-proxy-api/src/routers"
	"github.com/qsoyq/go-proxy-api/src/routers/convert/xml"
)

func setup() *gin.Engine {
	r := routers.SetupRouter()
	xml.AddXMLRouter(r)
	return r
}

func main() {
	r := setup()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Listening on port %s", port)
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

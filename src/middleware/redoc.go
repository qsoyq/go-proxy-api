package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	_ "github.com/mvrilo/go-redoc/_examples/gen/docs"
)

func RedocHandler(doc redoc.Redoc) gin.HandlerFunc {
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

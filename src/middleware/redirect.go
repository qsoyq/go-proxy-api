package middleware

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	_ "github.com/mvrilo/go-redoc/_examples/gen/docs"
)

func HTTPNotFoundRedirectHandler(scheme string, host string) gin.HandlerFunc {
	if scheme != "http" && scheme != "https" {
		panic("[HTTPNotFoundRedirectHandler] invalid scheme")
	}
	if len(host) == 0 {
		panic("[HTTPNotFoundRedirectHandler] invalid host")
	}
	fmt.Printf("[HTTPNotFoundRedirectHandler] enable http not found redirect, scheme: %s, host: %s\n", scheme, host)
	return func(ctx *gin.Context) {
		ctx.Next()
		if ctx.Writer.Status() == http.StatusNotFound {
			fmt.Println("status status 404: ", ctx.Request.URL.Path)

			currentPath := ctx.Request.URL.Path
			currentQuery := ctx.Request.URL.RawQuery

			// 构建新的 URL
			newURL := url.URL{
				Scheme:   scheme,
				Host:     host,
				Path:     currentPath,
				RawQuery: currentQuery,
			}
			ctx.Redirect(http.StatusTemporaryRedirect, newURL.String())
			fmt.Println("redirect from ", ctx.Request.URL.String(), "to ", newURL.String())
			ctx.Abort()
		}
	}
}

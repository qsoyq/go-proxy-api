package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	_ "github.com/mvrilo/go-redoc/_examples/gen/docs"
	"github.com/qsoyq/go-proxy-api/src/constants"
)

func RedocHandler(doc redoc.Redoc) gin.HandlerFunc {
	handler := doc.Handler()
	return func(ctx *gin.Context) {

		var buf bytes.Buffer
		writer := &responseWriter{ResponseWriter: ctx.Writer, buffer: &buf}

		// 针对 openapi.json 拦截响应
		if strings.HasSuffix(ctx.Request.URL.Path, doc.SpecPath) {
			ctx.Writer = writer
		}

		handler(ctx.Writer, ctx.Request)

		if strings.HasSuffix(ctx.Request.URL.Path, doc.SpecPath) {
			// 动态修改 API 版本号
			bodyBytes, err := io.ReadAll(&buf)
			if err != nil {
				return
			}
			var response map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &response); err == nil {
				// 在这里修改 JSON 响应
				if info, ok := response["info"].(map[string]interface{}); ok {
					info["version"] = constants.VERSION
				}
				modifiedBody, _ := json.Marshal(response)
				writer.ResponseWriter.Write(modifiedBody)
			}
			ctx.Writer = writer.ResponseWriter
			ctx.Abort()
		}
		ctx.Next()
	}
}

type responseWriter struct {
	gin.ResponseWriter
	buffer  *bytes.Buffer
	ifWrite bool
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.buffer.Write(b)
	rw.ifWrite = true
	return 0, nil
}

func (rw *responseWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
}

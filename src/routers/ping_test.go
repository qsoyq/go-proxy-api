package routers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qsoyq/go-proxy-api/src/routers"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	router := routers.SetupRouter()

	// GET 请求测试
	{
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	}
}

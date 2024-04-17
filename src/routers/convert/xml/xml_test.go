package xml

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qsoyq/go-proxy-api/src/routers"

	"github.com/stretchr/testify/assert"
)

func TestConvertXMLToJSONRoute(t *testing.T) {
	router := routers.SetupRouter()
	AddXMLRouter(router)

	xmlStr := "<person><name>John Doe</name><age>30</age></person>"
	// GET 请求测试
	{
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/convert/xml/json", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 422, w.Code)
	}

	{
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/convert/xml/json", nil)

		q := req.URL.Query()
		q.Add("content", xmlStr)
		req.URL.RawQuery = q.Encode()

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var resp map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			assert.Fail(t, "解析 json 失败", w.Body.String())
		}
		_, ok := resp["content"]
		assert.True(t, ok, w.Body.String())

		content, _ := resp["content"].(string)

		var m map[string]interface{}
		if err := json.Unmarshal([]byte(content), &m); err != nil {
			assert.Fail(t, "解析 json 失败", w.Body.String())
		}

		if _, ok := m["person"]; !ok {
			assert.Fail(t, fmt.Sprintf("%v", m))
		}

		person, ok := m["person"].(map[string]interface{})
		assert.True(t, ok, w.Body.String())
		if ok {
			name, ok := person["name"].(string)
			assert.True(t, ok, w.Body.String())
			assert.Equal(t, name, "John Doe")
		}
	}

	// POST 请求测试
	{
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/convert/xml/json", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 422, w.Code)
	}

	{
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/convert/xml/json", nil)

		q := req.URL.Query()
		q.Add("content", xmlStr)
		req.URL.RawQuery = q.Encode()

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var resp map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			assert.Fail(t, "解析 json 失败", w.Body.String())
		}
		_, ok := resp["content"]
		assert.True(t, ok, w.Body.String())

		content, _ := resp["content"].(string)

		var m map[string]interface{}
		if err := json.Unmarshal([]byte(content), &m); err != nil {
			assert.Fail(t, "解析 json 失败", w.Body.String())
		}

		if _, ok := m["person"]; !ok {
			assert.Fail(t, fmt.Sprintf("%v", m))
		}

		person, ok := m["person"].(map[string]interface{})
		assert.True(t, ok, w.Body.String())
		if ok {
			name, ok := person["name"].(string)
			assert.True(t, ok, w.Body.String())
			assert.Equal(t, name, "John Doe")
		}
	}
}

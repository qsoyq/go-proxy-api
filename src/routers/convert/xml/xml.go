package xml

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sbabiv/xml2map"

	"github.com/gin-gonic/gin"
)

type ConvertXMLInput struct {
	Content string `json:"content" form:"content"`
}

func AddXMLRouter(router *gin.Engine) {
	group := router.Group("/api/convert/xml")
	group.GET("/json", toJsonGet)
	group.POST("/json", toJsonGet)
}

func toJsonGet(c *gin.Context) {
	var input ConvertXMLInput
	switch c.Request.Method {
	case "GET":
		c.ShouldBindQuery(&input)
	case "POST":
		c.ShouldBindJSON(&input)

	}

	if len(input.Content) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "content must not be empty"})
		return
	}

	m, err := xml2map.NewDecoder(strings.NewReader(input.Content)).Decode()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": fmt.Sprintf("Error unmarshalling XML: %s", err)})
		return
	}

	output, err := json.Marshal(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": fmt.Sprintf("序列化失败: %s", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"content": string(output)})
}

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
	// xml 字符串
	Content string `json:"content" form:"content" binding:"required" example:"<note><to>value</to></note>"`
}

type ConvertXMLOutput struct {
	// json 字符串
	Content string `json:"content" form:"content" binding:"required" example:"{\"note\": {\"to\": \"value\"}}"`
}

func AddXMLRouter(router *gin.Engine) {
	group := router.Group("/api/convert/xml")
	group.GET("/json", toJson)
	group.POST("/json", toJson)
}

func toJson(c *gin.Context) {
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

// @id				convert.xml.json.get
// @Summary		XML to JSON
// @Description	将传入的 xml 字符串转成 json字符串并返回
// @Tags			convert
// @Accept			json
// @Produce 		json
// @Param			content	query	string	true	"xml字符串"
// @Router			/convert/xml/json [get]
// @Success      200  {object}  ConvertXMLOutput
func _(c *gin.Context) {}

// @id				convert.xml.json.post
// @Summary		XML to JSON
// @Description	将传入的 xml 字符串转成 json字符串并返回
// @Tags			convert
// @Accept			json
// @Produce 		json
// @Param			content	body	ConvertXMLInput true "-"
// @Router			/convert/xml/json [post]
// @Success      200  {object}  ConvertXMLOutput
func _(c *gin.Context) {}

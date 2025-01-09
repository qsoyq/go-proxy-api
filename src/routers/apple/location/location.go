package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qsoyq/go-proxy-api/src/errors"
)

func AddLocaltionRouter(router *gin.Engine) {
	group := router.Group("/api/apple/location")
	group.GET("/:code", code)
}

func code(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		ctx.JSON(http.StatusUnprocessableEntity, errors.BadEntity("path.code", "missing code", "value_error.code.scheme"))
		return
	}
	ctx.String(http.StatusOK, code)
}

// @id			apple.location.code
// @Summary		Location Code
// @Description	返回 code 文本
// @Tags		Apple
// @Produce		text/plain
// @Param		code	path	string	true	"地区代码"	Example(US)
// @Router		/apple/location/{code} [get]
// @Success		200	{string}	string	地区代码
func _(c *gin.Context) {}

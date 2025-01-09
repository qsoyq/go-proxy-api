package svg

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"strconv"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qsoyq/go-proxy-api/src/errors"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func AddSvgRouter(router *gin.Engine) {
	group := router.Group("/api/convert/svg")
	group.GET("/png", toPng)
}

func svgToPng(svgData []byte) ([]byte, error) {
	w, h := 512, 512
	in := bytes.NewReader(svgData)

	icon, _ := oksvg.ReadIconStream(in)
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	var buf bytes.Buffer
	err := png.Encode(&buf, rgba)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func toPng(ctx *gin.Context) {
	url := strings.TrimSpace(ctx.Query("url"))
	download, err := strconv.ParseBool(ctx.Query("download"))
	if err != nil {
		download = false
	}

	if url == "" {
		ctx.JSON(http.StatusUnprocessableEntity, errors.BadEntity("query.url", "invalid url", "value_error.url.scheme"))
		return
	}
	// 从 SVG URL 获取数据
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("[SVG] failed to fetch SVG. reason: %s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch SVG"})
		return
	}
	defer resp.Body.Close()

	svgData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[SVG] failed to read SVG data. reason: %s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read SVG data"})
		return
	}

	// 转换 SVG 为 PNG
	pngData, err := svgToPng(svgData)
	if err != nil {
		fmt.Printf("[SVG] to png failed. reason: %s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to convert SVG to PNG"})
		return
	}
	// 返回 PNG 数据
	ctx.Header("Content-Type", "image/png")
	if download {
		ctx.Header("Content-Disposition", "attachment; filename=converted.png")
	}
	ctx.Writer.Write(pngData)
}

//	@id				convert.svg.png.get
//	@Summary		SVG to PNG
//	@Description	将 SVG 图片转为 png
//	@Tags			convert
//	@Produce		image/png
//	@Param			url			query	string	true	"svg图片地址"	example(https://www.docker.com/wp-content/uploads/2024/01/icon-docker-square.svg)
//	@Param			download	query	bool	false	"是否下载"		example(false)
//	@Router			/convert/svg/png [get]
//	@Success		200
func _(c *gin.Context) {}

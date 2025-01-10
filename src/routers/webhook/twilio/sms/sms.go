package sms

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddTwilioSmsRouter(router *gin.Engine) {
	group := router.Group("/api/webhook/twilio/sms")
	group.GET("/:sid/:token/:from/:to/:body", sms)
}

type SMSInputScheme struct {
	// Accoutn Sid
	Sid string `uri:"sid" form:"sid" binding:"required"`
	// Auth Token
	Token string `uri:"token" form:"token" binding:"required"`
	// 发件人
	From string `uri:"from" form:"from" binding:"required"`
	// 收件人
	To string `uri:"to" form:"to" binding:"required"`
	// 短信内容
	Body string `uri:"body" form:"body" binding:"required"`
}

func sms(ctx *gin.Context) {
	var params SMSInputScheme
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	form := url.Values{}
	form.Add("From", params.From)
	form.Add("To", params.To)
	form.Add("Body", params.Body)
	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", params.Sid)

	req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(params.Token))
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		ctx.JSON(http.StatusInternalServerError, "send sms failed.")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[DEBUG] [/api/webhook/twilio/sms] Error reading body:", err)
		} else {
			fmt.Printf("[DEBUG] [/api/webhook/twilio/sms] send sms failed, detail: %s\n", string(body))
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

// @id			twilio.sms.get
// @Summary		Twilio SMS
// @Description	通过 Twilio API 发送短信
// @Tags			Webhook
// @Produce		json
// @Param			sid		path	string	true	"sid"			Example(AC64f796e3a022cd)
// @Param			token	path	string	true	"auth token"	Example(32a8fc7ef68c1a6a79)
// @Param			from	path	string	true	"发送号码"			Example(+19711231234)
// @Param			to		path	string	true	"接收号码"			Example(+19711231234)
// @Param			body	path	string	true	"短信内容"			Example(helloworld)
// @Router			/webhook/twilio/sms/{sid}/{token}/{from}/{to}/{body} [get]
// @Success		200	{object}	errors.Success	地区代码
func _(c *gin.Context) {}

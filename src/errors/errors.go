package errors

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func BadEntity(loc, msg, _type string) gin.H {
	return gin.H{
		"detail": []gin.H{
			{
				"loc":  strings.Split(loc, "."),
				"msg":  msg,
				"type": _type,
			},
		},
	}
}

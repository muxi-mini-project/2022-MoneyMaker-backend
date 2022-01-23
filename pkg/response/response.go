package response

import "github.com/gin-gonic/gin"

type Rep struct {
	Msg  string
	Code int
}

func SendResponse(c *gin.Context, msg string, code int) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}

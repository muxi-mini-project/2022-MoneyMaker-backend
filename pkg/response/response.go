package response

import "github.com/gin-gonic/gin"

func SendResponse(c *gin.Context, msg string, code int) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}

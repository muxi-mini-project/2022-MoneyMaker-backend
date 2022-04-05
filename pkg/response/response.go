package response

import "github.com/gin-gonic/gin"

type Resp struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context, msg string, code int) {
	c.JSON(code, Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

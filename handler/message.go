package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "返回用户购买、收藏后的信息"
//@Description "信息已根据时间按升序排列，id越大越新，消息通知的api"
//@Tags Return Message
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"success","infor":[]tables.Message}
//@Failure 500 "error happened"
//@Router /money/message [get]
func Returnmsg(c *gin.Context) {
	var msgs []tables.Message
	id, exists := c.Get("id")
	if !exists {
		response.SendResponse(c, "error happened", 500)
		return
	}

	mysql.DB.Order("id desc").Where("my=?", id).Find(&msgs)

	c.JSON(200, gin.H{
		"msg":   "success",
		"infor": msgs,
	})
}

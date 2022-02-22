package handler

import (
	"log"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "返回用户购买、收藏后的信息"
//@Description "消息通知的api"
//@Tags Message
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"success","infor":[]tables.Message}
//@Failure 500 {string} json{"msg":"error happened","infor":[]tables.Message}
//@Router /money/message [get]
func Returnmsg(c *gin.Context) {
	var msgs []tables.Message
	id, exists := c.MustGet("id").(string)
	if !exists {
		response.SendResponse(c, "error happened", 500)
		log.Println(exists)
		return
	}

	mysql.DB.Order("id desc").Where("my=?", id).Find(&msgs)

	c.JSON(200, gin.H{
		"msg":   "success",
		"infor": msgs,
	})
}

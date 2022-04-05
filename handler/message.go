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
//@Success 200 {object} response.Resp "success"
//@Failure 500 {object} response.Resp "error happened in the server"
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

	c.JSON(200, response.Resp{
		Code: 200,
		Msg:  "login successfully",
		Data: msgs,
	})
}

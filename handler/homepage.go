package handler

import (
	"log"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "主页内容"
//@Description "主页的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param page query string true "页码"
//@Success 200 {object} response.Resp "success"
//@Success 500 {object} response.Resp "error happened in server"
//@Router /money/homepage [get]
func Homepage(c *gin.Context) {
	var goods []tables.Good

	page := c.DefaultQuery("page", "1")
	num := easy.STI(page)
	err := mysql.DB.Limit(10).Offset(num*10).Order("feed_back desc").Order("scores desc").Where("goodsin=?", "yes").Find(&goods).Error

	if err != nil || num == -1 {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err, num)
		return
	}

	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}

	c.JSON(200, response.Resp{
		Code: 200,
		Msg:  "successfully",
		Data: goods,
	})
}

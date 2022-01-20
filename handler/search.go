package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "搜索"
//@Description "搜索"
//@Tags Search
//@Accept application/json
//@Produce application/json
//@Param content body string true "content"
//@Success 200 "search successfully"
//@Success 204 "search nothing"
//@Router /money/search [post]
func Search(c *gin.Context) {
	var goods []tables.Good
	content := c.Request.FormValue("content")
	if err := mysql.DB.Select("summary", "avatar", "id", "goodszone").Where(`summary like "%?%"`, content).Find(&goods).Error; err != nil {
		response.SendResponse(c, "find nothing", 200)
		return
	}
	c.JSON(200, gin.H{
		"msg":   "success",
		"infor": goods,
	})
}

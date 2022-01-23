package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "主页内容"
//@Description "order=1->返回前十个商品的内容，summary不需要展示出来，是在商品详情页里，主页的api"
//@Tags Homepage
//@Accept application/json
//@Produce application/json
//@Param page query string true "page"
//@Success 200 {string} json{"msg":"success","infor":[]tables.Good}
//@Success 500 {string} json{"msg":"error"}
//@Router /money/homepage [post]
func Homepage(c *gin.Context) {
	var goods []tables.Good
	page := c.Query("page")
	num := easy.STI(page)
	err := mysql.DB.Order("feed_back desc").Order("scores desc").Find(&goods).Error

	if err != nil || num == -1 {
		response.SendResponse(c, "error", 500)
		return
	}

	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}
	if len(goods) < 10 {
		c.JSON(200, gin.H{
			"msg":   "success",
			"goods": goods,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":   "success",
			"goods": goods[:num*10],
		})
	}
}

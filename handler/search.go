package handler

import (
	"fmt"
	"log"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "搜索并返回已排序的商品信息"
//@Description "order=1->返回前十个商品的内容，summary不需要展示出来-》在商品详情页里,搜索的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param content formData string true "搜索框输入的内容"
//@Param page query string true "页码"
//@Success 200 {string} json{"msg":"search successfully","infor":[]tables.Good}
//@Success 204 {string} json{"msg":"find nothing"}
//@Router /money/search [post]
func Search(c *gin.Context) {
	//根据举报次数与评分进行优先返回
	var goods []tables.Good

	page := c.DefaultQuery("page", "1")
	num := easy.STI(page)
	content := c.Request.FormValue("content")
	err := mysql.DB.Limit(10).Offset(num * 10).Order("feed_back asc").Order("scores desc").Where(fmt.Sprintf(`summary like "%%%s%%" AND goodsin=%s`, content, "yes")).Find(&goods).Error

	if err != nil || num == -1 {
		response.SendResponse(c, "find nothing", 204)
		log.Println(err)
		return
	}
	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}
	//这里不需要返回图片的url

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

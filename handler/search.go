package handler

import (
	"fmt"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "搜索并返回已排序的商品信息"
//@Description "order=1->返回前十个商品的内容，summary不需要展示出来-》在商品详情页里,搜索的api"
//@Tags Search
//@Accept application/json
//@Produce application/json
//@Param content formData string true "content"
//@Param page query string true "page"
//@Success 200 {string} json{"msg":"search successfully","infor":[]tables.Good}
//@Success 204 {string} json{"msg":"find nothing"}
//@Router /money/search [post]
func Search(c *gin.Context) {
	//title
	//avatar
	//summary
	//zone
	//根据举报次数与评分进行优先返回
	var goods []tables.Good
	page := c.Query("page")
	num := easy.STI(page)
	content := c.Request.FormValue("content")
	if err := mysql.DB.Order("feed_back asc").Order("scores desc").Where(fmt.Sprintf(`summary like "%%%s%%"`, content)).Find(&goods).Error; err != nil {
		response.SendResponse(c, "find nothing", 204)
		return
	}
	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}
	//这里不需要返回图片的url

	//fmt.Println(goods)
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

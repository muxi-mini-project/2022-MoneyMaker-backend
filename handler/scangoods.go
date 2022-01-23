package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "查询某个商品的详细信息"
//@Description "点击进入商品详情页的api"
//@Tags Scan
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {string} json{"msg":"success","infor":tables.User,"user":tables.User,"tips":"如果goodsin是no则代表已经下架，此时则不显示开启交易按钮"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/goods/scanning [get]
func Scan(c *gin.Context) {
	var good tables.Good
	var user tables.User
	goodsidstr := c.Query("goodsid")
	goodsid := easy.STI(goodsidstr)
	err := mysql.DB.Where("goods_id=?", goodsid).Find(&good).Error
	if goodsid == -1 || err != nil {
		response.SendResponse(c, "error happened!", 500)
		return
	}
	good.Way = ""
	good.Buyer = ""
	mysql.DB.Model(&tables.User{}).Where("id=?", good.ID).Find(&user)
	user.Buygoods = ""
	c.JSON(200, gin.H{
		"msg":   "success",
		"infor": good,
		"user":  user,
	})

}

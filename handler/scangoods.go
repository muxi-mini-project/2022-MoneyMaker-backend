package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "查询某个商品的详细信息"
//@Description "点击进入商品详情页的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {string} json{"msg":"success","infor":tables.User,"user":tables.User,"tips":"如果goodsin是no则代表已经下架，此时则不显示开启交易按钮"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Failure 304 {string} json{"msg":"error in database"}
//@Router /money/goods/scanning [get]
func Scan(c *gin.Context) {
	var good tables.Good
	var user tables.User
	var msg = ""

	goodsidstr := c.Query("goodsid")
	goodsid := easy.STI(goodsidstr)
	//err := mysql.DB.Where("goods_id=?", goodsid).Find(&good).Error

	good, err := model.GetOrderGood(goodsid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}

	if goodsid == -1 {
		response.SendResponse(c, "error happened!", 500)
		return
	}

	good.Way = ""
	good.Buyer = ""
	//mysql.DB.Model(&tables.User{}).Where("id=?", good.ID).Find(&user)

	user, err = model.GetOrderUser(good.ID)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}
	if good.FeedBack >= 100 {
		msg = "该商品被举报次数过多,请谨慎交易!"
	}

	user.Buygoods = ""

	c.JSON(200, gin.H{
		"msg":   "success," + msg,
		"infor": good,
		"user":  user,
	})

}

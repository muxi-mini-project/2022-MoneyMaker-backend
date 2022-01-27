package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "用户进行购买"
//@Description "点击购买时的api"
//@Tags Buy
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {string} json{"msg":"success","way":"联系方式对应的url"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/goods/shopping [get]
func Buy(c *gin.Context) {
	//买完之后就展示联系方式，并把用户名字放在buyer中
	var (
		good   tables.Good
		strstr string
		user   tables.User
	)

	goodsidstring := c.Query("goodsid")
	stuid, exists := c.MustGet("id").(string)
	goodsid := easy.STI(goodsidstring)

	if !exists || goodsid == -1 {
		response.SendResponse(c, "error happened", 500)
		return
	}

	mysql.DB.Select("buyer", "way", "id").Where("goods_id=?", goodsid).Find(&good)

	//存储消息通知
	easy.Returnbuy(stuid, good.ID)

	if good.Buyer != "" {
		strstr = easy.New(good.Buyer, stuid)
	} else {
		strstr = strstr + stuid
	}

	mysql.DB.Model(&tables.Good{}).Where("goods_id=?", goodsid).Update("buyer", strstr)
	//mysql.DB.Where("goods_id=?", goodsid).Find(&good)
	//fmt.Println(good)

	//用户购买后把商品id存入buygoods
	mysql.DB.Model(&tables.User{}).Where("id=?", stuid).Find(&user)
	strstr = ""
	if user.Buygoods != "" {
		strstr = easy.New(user.Buygoods, goodsidstring)
	} else {
		strstr = strstr + goodsidstring
	}

	//re := easy.New(user.Buygoods, goodsidstring)
	mysql.DB.Model(&tables.User{}).Where("id=?", stuid).Update("buygoods", strstr)

	//Way存放图片对应的地址，再使用staticfs打开

	c.JSON(200, gin.H{
		"msg": "buy successfully",
		"way": good.Way,
	})
}

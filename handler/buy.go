package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "用户进行购买"
//@Description "点击购买时的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {object} response.Resp "success"
//@Failure 500 {object} response.Resp "error happened in server"
//@Failure 500 {object} response.Resp "error in database"
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
		response.SendResponse(c, "error happened in server", 500)
		log.Println("err")
		return
	}

	//mysql.DB.Select("buyer", "way", "id").Where("goods_id=?", goodsid).Find(&good)
	good, err := model.GetOrderGood(goodsid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}
	//存储消息通知
	easy.Returnbuy(stuid, good.ID)

	if good.Buyer != "" {
		strstr = easy.New(good.Buyer, stuid)
	} else {
		strstr = strstr + stuid
	}

	//mysql.DB.Model(&tables.Good{}).Where("goods_id=?", goodsid).Update("buyer", strstr)
	err = model.UpdateGoodBuyer(goodsid, strstr)
	if err != nil {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err)
		return
	}

	user, err = model.GetOrderUser(stuid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}
	strstr = ""
	if user.Buygoods != "" {
		strstr = easy.New(user.Buygoods, goodsidstring)
	} else {
		strstr = strstr + goodsidstring
	}

	//re := easy.New(user.Buygoods, goodsidstring)
	//mysql.DB.Model(&tables.User{}).Where("id=?", stuid).Update("buygoods", strstr)
	err = model.UpdateBuygoods(stuid, strstr)
	if err != nil {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err)
		return
	}

	//Way存放图片对应的地址，再使用staticfs打开

	c.JSON(200, response.Resp{
		Code: 200,
		Msg:  "buy successfully",
		Data: good.Way,
	})
}

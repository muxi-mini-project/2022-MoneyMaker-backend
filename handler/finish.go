package handler

import (
	"fmt"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type Two struct {
	Good  tables.Good
	Buyer []string
}

//@Summary "返回用户与卖家未完成的订单"
//@Description "返回订单,需要点完成的是‘my sell’->[]string 是与我做交易的人的id,因为一个商品可能被多个人购买，所以string切片的长度就是‘完成订单’的订单数,点评价的是‘my buy’"
//@Tags return trade
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"success","my buy":[]Two,"my sell":map[Two][]string}
//@Failure 500 {string} json{"error happened"}
//@Router /money/my/goods/unfish [get]
func UnFinish(c *gin.Context) {
	var goods []tables.Good
	var user tables.User
	var good tables.Good
	var strbuy []tables.Good
	var strsell []Two
	var two Two
	idstr, exists := c.Get("id")
	id, ok := idstr.(string)
	if !exists || !ok {
		//fmt.Println("1", exists, ok)
		response.SendResponse(c, "error", 500)
		return
	}
	//确认完成
	//获取我已购买我发布的商品的用户的id,可能不止一个商品,mysell
	mysql.DB.Model(&tables.Good{}).Select("price", "title", "goods_id", "buyer").Where("id=?", id).Find(&goods)
	for _, v := range goods {
		buyer := strings.Split(v.Buyer, ",")
		two.Buyer = buyer
		two.Good = v
		strsell = append(strsell, two)
		fmt.Println("1", buyer)
	}

	//评价
	mysql.DB.Where("id=?", id).Find(&user)
	hasbuy := strings.Split(user.Buygoods, ",")
	//fmt.Println(hasbuy)
	for _, v := range hasbuy {
		num := easy.STI(v)
		if num == -1 {
			//fmt.Println("2", num)
			response.SendResponse(c, "error", 500)
			return
		} else {
			mysql.DB.Model(&tables.Good{}).Select("price", "title", "goods_id").Where("goods_id=?", num).Find(&good)
			strbuy = append(strbuy, good)
		}
		fmt.Println("2")
	}
	c.JSON(200, gin.H{
		"msg":     "success",
		"my buy":  strbuy,
		"my sell": strsell,
	})
}

//@Summary "商家完成订单"
//@Description "点击确认完成时的api"
//@Tags Finish trade
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {string} json{"msg":"success"}
//@Failure 500 "error happened"
//@Router /money/my/goods/finish [get]
func Finsh(c *gin.Context) {
	var good tables.Good
	var user tables.User
	var re string
	//点击完成之后把这个购买者从buyer中删去，以及goodsid从uesr中删除,有多个的情况下，则只删除一个
	goodsid := c.Query("goodsid")
	strid, exists := c.Get("id")
	id, ok := strid.(string)
	if !exists || !ok {
		response.SendResponse(c, "error", 500)
		return
	}
	mysql.DB.Where("goods_id=?", goodsid).Find(&good)
	mysql.DB.Where("id=?", id).Find(&user)

	re = easy.Delete(good.Buyer, id)
	num := easy.STI(goodsid)
	if num == -1 {
		fmt.Println("2", num)
		response.SendResponse(c, "error", 500)
		return
	}
	mysql.DB.Model(&tables.Good{}).Where("goods_id=?", goodsid).Update("buyer", re)

	re = ""
	re = easy.Delete(user.Buygoods, goodsid)
	mysql.DB.Model(&tables.User{}).Where("id=?", id).Update("buygoods", re)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}

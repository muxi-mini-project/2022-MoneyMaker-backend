package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

//@Summary "查看我的橱窗"
//@Description "橱窗"
//@Tags MYgoods
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"check successfully","infot":[]tables.Good}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/my/goods [get]
func Mygoods(c *gin.Context) {
	var goods []tables.Good
	id, exists := c.Get("id")
	stuid, ok := id.(string)
	if !ok || !exists {
		response.SendResponse(c, "error happened!", 500)
		return
	}
	mysql.DB.Where("id=?", stuid).Find(&goods)
	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}
	c.JSON(200, gin.H{
		"msg":   "successfull",
		"infor": goods,
	})
}

//@Summary "查看我的购物车"
//@Description "购物车"
//@Tags MYCart
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"check successfully","infot":[]tables.Good}
//@Success 204 {string} json{"msg":"check successfully","infot":"nothing"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/my/cart [get]
func Mycart(c *gin.Context) {
	var (
		cart  tables.Cart
		goods []tables.Good
		good  tables.Good
	)

	id, exists := c.Get("id")
	stuid, ok := id.(string)

	if !ok || !exists {
		response.SendResponse(c, "error happened!", 500)
		return
	}

	mysql.DB.Where("id=?", stuid).Find(&cart)

	if cart.Goodsid == "" {
		response.SendResponse(c, "nothing", 204)
		return
	}

	goodsstrs := strings.Split(cart.Goodsid, ",")

	for _, v := range goodsstrs {
		goodsid := easy.STI(v)

		if goodsid != -1 {
			err := mysql.DB.Where("goods_id=?", goodsid).Find(&good).Error

			if err != nil {
				response.SendResponse(c, "error", 500)
				return
			}

			goods = append(goods, good)
		}
	}

	c.JSON(200, gin.H{
		"msg":   "success",
		"goods": goods,
	})
}

//@Summary "返回我的信息"
//@Description "我的个人信息的api"
//@Tags Message
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"avatar 是头像对应的url","infor":tables.User}
//@Failure 500 {string} json{"msg":"error happened","infor":tables.User}
//@Router /money/my/message [get]
func Mymessage(c *gin.Context) {
	var user tables.User

	id, exists := c.Get("id")
	if !exists {
		response.SendResponse(c, "return error!", 500)
		return
	}

	mysql.DB.Where("id=?", id).Find(&user)

	user.Buygoods = ""
	c.JSON(200, gin.H{
		"msg":   "success",
		"infor": user,
	})
}

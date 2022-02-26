package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

//@Summary "查看我的橱窗"
//@Description "橱窗"
//@Tags My
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"check successfully","infot":[]tables.Good}
//@Failure 500 {string} json{"msg":"error happened in server"}
//@Router /money/my/goods [get]
func Mygoods(c *gin.Context) {
	var goods []tables.Good
	id, exists := c.Get("id")
	stuid, ok := id.(string)
	if !ok || !exists {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(ok, exists)
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
//@Tags My
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"check successfully","infot":[]tables.Good}
//@Success 204 {string} json{"msg":"check successfully","infot":"nothing"}
//@Failure 500 {string} json{"msg":"error happened in server"}
//@Failure 304 {string} json{"msg":"error in database"}
//@Router /money/my/cart [get]
func Mycart(c *gin.Context) {
	var (
		cart  tables.Cart
		goods []tables.Good
		good  tables.Good
	)

	stuid, exists := c.MustGet("id").(string)

	if !exists {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(exists)
		return
	}

	//mysql.DB.Where("id=?", stuid).Find(&cart)
	cart, err := model.GetOrderCart(stuid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}

	if cart.Goodsid == "" {
		response.SendResponse(c, "nothing", 204)
		return
	}

	goodsStrs := strings.Split(cart.Goodsid, ",")

	for _, v := range goodsStrs {
		goodsid := easy.STI(v)

		if goodsid != -1 {
			model.GetOrderGood(goodsid)
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
//@Tags My
//@Accept application/json
//@Produce application/json
//@Success 200 {string} json{"msg":"avatar 是头像对应的url","infor":tables.User}
//@Failure 500 {string} json{"msg":"error happened in server","infor":tables.User}
//@Failure 304 {string} json{"msg":"error in database"}
//@Router /money/my/message [get]
func Mymessage(c *gin.Context) {
	var user tables.User

	id, exists := c.MustGet("id").(string)
	if !exists {
		response.SendResponse(c, "error happened in server!", 500)
		return
	}

	//mysql.DB.Where("id=?", id).Find(&user)
	user, err := model.GetOrderUser(id)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}

	user.Buygoods = ""
	c.JSON(200, gin.H{
		"msg":   "success",
		"infor": user,
	})
}

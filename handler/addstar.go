package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

//@Summary "新增收藏"
//@Description "add"
//@Tags Add
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 "add successfully"
//@Failure 500 "error happened"
//@Router /money/new_star [patch]
func Addstar(c *gin.Context) {
	//用户收藏后在cart里就会新增这个商品的goodsid
	var cart tables.Cart
	var re string
	id, exists := c.Get("id")
	goodsid := c.Query("goodsid")
	stuid, ok := id.(string)
	if !exists || !ok {
		response.SendResponse(c, "error happened", 500)
	}
	mysql.DB.Where("id=?", stuid).Find(&cart)
	user := strings.Split(cart.Goodsid, ",")
	user = append(user, goodsid)
	for i, v := range user {
		if i < len(user)-1 {
			re = re + v + ","
		}
		re = re + v
	}
	err := mysql.DB.Model(&tables.Cart{}).Where("id=?", stuid).Update("goodsid", re).Error
	if err != nil {
		response.SendResponse(c, "error happened", 500)
	}
	response.SendResponse(c, "add successfully!", 200)
}

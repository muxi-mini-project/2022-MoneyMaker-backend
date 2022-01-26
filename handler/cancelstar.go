package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "取消收藏"
//@Description "在购物车里取消收藏的api"
//@Tags Cancel
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {string} json{"msg":"cancel successfully"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/my/cancellation [post]
func Cancelstar(c *gin.Context) {
	var cart tables.Cart
	//var str string
	id, exists := c.Get("id")
	goodsid := c.Query("goodsid")
	stuid, ok := id.(string)

	//mysql.DB.Model(&tables.Cart{}).Where("id=?",stuid)

	mysql.DB.Where("id=?", stuid).Find(&cart)
	re := easy.Delete(cart.Goodsid, goodsid)
	err := mysql.DB.Model(&tables.Cart{}).Where("id=?", id).Update("goodsid", re).Error

	if !ok || !exists || err != nil {
		response.SendResponse(c, "error happened!", 500)
	}

	response.SendResponse(c, "cancel successfully", 200)
}

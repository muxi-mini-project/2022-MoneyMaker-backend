package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "取消收藏"
//@Description "在购物车里取消收藏的api"
//@Tags Star
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {string} json{"msg":"cancel successfully"}
//@Failure 500 {string} json{"msg":"error happened in server"}
//@Router /money/my/cancellation [post]
func Cancelstar(c *gin.Context) {
	var cart tables.Cart
	//var str string
	stuid, exists := c.MustGet("id").(string)
	goodsid := c.Query("goodsid")

	//mysql.DB.Where("id=?", stuid).Find(&cart)
	cart = model.GetOrderCart(stuid)
	re := easy.Delete(cart.Goodsid, goodsid)

	err := mysql.DB.Model(&tables.Cart{}).Where("id=?", stuid).Update("goodsid", re).Error
	if !exists || err != nil {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err, exists)
	}

	response.SendResponse(c, "cancel successfully", 200)
}

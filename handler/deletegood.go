package handler

import (
	"log"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "商家下架商品"
//@Description "下架商品的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {string} {"msg":"delete successfully"}
//@Failure 500 {string} {"msg":"error happened in server"}
//@Router /money/goods/deletion [delete]
func Deletegood(c *gin.Context) {
	goodsstr := c.Query("goodsid")
	goodsid := easy.STI(goodsstr)

	err := mysql.DB.Model(&tables.Good{}).Where("goods_id=?", goodsid).Update("goodsin", "no").Error

	if goodsid == -1 || err != nil {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err, goodsid)
		return
	}

	response.SendResponse(c, "delete successfully", 200)
}

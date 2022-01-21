package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "商家下架商品"
//@Description "购买"
//@Tags Delete
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 "delete successfully"
//@Failure 500 "error happened"
//@Router /money/goods/deletion [get]
func Deletegood(c *gin.Context) {
	goodsidstring := c.Query("goodsid")
	goodsid := easy.STI(goodsidstring)
	if goodsid != -1 {
		err := mysql.DB.Where("goodsid=?", goodsid).Delete(&tables.Good{}).Error
		if err != nil {
			response.SendResponse(c, "failed to delete!", 500)
		}
	}
	response.SendResponse(c, "delete successfully", 200)
}

package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//举报
//@Summary "接收举报"
//@Description "举报的api"
//@Tags Feedback
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {string} json{"msg":"举报成功!"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/goods/feedback [post]
func Feedback(c *gin.Context) {
	var good tables.Good
	goodsstr := c.Query("goodsid")
	goodsid := easy.STI(goodsstr)
	mysql.DB.Select("feed_back").Where("goods_id=?", goodsid).Find(&good)
	if goodsid == -1 {
		response.SendResponse(c, "error happened", 500)
		return
	}

	mysql.DB.Model(&tables.Good{}).Where("goods_id=?", goodsid).Update("feed_back", good.FeedBack+1)
	response.SendResponse(c, "举报成功", 200)
}

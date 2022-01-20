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
//@Description "举报次数增加"
//@Tags Feedback
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 "upload successfully"
//@Failure 500 "error happened"
//@Router /money/goods/feedback [post]
func Feedback(c *gin.Context) {
	var good tables.Good
	goodsstr := c.Query("goodsid")
	goodsid := easy.STI(goodsstr)
	mysql.DB.Select("feedback").Where("goodsid=?", goodsid).Find(&good)
	if goodsid != -1 {
		mysql.DB.Model(&tables.Good{}).Where("goodsid=?", goodsid).Update("feedback", good.FeedBack+1)
		response.SendResponse(c, "举报成功", 200)
		return
	}
	response.SendResponse(c, "error happened", 500)
}

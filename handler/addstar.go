package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "新增收藏"
//@Description "添加至购物车时的api"
//@Tags Add
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {string} json{"msg":"add successfully" "msg":"你已经收藏过该商品了"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/new_star [patch]
func Addstar(c *gin.Context) {
	//用户收藏后在cart里就会新增这个商品的goodsid
	var (
		cart tables.Cart
		re   string
		good tables.Good
	)

	id, exists := c.Get("id")
	goodsid := c.Query("goodsid")
	stuid, ok := id.(string)

	if !exists || !ok {
		response.SendResponse(c, "error happened", 500)
	}

	mysql.DB.Where("id=?", stuid).Find(&cart)
	if cart.Goodsid != "" {
		re, ok = easy.NewSingle(cart.Goodsid, goodsid)
	} else {
		re = re + goodsid
	}

	if ok {
		err := mysql.DB.Model(&tables.Cart{}).Where("id=?", stuid).Update("goodsid", re).Error
		if err != nil {
			response.SendResponse(c, "error happened", 500)
			return
		}
	}

	goodsidint := easy.STI(goodsid)
	if goodsidint == -1 {
		response.SendResponse(c, "error happened", 500)
		return
	}

	mysql.DB.Where("goods_id=?", goodsidint).Find(&good)

	//保存信息
	easy.Returnstar(stuid, good.ID)

	if ok {
		response.SendResponse(c, "add successfully!", 200)
	} else {
		response.SendResponse(c, "你已经收藏过该商品哦!", 200)
	}

}

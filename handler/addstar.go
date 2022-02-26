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

//@Summary "新增收藏"
//@Description "添加至购物车时的api"
//@Tags Star
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品的编号"
//@Success 200 {string} json{"msg":"add successfully" "msg":"你已经收藏过该商品了"}
//@Failure 500 {string} json{"msg":"error happened in server"}
//@Failure 304 {string} json{"msg":"error in database"}
//@Router /money/new_star [patch]
func Addstar(c *gin.Context) {
	//用户收藏后在cart里就会新增这个商品的goodsid
	var (
		cart tables.Cart
		re   string
		good tables.Good
		ok   bool
	)

	stuid, exists := c.MustGet("id").(string)
	goodsid := c.Query("goodsid")

	if !exists {
		response.SendResponse(c, "error happened", 500)
	}

	//mysql.DB.Where("id=?", stuid).Find(&cart)
	cart, err := model.GetOrderCart(stuid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}
	if cart.Goodsid != "" {
		re, ok = easy.NewSingle(cart.Goodsid, goodsid)
	} else {
		re = re + goodsid
	}

	if ok {
		err := mysql.DB.Model(&tables.Cart{}).Where("id=?", stuid).Update("goodsid", re).Error
		if err != nil {
			response.SendResponse(c, "error happened in server", 500)
			log.Println(err)
			return
		}
	}

	goodsidInt := easy.STI(goodsid)
	if goodsidInt == -1 {
		response.SendResponse(c, "error happened in server", 500)
		log.Println("STI错误")
		return
	}

	//mysql.DB.Where("goods_id=?", goodsidint).Find(&good)
	good, err = model.GetOrderGood(goodsidInt)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}

	//保存信息
	easy.Returnstar(stuid, good.ID)

	if ok {
		response.SendResponse(c, "add successfully!", 200)
	} else {
		response.SendResponse(c, "你已经收藏过该商品哦!", 200)
	}

}

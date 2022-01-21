package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var Goodsway string
var Goodsurl string
var Goodsidstring string

//@Summary "用户进行购买"
//@Description "购买"
//@Tags Buy
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {object} {"msg":"success","img":"联系方式对应的url"}
//@Failure 500 "error happened"
//@Router /money/goods/shopping [get]
func Buy(c *gin.Context) {
	//买完之后就展示联系方式，并把用户名字放在buyer中
	var good tables.Good
	var strstr string
	goodsidstring := c.Query("goodsid")
	id, exists := c.Get("id")
	stuid, ok := id.(string)
	goodsid := easy.STI(goodsidstring)
	if !exists || !ok || goodsid == -1 {
		response.SendResponse(c, "error happened", 500)
		return
	}

	mysql.DB.Select("buyer", "way").Where("goodsid=?", goodsid).Find(&good)

	str := strings.Split(good.Buyer, ",")
	str = append(str, stuid)

	for i, v := range str {
		if i < len(str)-1 {
			strstr = strstr + v + ","
		}
		strstr = strstr + v
	}
	mysql.DB.Model(&tables.Good{}).Where("goodsid=?", stuid).Update("buyer", strstr)

	//Way存放图片对应的地址，再使用staticfile打开
	Goodsway = good.Way
	Goodsurl = "localhost:8080/goods/way/" + strconv.Itoa(goodsid)
	c.JSON(200, gin.H{
		"msg": "buy successfully",
		"img": Goodsurl,
	})
}

package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"
	"miniproject/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Tmp struct {
	Title   string
	Zone    string
	Summary string
	Price   int
}

//@Summary "上架商品"
//@Description "新增一个商品"
//@Tags New goods
//@Accept application/json
//@Produce application/json
//@Param avatar formdata file true "商品图二进制文件"
//@Param way formdata file true "联系方式二进制文件"
//@Success 200 "upload successfully"
//@Failure 500 "error happened"
//@Router /money/goods/addition [post]
func Addgood(c *gin.Context) {
	//新增一个商品,goodsid不需要去获取，设置了自增就可以，只要管其他的字段
	var good1 tables.Good
	var good2 tables.Good
	var tmp Tmp
	id, exists := c.Get("id")
	stuid, ok := id.(string)
	if !exists || !ok {
		response.SendResponse(c, "error happened", 500)
	}
	sql := `SELECT max(goodsid) FROM goods`
	if err := mysql.DB.Exec(sql).Find(&good1).Error; err != nil {
		Init(c)
	}
	good2.ID = stuid
	good2.GoodsID = good1.GoodsID
	wayA := upload.UploadA(c, good2.GoodsID)
	good2.Avatar = wayA

	if err := c.ShouldBindJSON(&tmp); err != nil {
		response.SendResponse(c, "error happened", 500)
	}

	wayW := upload.UploadW(c, good2.GoodsID)
	good2.Way = wayW

	good2.Price = tmp.Price
	good2.Goodszone = tmp.Zone
	good2.Summary = tmp.Summary
	good2.Title = tmp.Title

	mysql.DB.Create(&good2)
	response.SendResponse(c, "upload successfully!", 200)

}

//防止第一个新增商品出现错误
func Init(c *gin.Context) {
	var good tables.Good
	var user tables.User
	good.GoodsID = 0
	good.ID = "0"
	user.ID = "0"
	err1 := mysql.DB.Select("id").Create(&user).Error
	err2 := mysql.DB.Select("id", "goodsid").Create(&good).Error
	if err1 != nil || err2 != nil {
		response.SendResponse(c, "error happened!", 500)
	}
}

package handler

import (
	"fmt"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"
	"miniproject/pkg/upload"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tmp struct {
	Title   string `json:"title" binding:"required"`
	Zone    string `json:"zone" binding:"required"`
	Summary string `json:"summary"`
	Price   int    `json:"price" binding:"required"`
}

//@Summary "上架商品"
//@Description "新增一个商品时的api"
//@Tags New goods
//@Accept multipart/form-data
//@Produce application/json
//@Param avatar formData file true "商品图二进制文件"
//@Param way formData file true "联系方式二进制文件"
//@Success 200 {string} json{"msg":"upload successfully"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/goods/addition [post]
func Addgood(c *gin.Context) {
	//新增一个商品,goodsid不需要去获取，设置了自增就可以，只要管其他的字段
	var good1 tables.Good
	var good2 tables.Good
	var tmp Tmp
	var msga string
	var msgw string
	id, exists := c.Get("id")
	stuid, ok := id.(string)
	if !exists || !ok {
		response.SendResponse(c, "error happened", 500)
		return
	}
	//sql := `SELECT max(goodsid) FROM goods`
	err1 := mysql.DB.Model(&tables.Good{}).Last(&good1).Error
	err2 := c.ShouldBindJSON(&tmp)
	if err1 != nil || err2 != nil {
		response.SendResponse(c, "error happened", 500)
		return

	}
	fmt.Println(good1)
	good2.ID = stuid
	//获取到当前的最大值之后再加一即可
	good2.GoodsID = good1.GoodsID + 1
	//存放图片到本地
	oka := upload.UploadA(c, good2.GoodsID)
	if !oka {
		msga = "the avatar failed to upload!"
	}
	//good2.Avatar = wayA

	//存放图片到本地
	okw := upload.UploadW(c, good2.GoodsID)
	if !okw {
		msgw = "the way failed to upload!"
	}
	//good2.Way = wayW

	good2.Price = tmp.Price
	good2.Goodszone = tmp.Zone
	if tmp.Summary == "" {
		good2.Summary = "该商品暂无其他描述"
	} else {
		good2.Summary = tmp.Summary
	}

	good2.Title = tmp.Title
	//直接存url
	good2.Avatar = "localhost:8080/images/avatar/" + strconv.Itoa(good2.GoodsID) + ".jpg"
	good2.Way = "localhost:8080/images/way/" + strconv.Itoa(good2.GoodsID) + ".jpg"

	mysql.DB.Model(&tables.Good{}).Create(&good2)

	response.SendResponse(c, "upload successfully!"+msga+","+msgw, 200)
}

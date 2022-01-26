package handler

import (
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
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
//@Param title formData string true "标题"
//@Param zone formData string true "分区"
//@Param price formData integer true "价格"
//@Param summary formData string true "详情"
//@Param avatar formData file true "商品图二进制文件"
//@Param way formData file true "联系方式二进制文件"
//@Success 200 {string} json{"msg":"upload successfully"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/goods/addition [post]
func Addgood(c *gin.Context) {
	//新增一个商品,goodsid不需要去获取，设置了自增就可以，只要管其他的字段
	var (
		good1 tables.Good
		good2 tables.Good
		msga  string
		msgw  string
	)

	id, exists := c.Get("id")
	stuid, ok := id.(string)
	if !exists || !ok {
		response.SendResponse(c, "error happened", 500)
		return
	}

	err1 := mysql.DB.Model(&tables.Good{}).Last(&good1).Error
	summary := c.PostForm("title")
	zone := c.PostForm("zone")
	price := c.PostForm("price")
	title := c.PostForm("title")

	tmp := easy.STI(price)
	if tmp == -1 || err1 != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}

	good2.ID = stuid
	//获取到当前的最大值之后再加一即可
	good2.GoodsID = good1.GoodsID + 1

	//存放图片到本地
	oka := upload.UploadAvatar(c, good2.GoodsID)
	if !oka {
		msga = "the avatar failed to upload!"
	}

	//存放图片到本地
	okw := upload.UploadWay(c, good2.GoodsID)
	if !okw {
		msgw = "the way failed to upload!"
	}
	//good2.Way = wayW

	good2.Price = tmp
	good2.Goodszone = zone
	if summary == "" {
		good2.Summary = "该商品暂无其他描述"
	} else {
		good2.Summary = summary
	}

	good2.Title = title

	//直接存url
	good2.Avatar = "localhost:8080/images/avatar/" + strconv.Itoa(good2.GoodsID) + ".jpg"
	good2.Way = "localhost:8080/images/way/" + strconv.Itoa(good2.GoodsID) + ".jpg"
	good2.Goodsin = "yes"

	mysql.DB.Model(&tables.Good{}).Create(&good2)

	response.SendResponse(c, "upload successfully!"+msga+","+msgw, 200)
}

package handler

import (
	"miniproject/model"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type All struct {
	One     int
	Two     int
	Three   int
	Four    int
	Five    int
	Sum     int
	Average float64
	Person  int
}

type Description struct {
	Score   int    `json:"score" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}

//@Summary "获取某个商品的所有评论"
//@Description "商品详情页点击评价时的api "scores":所有分值情况, "infor":"评论信息以及学号""
//@Tags Comment
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {string} {"infor":[]tables.Comment,"score":All}
//@Failure 500 {string} json{"msg":"err"}
//@Router /money/goods/comments [get]
func Getcomment(c *gin.Context) {
	//先获取goodsid
	var re []tables.Comment
	var all All

	goodsidstring := c.Query("goodsid")
	goodsid, err := strconv.Atoi(goodsidstring)
	if err != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}

	//mysql.DB.Model(&tables.Comment{}).Where("goods_id=?", goodsid).Find(&re)
	re = model.GetGoodComment(goodsid)

	for i := 0; i < len(re); i++ {
		all.Sum += re[i].Score
		switch re[i].Score {
		case 1:
			all.One++
		case 2:
			all.Two++
		case 3:
			all.Three++
		case 4:
			all.Four++
		case 5:
			all.Five++
		}
	}
	all.Average = float64(all.Sum) / float64(len(re))
	all.Person = len(re)
	//fmt.Println(all)
	//若果要返回一个自定义结构体，那么它的字段应该要大写，否则会无法识别
	c.JSON(200, gin.H{
		"infor":  re,
		"scores": all,
		"msg":    "success",
	})
}

//@Summary "用户对某个商品的评论"
//@Description "用户做出评价，点击评价时的api"
//@Tags Givecomment
//@Accept application/json
//@Produce application/json
//@Param comment body Description true "评论"
//@Param goodsid query string true "商品编号"
//@Success 200 {string} json{"msg":"give successfully"}
//@Failure 500 {string} json{"msg":"error happened"}
//@Router /money/goods/comment [post]
func Givecomment(c *gin.Context) {
	var (
		re  []tables.Comment
		des Description
		cmt tables.Comment
	)

	goodsid := c.Query("goodsid")
	userid, exists := c.MustGet("id").(string)

	if err := c.ShouldBindJSON(&des); err != nil || !exists {
		response.SendResponse(c, "error happened", 500)
		return
	}

	//获取当前时间
	tm := time.Now().Format("2006-01-02 15:04:05")

	cmt.ID = userid
	cmt.GoodsID = easy.STI(goodsid)
	cmt.Givetime = tm
	cmt.Score = des.Score
	cmt.Comment = des.Comment
	err := mysql.DB.Create(&cmt).Error

	if cmt.GoodsID == -1 || err != nil {
		//fmt.Println("2", cmt.GoodsID, err, ok)
		response.SendResponse(c, "error happened", 500)
		return
	}

	//更新商品的平均分,得在创建一条新评论之后
	mysql.DB.Select("score").Where("goods_id=?", goodsid).Find(&re)

	if str := Average(c, re, goodsid); str != "" {
		//fmt.Println("3", str)
		response.SendResponse(c, "error happened", 500)
		return
	}

	response.SendResponse(c, "give successfully", 200)
}

func Average(c *gin.Context, re []tables.Comment, goodsid string) string {
	var good tables.Good
	var sum = 0
	id := easy.STI(goodsid)
	if id == -1 {
		return "error"
	}
	if len(re) == 0 {
		good.Scores = 0
	}
	for _, v := range re {
		sum += v.Score
	}
	good.Scores = float64(sum) / float64(len(re))

	mysql.DB.Model(&tables.Good{}).Where("goods_id=?", id).Update("scores", good.Scores)
	return ""
}

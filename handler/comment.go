package handler

import (
	"fmt"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type All struct {
	one     int
	two     int
	three   int
	four    int
	five    int
	sum     int
	average int
}

type Description struct {
	Score   int    `gorm:"score" json:"score"`
	Comment string `gorm:"comment" json:"comment"`
}

//@Summary "获取某个商品的所有评论"
//@Description "分值已处理,评论无序"
//@Tags Comment
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 {object} []tables.Comment"{"all":所有分值情况, "infor":"评论信息以及学号"}"
//@Failure 500 "err"
//@Router /money/goods/comments [get]
func Getcomment(c *gin.Context) {
	//先获取goodsid
	var re []tables.Comment
	//var ans = make(map[string]string)
	var all All
	goodsidstring := c.Query("goodsid")
	goodsid, err := strconv.Atoi(goodsidstring)
	if err != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}
	mysql.DB.Select("comment", "id", "score", "commentid").Where("goodsid=?", goodsid).Find(&re)
	for i := 0; i < len(re); i++ {
		all.sum += re[i].Score
		switch re[i].Score {
		case 1:
			all.one++
		case 2:
			all.two++
		case 3:
			all.three++
		case 4:
			all.four++
		case 5:
			all.five++
		}
	}
	all.average = all.sum / len(re)
	c.JSON(200, gin.H{
		"scores": all,
		"infor":  re,
	})
}

//@Summary "用户对某个商品的评论"
//@Description "用户评价"
//@Tags Givecomment
//@Accept application/json
//@Produce application/json
//@Param comment body Description true "comment"
//@Param goodsid query string true "goodsid"
//@Success 200 "upload successfullu"
//@Failure 500 "error happened"
//@Router /money/goods/comment [post]
func Givecomment(c *gin.Context) {
	var des Description
	var cmt tables.Comment
	goodsid := c.Query("goodsid")
	fmt.Println(goodsid)
	id, exists := c.Get("id")
	if err := c.ShouldBindJSON(&des); err != nil && !exists {
		response.SendResponse(c, "error happened", 500)
		return
	}

	userid, ok := id.(string)
	cmt.ID = userid
	cmt.GoodsID = easy.STI(goodsid)
	err := mysql.DB.Select("id", "comment", "score", "goodsid").Create(&cmt).Error
	if cmt.GoodsID == -1 || !ok || err != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}
	response.SendResponse(c, "success", 200)
}

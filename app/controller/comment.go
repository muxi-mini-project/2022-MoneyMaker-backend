package controller

import (
	"fmt"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"strconv"

	"github.com/gin-gonic/gin"
)

//@Summary "获取某个商品的评论"
//@Description "分值已处理,评论无序"
//@Id get-comment
//@Tags Comment
//@Accept application/json
//@Produce application/json
//@Param goodsid path string true "goodsid"
//@Success 200 {object} []tables.Comment"{"all":所有分值情况, "infor":"评论信息以及学号"}"
//@Failure 500 "err"
//@Router /money/goods/comment/{goodsid} [get]
type All struct {
	one     int
	two     int
	three   int
	four    int
	five    int
	sum     int
	average int
}

func Comment(c *gin.Context) {
	//先获取goodsid
	var re []tables.Comment
	var all All
	goodsidstring := c.Param("goodsid")
	goodsid, err := strconv.Atoi(goodsidstring)
	if err != nil {
		c.JSON(500, gin.H{
			"err": "err",
		})
		fmt.Println("转化错误")
		return
	}
	mysql.DB.Select("comment", "id", "score").Where("goodsid=?", goodsid).Find(&re)
	for i := 0; i < len(re); i++ {
		all.sum += re[i].Score
		if re[i].Score == 1 {
			all.one++
		}
		if re[i].Score == 2 {
			all.two++
		}
		if re[i].Score == 3 {
			all.three++
		}
		if re[i].Score == 4 {
			all.four++
		}
		if re[i].Score == 5 {
			all.five++
		}
	}
	all.average = all.sum / len(re)
	c.JSON(200, gin.H{
		"scores": all,
		"infor":  re,
	})

}

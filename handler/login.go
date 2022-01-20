package handler

import (
	model "miniproject/model/getstu"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"
	"miniproject/pkg/token"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id" binding:"required"` //一定要输入的加上了required
	Password string `json:"password" binding:"required"`
}

//@Summary "用户登录"
//@Description "进行核对"
//@Tags Givecomment
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "goodsid"
//@Success 200 "login successfully"
//@Failure 500 "error happened"
//@Router /money/goods/comment [post]
func Login(c *gin.Context) {
	var user user
	if err := c.ShouldBindJSON(&user); err != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}

	_, err := model.GetUserInfoFormOne(user.ID, user.Password)

	if err != nil {
		//用户认证信息错误返回401状态码
		c.JSON(401, gin.H{
			"error": "密码或学号错误",
		})
		return
	}
	Create(user.ID)
	token, err := token.GenerateToken(user.ID)
	if err != nil {
		response.SendResponse(c, "token生成错误", 500)
	}
	c.JSON(200, gin.H{
		"msg":   "登录成功,请保留token并将其放在之后的请求头中",
		"token": token,
	})
}

func Create(id string) {
	var user tables.User
	if err := mysql.DB.Where("id=?", id).Find(&user).Error; err != nil {
		mysql.DB.Model(&tables.User{}).Create(map[string]interface{}{
			"id": id,
		})
	}
}

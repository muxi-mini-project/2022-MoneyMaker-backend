package controller

import (
	"fmt"
	model "miniproject/model/getstu"
	"miniproject/utils/token"

	"github.com/gin-gonic/gin"
)

type Json struct {
	ID       string `json:"id" binding:"required"` //一定要输入的加上了required
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user Json
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("bindjson错误")
	}

	_, err := model.GetUserInfoFormOne(user.ID, user.Password)

	if err == nil {
		token, err := token.GenerateToken(user.ID)
		if err != nil {
			//log.Fatal(err)
			fmt.Println("token返回错误")
		}
		c.JSON(200, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
	} else {
		//用户认证信息错误返回401状态码
		c.JSON(401, gin.H{
			"error": "密码或学号错误",
		})
	}
}

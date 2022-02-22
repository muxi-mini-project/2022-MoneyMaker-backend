package handler

import (
	"encoding/base64"
	"log"
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
//@Description "登录的api"
//@Tags Login
//@Accept application/json
//@Produce application/json
//@Param user body user true "id 学号 password 密码进行base64加密后的字符串"
//@Success 200 {string} json{""msg":   "登录成功","token": token,"tips": "请保留token并将其放在之后的请求头中"}
//@Failure 401 {string} json{"msg":"unauthorization"}
//@Failure 500 {string} json{"msg":"token生成错误"}
//@Router /money/entrance [post]
func Login(c *gin.Context) {
	var user user
	var acc tables.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}
	//第一次登陆进if
	if err := mysql.DB.Model(&tables.User{}).Where("id=?", user.ID).Find(&acc).Error; err != nil {
		psd, err := base64.StdEncoding.DecodeString(user.Password)
		if err != nil {
			response.SendResponse(c, "fail to decode password", 500)
			log.Println(err)
			return
		}

		stu, err := model.GetUserInfoFormOne(user.ID, string(psd))

		if err != nil {
			//用户认证信息错误返回401状态码
			c.JSON(401, gin.H{
				"error": "密码或学号错误",
			})
			return
		}

		token, err := token.GenerateToken(user.ID)
		if err != nil {
			response.SendResponse(c, "token生成错误", 500)
			log.Println(err)
			return
		}

		mysql.Create(user.ID, stu.User.Name, string(psd))

		c.JSON(200, gin.H{
			"msg":   "登录成功",
			"token": token,
			"tips":  "请保留token并将其放在之后的请求头中",
		})
	} else {
		psd, err := base64.StdEncoding.DecodeString(user.Password)
		if err != nil {
			response.SendResponse(c, "fail to decode password", 500)
			log.Println(err)
			return
		}

		if acc.Password == string(psd) {
			token, err := token.GenerateToken(user.ID)
			if err != nil {
				response.SendResponse(c, "token生成错误", 500)
				log.Println(err)
				return
			}

			c.JSON(200, gin.H{
				"msg":   "登录成功",
				"token": token,
				"tips":  "请保留token并将其放在之后的请求头中",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "密码或学号错误",
			})
			return
		}
	}

}

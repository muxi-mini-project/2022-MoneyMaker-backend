package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/getstu"
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
//@Param user body user true "id 学号 password 密码"
//@Success 200 {object} response.Resp "登录成功"
//@Failure 401 {object} response.Resp "unauthorized"
//@Failure 500 {object} response.Resp "token生成错误"
//@Router /entrance [post]
func Login(c *gin.Context) {
	var user user
	var acc tables.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}
	//第一次登陆进if
	acc = model.GetLoginInfor(user.ID)
	if acc.ID != user.ID {

		stu, err := getstu.GetUserInfoFormOne(user.ID, user.Password)

		if err != nil {
			//用户认证信息错误返回401状态码
			c.JSON(200, response.Resp{
				Code: 401,
				Msg:  "密码或学号错误",
				Data: nil,
			})
			return
		}

		token, err := token.GenerateToken(user.ID)
		if err != nil {
			response.SendResponse(c, "token生成错误", 500)
			log.Println(err)
			return
		}

		mysql.Create(user.ID, stu.User.Name, user.Password)

		c.JSON(200, response.Resp{
			Code: 200,
			Msg:  "login successfully",
			Data: token,
		})
	} else {

		if acc.Password == user.Password {
			token, err := token.GenerateToken(user.ID)
			if err != nil {
				response.SendResponse(c, "token生成错误", 500)
				log.Println(err)
				return
			}

			c.JSON(200, response.Resp{
				Code: 200,
				Msg:  "login successfully",
				Data: token,
			})
		} else {
			c.JSON(200, response.Resp{
				Code: 401,
				Msg:  "密码或学号错误",
				Data: nil,
			})
			return
		}
	}

}

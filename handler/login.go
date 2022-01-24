package handler

import (
	model "miniproject/model/getstu"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/avatar"
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
//@Tags Givecomment
//@Accept application/json
//@Produce application/json
//@Param user body user true "用户信息"
//@Success 200 {string} json{"msg":"login successfully"}
//@Failure 401 {string} json{"msg":"error happened"}
//@Router /money/entrance [post]
func Login(c *gin.Context) {
	var user user
	if err := c.ShouldBindJSON(&user); err != nil {
		response.SendResponse(c, "error happened", 500)
		return
	}

	stu, err := model.GetUserInfoFormOne(user.ID, user.Password)

	if err != nil {
		//用户认证信息错误返回401状态码
		c.JSON(401, gin.H{
			"error": "密码或学号错误",
			"msg":   "如果账号密码无误请多次尝试登录",
		})
		return
	}

	Create(user.ID, stu.User.Name)

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		response.SendResponse(c, "token生成错误", 500)
	}

	c.JSON(200, gin.H{
		"msg":   "登录成功,请保留token并将其放在之后的请求头中",
		"token": token,
	})
}

//登录成功初始化这个用户的信息
func Create(id string, name string) {
	var user tables.User
	//不存在就会报错
	mysql.DB.Where("id=?", id).Find(&user)
	avatar := avatar.GetAvatar()
	if user.ID != id {
		mysql.DB.Model(&tables.User{}).Create(map[string]interface{}{
			"id":       id,
			"avatar":   avatar, //随机分配头像
			"nickname": name,
			"buygoods": "",
		})

		//新建一个空的购物车
		mysql.DB.Model(&tables.Cart{}).Create(map[string]interface{}{
			"id": id})

	}

}

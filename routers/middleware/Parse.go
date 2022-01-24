package middleware

import (
	"miniproject/config"
	"miniproject/pkg/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Parse(c *gin.Context) {
	//dvar stu controller.Json
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "请求头中的auth为空",
		})
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "请求头中的auth格式有误",
		})
		c.Abort()
		return
	}

	token, err := token.ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "token无效",
		})
		c.Abort()
		return
	}

	id := token.ID
	c.Set("id", id)
	issuer := token.Issuer
	//_, err := model.GetUserInfoFormOne()

	if issuer != config.Issuer {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "发布者错误",
		})
		c.Abort()
		return
	}
}

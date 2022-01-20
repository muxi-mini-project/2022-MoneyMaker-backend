package main

import (
	"miniproject/model/mysql"
	router "miniproject/routers"

	"github.com/gin-gonic/gin"
)

//@title miniproject
//@version 1.0.0
//@description 赚圈圈API
//@termsOfService http://swagger.io/terrms/

//@contact.name yyj
//@contact.email 2105753640@qq.com

//@host localhost:8080
//@BasePath /api/vi
//@Schemes http

func main() {
	mysql.DataInit()
	gin.SetMode(gin.ReleaseMode)
	router := router.Router()
	router.Run(":8080")
}

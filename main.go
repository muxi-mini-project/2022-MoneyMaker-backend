package main

import (
	"miniproject/model/mysql"
	router "miniproject/routers"

	"github.com/gin-gonic/gin"
)

//@title miniproject
//@version 1.0.0
//@description "赚圈圈API 返回的goods如果其中的goodsin为yes代表它可以进行交易，即：未下架"
//@termsOfService http://swagger.io/terrms/

//@contact.name yyj
//@contact.email 2105753640@qq.com

//@host localhost:8080
//@BasePath /api/vi
//@Schemes http

func main() {
	mysql.Init()
	gin.SetMode(gin.ReleaseMode)
	router := router.Router()
	router.Run(":8080")
}

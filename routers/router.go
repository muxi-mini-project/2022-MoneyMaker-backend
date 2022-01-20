package router

import (
	"miniproject/app/controller"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	router := gin.Default()

	//router.POST("/api/vi/money/entrance", controller.Login)
	router.GET("/api/vi/money/goods/comment/:goodsid", controller.Comment)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

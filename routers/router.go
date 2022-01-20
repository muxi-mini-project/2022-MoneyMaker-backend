package router

import (
	"miniproject/handler"
	"miniproject/routers/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.POST("/api/vi/entrance", handler.Login)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group := router.Group("/api/vi/money")
	group.Use(middleware.Parse)
	{
		groupcomment := group.Group("/goods")
		{
			groupcomment.GET("/comments", handler.Getcomment)
			groupcomment.POST("/comment", handler.Givecomment)
			groupcomment.PUT("/feedback", handler.Feedback)
		}

		groupmy := group.Group("/my")
		{
			groupmy.GET("/cart", handler.Mycart)
			groupmy.GET("/goods", handler.Mygoods)
		}
		group.POST("/search", handler.Search)
	}
	return router
}

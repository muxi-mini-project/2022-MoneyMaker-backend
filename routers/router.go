package router

import (
	"miniproject/handler"
	"miniproject/routers/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/vi/entrance", handler.Login) //ok
	group := router.Group("/api/vi/money")
	group.Use(middleware.Parse) //ok
	{
		groupgoods := group.Group("/goods")
		{
			groupgoods.GET("/comments", handler.Getcomment)    //ok
			groupgoods.POST("/comment", handler.Givecomment)   //ok
			groupgoods.PUT("/feedback", handler.Feedback)      //ok
			groupgoods.DELETE("/deletion", handler.Deletegood) //ok
			groupgoods.POST("/addition", handler.Addgood)      //ok
			groupgoods.POST("/shopping", handler.Buy)          //ok
			groupgoods.GET("/scanning", handler.Scan)          //ok
		}

		groupmy := group.Group("/my")
		{
			groupmy.GET("/cart", handler.Mycart)
			groupmy.GET("/goods", handler.Mygoods)            //ok
			groupmy.PATCH("/new_star", handler.Addstar)       //ok
			groupmy.POST("/cancellation", handler.Cancelstar) //ok
			groupmy.GET("/message", handler.Mymessage)        //ok
			groupmy.GET("/goods/unfinish", handler.UnFinish)
			groupmy.POST("/goods/finish", handler.Finsh)
		}
		group.POST("/search", handler.Search)    //ok
		group.GET("/homepage", handler.Homepage) //ok
		group.GET("/message", handler.Returnmsg) //ok
	}
	//购买后返回联系方式的接口
	router.StaticFS("/images/way", http.Dir("C:\\source\\go\\miniproject\\goods\\way"))
	//查看商品时，返回的图片
	router.StaticFS("/images/avatar", http.Dir("C:\\source\\go\\miniproject\\goods\\avatar"))
	return router
}

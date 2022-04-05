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
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	{
		groupgoods := group.Group("/goods")
		{
			//获取商品的评论
			groupgoods.GET("/comments", handler.Getcomment) //ok

			//发表评论
			groupgoods.POST("/comment", handler.Givecomment) //ok

			//进行举报
			groupgoods.PUT("/feedback", handler.Feedback) //ok

			//下架商品
			groupgoods.DELETE("/deletion", handler.Deletegood) //ok

			//新增商品
			groupgoods.POST("/addition", handler.Addgood) //ok

			//购买商品
			groupgoods.POST("/shopping", handler.Buy) //ok

			//商品详情页
			groupgoods.GET("/scanning", handler.Scan) //ok
		}

		groupmy := group.Group("/my")
		{
			//我的购物车
			groupmy.GET("/cart", handler.Mycart)

			//我上架的商品
			groupmy.GET("/goods", handler.Mygoods) //ok

			//收藏商品
			groupmy.PATCH("/new_star", handler.Addstar) //ok

			//取消收藏
			groupmy.POST("/cancellation", handler.Cancelstar) //ok

			//个人信息
			groupmy.GET("/message", handler.Mymessage) //ok

			//未完成的订单->订单反馈
			groupmy.GET("/goods/unfinish", handler.UnFinish) //ok

			//确认完成
			groupmy.POST("/goods/finish", handler.Finsh) //ok

			// 修改昵称
			groupmy.GET("/name", handler.ChangeNickname)
		}

		//根据内容对商品的summary搜索
		group.POST("/search", handler.Search) //ok

		//主页推送
		group.GET("/homepage", handler.Homepage) //ok

		//消息通知
		group.GET("/message", handler.Returnmsg) //ok
	}
	//购买后返回联系方式的接口
	router.StaticFS("/images/way", http.Dir("./goods\\way"))
	//查看商品时，返回的图片
	router.StaticFS("/images/avatar", http.Dir("./goods\\avatar"))
	return router
}

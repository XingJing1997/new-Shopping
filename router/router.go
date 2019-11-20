package router

import (
	"github.com/gin-gonic/gin"
	"shopping/config"
	"shopping/controller"
	"shopping/util/tools"
)

func InitRouter() {
	router := gin.Default()
	router.Use(tools.Cors())
	user := router.Group("/user")
	{
		user.POST("/register", controller.RegisterHandler)
		user.GET("/login", controller.LoginHandler)
		user.Use(tools.ValidateToken())
		user.POST("/modify", controller.ModifyUser)
		user.POST("modifyInfo", controller.ModifyUserInfo)
		user.POST("/pushImage", controller.UploadImage)
		user.GET("/showInfo", controller.ShowUserInfo)
	}
	router.Use(tools.ValidateToken())

	UserProduct := router.Group("/userProduct")
	{
		UserProduct.GET("/showProducts", controller.ShowProducts)
	}
	UserCart := router.Group("/userCart")
	{
		UserCart.POST("/addUserCart", controller.AddUserCarts)
		UserCart.GET("/DeleteUserCarts", controller.DeleteUserCarts)
		UserCart.POST("/UpdateCarts", controller.UpdateCarts)
		UserCart.GET("/PrintCarts", controller.PrintCarts)
	}
	Local := router.Group("/local")
	{
		Local.POST("/AddShipping", controller.AddShipping)
		Local.POST("/UpdateShippings", controller.UpdateShippings)
		Local.GET("/DeleteShipping", controller.DeleteShipping)
		Local.GET("/PrintShipping", controller.PrintShipping)
	}
	Order := router.Group("/order")
	{
		Order.POST("/CreateOrderByCart", controller.CreateOrderByCart)
		Order.POST("/CreateOrder", controller.CreateOrder)
		Order.GET("/DeleteOrder", controller.DeleteOrder)
		Order.GET("/UpdateOrder", controller.UpdateOrder)
		Order.GET("/PrintAllOrder", controller.PrintAllOrder)
		Order.GET("/PrintOrderStatus", controller.PrintOrderStatus)
		Order.GET("/PrintOrderDetail", controller.PrintOrderDetail)
	}
	Point := router.Group("point")
	{
		//发行积分，参数：积分发行商ID，具体账户（手机号），有效期
		Point.POST("/PushPoint", controller.PushPoint)
		//2。显示用户未领取的积分 参数: account == phone
		Point.GET("/ShowNoPoint", controller.ShowNoPoint)
		//3.显示用户积分 参数：userid
		Point.GET("/GetUserPoint", controller.GetUserPoint)
		//4.用户领取积分，参数publish 的ID和用户ID
		Point.GET("/GetPoint", controller.GetPoint)
	}
	router.Run(":" + config.GetPort())
}

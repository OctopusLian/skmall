/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 16:32:06
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 18:49:24
 */
package router

import (
	"skmall/middleware"

	"skmall/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	user_group := router.Group("/user")
	product_group := router.Group("/product")
	seckill_group := router.Group("/seckill")

	UserRouter(user_group)
	ProductRouter(product_group)
	SeckillRouter(seckill_group)
}

func UserRouter(router *gin.RouterGroup) {
	// 用户端注册登录
	router.POST("/send_email", api.SendEmail)
	router.POST("/front_user_register", api.FrontUserRegister)
	router.POST("/front_user_login", api.FrontUserLogin)

	// 管理端登录
	router.POST("/admin_login", api.AdminLogin)
	router.GET("/get_front_users", middleware.JwtTokenValid, api.FrontUserList)
}

func ProductRouter(router *gin.RouterGroup) {
	router.GET("/get_product_list", middleware.JwtTokenValid, api.GetProductList)
	router.POST("/product_add", middleware.JwtTokenValid, api.ProductAdd)
	router.POST("/product_del", middleware.JwtTokenValid, api.ProductDel)
	router.GET("/to_product_edit", middleware.JwtTokenValid, api.ProductToEdit)
	router.POST("/do_product_edit", middleware.JwtTokenValid, api.ProductDoEdit)
}

func SeckillRouter(router *gin.RouterGroup) {
	// 管理端
	router.GET("/get_seckill_list", middleware.JwtTokenValid, api.GetSeckillList)
	router.GET("/get_products", middleware.JwtTokenValid, api.GetProducts)
	router.POST("/seckill_add", middleware.JwtTokenValid, api.SecKillAdd)
	router.POST("/seckill_del", middleware.JwtTokenValid, api.SecKillDel)
	router.GET("/seckill_to_edit", middleware.JwtTokenValid, api.SeckillToEdit)
	router.POST("/seckill_do_edit", middleware.JwtTokenValid, api.ProductDoEdit)

	// 前端列表
	router.GET("/front/get_seckill_list", api.GetFrontSeckillList)
	// 前端详情
	router.GET("/front/seckill_detail", middleware.JwtTokenFrontValid, api.SecKillDetail)
	// 秒杀接口
	router.POST("/front/seckill", middleware.JwtTokenFrontValid, api.SecKill)
	// 获取下单结果
	router.GET("/front/get_seckill_result", middleware.JwtTokenFrontValid, api.GetSeckillResult)
}

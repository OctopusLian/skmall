/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:42:54
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:42:56
 */
package user

import (
	"skmall/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	// 用户端注册登录
	router.POST("/send_email", SendEmail)
	router.POST("/front_user_register", FrontUserRegister)
	router.POST("/front_user_login", FrontUserLogin)

	// 管理端登录
	router.POST("/admin_login", AdminLogin)
	router.GET("/get_front_users", middle_ware.JwtTokenValid, FrontUserList)
}

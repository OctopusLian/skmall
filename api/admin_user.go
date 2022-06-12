/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 18:40:44
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 18:45:47
 */
package api

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"skmall/controller"
	"skmall/utils"
)

func AdminLogin(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	in := controller.AdminUserRequest{
		Username: username,
		Password: password,
	}

	// grpc通信
	au := controller.AdminUser{}
	err := au.AdminUserLogin(ctx, &in, &controller.AdminUserResponse{})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "用户名或密码错误",
		})
	} else {
		admin_token, err1 := utils.GenToken(username, utils.AdminUserExpireDuration, utils.AdminUserSecretKey)

		if err1 != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "token错误",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":        rep.Code,
			"msg":         rep.Msg,
			"admin_token": admin_token,
			"user_name":   rep.UserName,
		})
	}

}

func FrontUserList(ctx *gin.Context) {
	currentPage := ctx.DefaultQuery("currentPage", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "没有查询到数据",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":         200,
		"msg":          "成功",
		"front_users":  rep.FrontUsers,
		"total":        rep.Total,
		"current_page": rep.Current,
		"page_size":    rep.PageSize,
	})

}

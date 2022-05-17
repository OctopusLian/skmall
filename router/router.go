/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:44:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:44:21
 */
package all_router

import (
	"skmall/controller/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	user_group := router.Group("/user")
	user.Router(user_group)
}

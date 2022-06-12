/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-12 15:47:31
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-12 16:15:21
 */
package controller

import (
	"context"
	"fmt"
	"skmall/data_source"
	"skmall/models"
	"skmall/utils"
	"strconv"

	"github.com/pkg/errors"
)

type AdminUser struct{}

type AdminUserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AdminUserResponse struct {
	Code     int32  `json:"code,omitempty"`
	Msg      string `json:"msg,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

func (a *AdminUser) AdminUserLogin(ctx context.Context, in *AdminUserRequest, out *AdminUserResponse) error {

	user_name := in.Username
	password := in.Password

	fmt.Println("++++++++++++++")
	fmt.Println(user_name)
	fmt.Println(password)

	md5_password := utils.Md5pwd(password)

	admin_user := models.AdminUser{}

	result := data_source.Db.Where("user_name = ?", user_name).Where("password =?", md5_password).Find(&admin_user)

	if result.Error != nil {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return errors.New("用户名或密码错误")
	}

	out.Code = 200
	out.Msg = "登录成功"
	out.UserName = admin_user.UserName
	return nil
}

type FrontUsersRequest struct {
	CurrentPage int32 `json:"current_page,omitempty"`
	Pagesize    int32 `json:"pagesize,omitempty"`
}

type FrontUser struct {
	Email      string `json:"email,omitempty"`
	Desc       string `json:"desc,omitempty"`
	Status     string `json:"status,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type FrontUsersResponse struct {
	Code       int32        `json:"code,omitempty"`
	Msg        string       `json:"msg,omitempty"`
	FrontUsers []*FrontUser `json:"front_users,omitempty"`
	Total      int32        `json:"total,omitempty"`
	Current    int32        `json:"current,omitempty"`
	PageSize   int32        `json:"page_size,omitempty"`
}

func (a *AdminUser) FrontUserList(ctx context.Context, in *FrontUsersRequest, out *FrontUsersResponse) error {

	currentPage := in.CurrentPage
	pageSize := in.Pagesize

	/*
		current offset limit
		1       0        2       2 * (1 - 1)
		2       2        2		 2 * (2 - 1)
		3       4         2		2 * (3 -1 )

		offset = limit * (current - 1)
	*/

	offsetNum := pageSize * (currentPage - 1)

	users := []models.FrontUser{}
	result := data_source.Db.Limit(pageSize).Offset(offsetNum).Find(&users)
	fmt.Println(len(users))

	if result.Error != nil {
		out.Code = 500
		out.Msg = "没有查询到数据"
	}

	var count int32
	users_count := []models.FrontUser{}
	data_source.Db.Find(&users_count).Count(&count)

	front_users_rep := []*FrontUser{}

	for _, front_user := range users {
		front_user_rep := FrontUser{}
		front_user_rep.Email = front_user.Email
		front_user_rep.Desc = front_user.Desc
		front_user_rep.Status = strconv.FormatInt(int64(front_user.Status), 10)
		front_user_rep.CreateTime = front_user.CreateTime.Format("2006-01-02 15:04:05")
		front_users_rep = append(front_users_rep, &front_user_rep)
	}
	out.Code = 200
	out.Msg = "成功"
	out.FrontUsers = front_users_rep
	out.Total = count
	out.Current = currentPage
	out.PageSize = pageSize
	return nil

}

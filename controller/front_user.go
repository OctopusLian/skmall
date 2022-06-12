package controller

import (
	"context"
	"fmt"
	"skmall/data_source"
	"skmall/models"
	"skmall/utils"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

type fu struct{}

var c = cache.New(60*time.Second, 10*time.Second)

type FrontUserRequest struct {
	Email     string `json:"email,omitempty"`
	Code      string `json:"code,omitempty"`
	Password  string `json:"password,omitempty"`
	Reassword string `json:"reassword,omitempty"`
}

type FrontUserResponse struct {
	Code     int32  `json:"code,omitempty"`
	Msg      string `json:"msg,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

func (f *fu) FrontUserRegister(ctx context.Context, in *FrontUserRequest, out *FrontUserResponse) error {
	// 用户注册
	//data_source.Db
	email := in.Email
	captche := in.Code
	password := in.Password

	code, is_ok := c.Get(email)

	if is_ok {
		if code != captche {
			out.Code = 500
			out.Msg = "邮箱验证码不正确"

		} else {
			// 保存数据到数据库

			md5_password := utils.Md5pwd(password)
			front_user := models.FrontUser{
				Email:      email,
				Password:   md5_password,
				Status:     1,
				CreateTime: time.Now(),
			}

			data_source.Db.Create(&front_user)

			out.Code = 200
			out.Msg = "注册成功,请登录"
		}
	} else {
		out.Code = 500
		out.Msg = "注册失败，请重新尝试"
	}

	// 验证码是否正确

	return nil

}

type FrontUserMailRequest struct {
	Email string `json:"email,omitempty"`
}

func (f *fu) FrontUserSendEmail(ctx context.Context, in *FrontUserMailRequest, out *FrontUserResponse) error {
	// 发送邮件
	email := in.Email

	front_user := models.FrontUser{}
	var count int
	data_source.Db.Where("email =?", email).Find(&front_user).Count(&count)

	if count < 1 {
		rand_num := utils.GenRandNum(6)
		utils.SendEmail(email, rand_num)
		c.Set(email, rand_num, cache.DefaultExpiration)
		out.Code = 200
		out.Msg = "发送成"

	} else {
		out.Code = 500
		out.Msg = "邮箱又存在，请使用其他邮箱"
	}

	return nil

}

func (f *fu) FrontUserLogin(ctx context.Context, in *FrontUserRequest, out *FrontUserResponse) error {

	email := in.Email
	password := in.Password

	md5_password := utils.Md5pwd(password)

	fmt.Println(email)
	fmt.Println(md5_password)

	front_user := models.FrontUser{}
	var count int
	data_source.Db.Where("email = ?", email).Where("password = ?", md5_password).Find(&front_user).Count(&count)

	if count < 1 {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return errors.New("用户名或密码错误")
	} else {
		out.Code = 200
		out.Msg = "登录成功"
		out.UserName = email
		return nil
	}

}

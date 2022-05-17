/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-17 23:40:03
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-17 23:40:05
 */
package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	_ "skmall/user_srv/data_source"
	//"zhiliao_user_srv/subscriber"
	//"zhiliao_user_srv/handler"
	//example "zhiliao_user_srv/proto/example"\

	"skmall/user_srv/controller"
	admin_user "skmall/user_srv/proto/admin_user"
	front_user "skmall/user_srv/proto/front_user"

	"github.com/micro/go-grpc"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("zhiliao.user.srv.zhiliao_user_srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//example.RegisterExampleHandler(service.Server(), new(handler.Example))
	front_user.RegisterFrontUserHandler(service.Server(), new(controller.FrontUser))
	admin_user.RegisterAdminUserHandler(service.Server(), new(controller.AdminUser))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("zhiliao.user.srv.zhiliao_user_srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("zhiliao.user.srv.zhiliao_user_srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

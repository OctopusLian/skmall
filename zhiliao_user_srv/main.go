package main

import (
	"zhiliao_user_srv/handler"
	pb "zhiliao_user_srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("zhiliao_user_srv"),
	)

	// Register handler
	pb.RegisterZhiliao_user_srvHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

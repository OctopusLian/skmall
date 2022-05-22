package main

import (
	"zhiliao_product_srv/handler"
	pb "zhiliao_product_srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("zhiliao_product_srv"),
	)

	// Register handler
	pb.RegisterZhiliao_product_srvHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

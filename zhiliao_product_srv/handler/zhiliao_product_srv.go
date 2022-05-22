package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	zhiliao_product_srv "zhiliao_product_srv/proto"
)

type Zhiliao_product_srv struct{}

// Return a new handler
func New() *Zhiliao_product_srv {
	return &Zhiliao_product_srv{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Zhiliao_product_srv) Call(ctx context.Context, req *zhiliao_product_srv.Request, rsp *zhiliao_product_srv.Response) error {
	log.Info("Received Zhiliao_product_srv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Zhiliao_product_srv) Stream(ctx context.Context, req *zhiliao_product_srv.StreamingRequest, stream zhiliao_product_srv.Zhiliao_product_srv_StreamStream) error {
	log.Infof("Received Zhiliao_product_srv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&zhiliao_product_srv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Zhiliao_product_srv) PingPong(ctx context.Context, stream zhiliao_product_srv.Zhiliao_product_srv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&zhiliao_product_srv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

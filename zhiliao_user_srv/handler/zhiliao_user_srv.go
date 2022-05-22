package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	zhiliao_user_srv "zhiliao_user_srv/proto"
)

type Zhiliao_user_srv struct{}

// Return a new handler
func New() *Zhiliao_user_srv {
	return &Zhiliao_user_srv{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Zhiliao_user_srv) Call(ctx context.Context, req *zhiliao_user_srv.Request, rsp *zhiliao_user_srv.Response) error {
	log.Info("Received Zhiliao_user_srv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Zhiliao_user_srv) Stream(ctx context.Context, req *zhiliao_user_srv.StreamingRequest, stream zhiliao_user_srv.Zhiliao_user_srv_StreamStream) error {
	log.Infof("Received Zhiliao_user_srv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&zhiliao_user_srv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Zhiliao_user_srv) PingPong(ctx context.Context, stream zhiliao_user_srv.Zhiliao_user_srv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&zhiliao_user_srv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

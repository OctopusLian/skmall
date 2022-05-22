package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	zhiliao_web "zhiliao_web/proto"
)

type Zhiliao_web struct{}

// Return a new handler
func New() *Zhiliao_web {
	return &Zhiliao_web{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Zhiliao_web) Call(ctx context.Context, req *zhiliao_web.Request, rsp *zhiliao_web.Response) error {
	log.Info("Received Zhiliao_web.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Zhiliao_web) Stream(ctx context.Context, req *zhiliao_web.StreamingRequest, stream zhiliao_web.Zhiliao_web_StreamStream) error {
	log.Infof("Received Zhiliao_web.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&zhiliao_web.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Zhiliao_web) PingPong(ctx context.Context, stream zhiliao_web.Zhiliao_web_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&zhiliao_web.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

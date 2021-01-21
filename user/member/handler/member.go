package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	member "member/proto"
)

type Member struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Member) Call(ctx context.Context, req *member.Request, rsp *member.Response) error {
	log.Info("Received Member.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Member) Stream(ctx context.Context, req *member.StreamingRequest, stream member.Member_StreamStream) error {
	log.Infof("Received Member.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&member.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Member) PingPong(ctx context.Context, stream member.Member_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&member.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

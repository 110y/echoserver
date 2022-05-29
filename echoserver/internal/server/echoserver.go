package server

import (
	"context"

	"google.golang.org/grpc/metadata"

	echoserverpb "github.com/110y/echoserver/echoserver/api/v1"
)

var _ echoserverpb.EchoServerServer = (*echoserver)(nil)

type echoserver struct{}

func (s *echoserver) Echo(ctx context.Context, req *echoserverpb.EchoRequest) (*echoserverpb.EchoResponse, error) {
	res := new(echoserverpb.EchoResponse)
	res.Message = req.Message

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		res.Headers = make(map[string]*echoserverpb.HeaderValue, md.Len())

		for key, val := range md {
			res.Headers[key] = &echoserverpb.HeaderValue{Value: val}
		}
	}

	return res, nil
}

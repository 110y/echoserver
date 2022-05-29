package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	echocallerpb "github.com/110y/echoserver/echocaller/api/v1"
	echoserverpb "github.com/110y/echoserver/echoserver/api/v1"
)

var _ echocallerpb.EchoCallerServer = (*echocaller)(nil)

type echocaller struct {
	echoserverClient echoserverpb.EchoServerClient
}

func (s *echocaller) Echo(ctx context.Context, req *echoserverpb.EchoRequest) (*echoserverpb.EchoResponse, error) {
	res, err := s.echoserverClient.Echo(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to call echoserver: %s", err.Error())
	}

	return res, nil
}

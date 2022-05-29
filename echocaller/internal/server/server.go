package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/110y/servergroup"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	echocallerpb "github.com/110y/echoserver/echocaller/api/v1"
	echoserverpb "github.com/110y/echoserver/echoserver/api/v1"
	"github.com/110y/echoserver/internal/httputil"
)

var (
	_ servergroup.Server  = (*Server)(nil)
	_ servergroup.Stopper = (*Server)(nil)
)

func NewServer(port int, echoServerAddr string, echoServerTransportInsecure bool) *Server {
	grpcServer := grpc.NewServer()

	echocallerServer := new(echocaller)

	reflection.Register(grpcServer)
	channelz.RegisterChannelzServiceToServer(grpcServer)
	echocallerpb.RegisterEchoCallerServer(grpcServer, echocallerServer)

	mux := runtime.NewServeMux()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if httputil.IsGRPCRequest(r) {
			grpcServer.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	})

	httpServer := &http.Server{Handler: h2c.NewHandler(handler, &http2.Server{})}

	return &Server{
		port:                        port,
		mux:                         mux,
		httpServer:                  httpServer,
		echocallerServer:            echocallerServer,
		echoServerAddr:              echoServerAddr,
		echoserverTransportInsecure: echoServerTransportInsecure,
	}
}

type Server struct {
	port                        int
	mux                         *runtime.ServeMux
	httpServer                  *http.Server
	echoServerAddr              string
	echocallerServer            *echocaller
	echoserverTransportInsecure bool
}

func (s *Server) Start(ctx context.Context) error {
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	if s.echoserverTransportInsecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	}

	// dialer := grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
	//     return (&net.Dialer{}).DialContext(ctx, "tcp4", addr)
	// })
	// opts = append(opts, dialer)

	conn, err := grpc.DialContext(ctx, s.echoServerAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to create a grpc connection to echoserver: %w", err)
	}

	s.echocallerServer.echoserverClient = echoserverpb.NewEchoServerClient(conn)

	gatewayOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	echocallerConn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", s.port), gatewayOpts...)
	if err != nil {
		return fmt.Errorf("failed to dial to the grpc server: %w", err)
	}

	if err = echocallerpb.RegisterEchoCallerHandler(ctx, s.mux, echocallerConn); err != nil {
		return fmt.Errorf("failed to register the echo server handler client: %w", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on the port: %w", err)
	}

	if err := s.httpServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve the grpc server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); httputil.IsUnexpectedListenerError(err) {
		return fmt.Errorf("failed to shutdown the http server: %w", err)
	}

	return nil
}

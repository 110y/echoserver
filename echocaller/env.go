package main

type environment struct {
	GRPCPort                    int    `env:"GRPC_PORT,default=5000"`
	EchoServerAddr              string `env:"ECHO_SERVER_ADDR,required"`
	EchoServerTransportInsecure bool   `env:"ECHO_SERVER_TRANSPORT_INSECURE,default=false"`
}

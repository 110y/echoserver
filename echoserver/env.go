package main

type environment struct {
	GRPCPort int `env:"GRPC_PORT,default=5000"`
}

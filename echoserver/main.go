package main

import (
	"context"
	"fmt"
	"os"

	"github.com/110y/run"
	"github.com/110y/servergroup"
	"github.com/sethvargo/go-envconfig"

	"github.com/110y/echoserver/echoserver/internal/server"
)

func main() {
	run.Run(start)
}

func start(ctx context.Context) int {
	env := new(environment)
	if err := envconfig.Process(ctx, env); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load environment variables: %s", err.Error())
		return 1
	}

	var group servergroup.Group

	gs := server.NewServer(env.GRPCPort)
	group.Add(gs)

	if err := group.Start(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start or stop the server: %s", err.Error())
		return 1
	}

	return 0
}

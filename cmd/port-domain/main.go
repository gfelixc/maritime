package main

import (
	"errors"
	"os/signal"
	"syscall"

	"github.com/gfelixc/maritime/port/ui/server"
	"golang.org/x/net/context"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	err := server.BootGRPCServer(ctx)
	if err != nil && !errors.Is(ctx.Err(), context.Canceled) {
		panic(err)
	}
}

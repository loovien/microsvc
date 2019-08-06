package main

import (
	"context"
	"github.com/labstack/gommon/log"
	appGRPC "github.com/loovien/microsvc/grpc"
	"github.com/loovien/microsvc/service/v1"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	todoSvr := v1.NewTodoService()
	go func() {
		err := appGRPC.ServerRUN(ctx, todoSvr, "localhost:8000")
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	<-ctx.Done()
	log.Printf("receive signal: %v\n", sig)
}

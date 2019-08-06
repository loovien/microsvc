package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/loovien/microsvc/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func ServerRUN(ctx context.Context, server v1.TodoServiceServer, addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	v1.RegisterTodoServiceServer(srv, server)
	go func() {
		err := srv.Serve(listen)
		panic(err)
	}()
	log.Printf("server run at: %s \n", addr)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	sig := <-sigChan
	srv.GracefulStop()
	<-ctx.Done()
	return errors.New(fmt.Sprintf("server stop for receive terminal signal: %v\n", sig))
}

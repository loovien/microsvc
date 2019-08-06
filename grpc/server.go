package grpc

import (
	"context"
	"github.com/loovien/microsvc/api/v1"
	"google.golang.org/grpc"
	"net"
)

func ServerRUN(ctx context.Context, server v1.TodoServiceServer, addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	v1.RegisterTodoServiceServer(srv, server)
	return srv.Serve(listen)
}

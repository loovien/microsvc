package main

import (
	"context"
	GRPC "github.com/loovien/microsvc/grpc"
	"github.com/loovien/microsvc/service/v1"
	"log"
)

func main() {
	addr := "localhost:8000"
	ctx := context.Background()
	todoSvr := v1.NewTodoService()
	err := GRPC.ServerRUN(ctx, todoSvr, addr)
	if err != nil {
		log.Printf("grpc start failure: %v\n", err)
	}
	log.Printf("grpc sever stop: %v\n", err)
}

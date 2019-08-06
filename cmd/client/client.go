package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/loovien/microsvc/api/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	addr := "localhost:8000"
	clientConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect grpc server: %s failure: %v\n", addr, err)
		return
	}
	defer clientConn.Close()
	todoClient := v1.NewTodoServiceClient(clientConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// NewTodo(ctx, todoClient)

	GetTodoList(ctx, todoClient)

}

func NewTodo(ctx context.Context, todoClient v1.TodoServiceClient) {

	todoReq := &v1.TodoRequest{
		Title:       "hello todo client",
		Description: "learn grpc client",
		Reminder:    &timestamp.Timestamp{Seconds: time.Now().Unix()},
	}
	resp, err := todoClient.NewTodo(ctx, todoReq)
	if err != nil {
		log.Fatalf("create todo list failure: %v\n", err)
	}
	log.Printf("create todo list success: %v\n", resp)

}

func GetTodoList(ctx context.Context, client v1.TodoServiceClient) {

	todoListReq := &v1.ListTodoRequest{
	}
	resp, err := client.ListTodo(ctx, todoListReq)
	if err != nil {
		log.Fatalf("query todo list failure: %v\n", err)
	}
	log.Printf("query todo list success: %v\n", resp)
}

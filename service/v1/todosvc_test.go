package v1

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/loovien/microsvc/api/v1"
	"testing"
	"time"
)

func TestTodoService_ListTodo(t *testing.T) {
	todo := NewTodoService()
	resp, err := todo.ListTodo(context.Background(), &v1.ListTodoRequest{
		StartTime: &timestamp.Timestamp{Seconds: time.Now().Unix() - 1000},
		Title: "luowen",
	})
	t.Log(resp, err)
}
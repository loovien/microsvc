package v1

import (
	"context"
	"github.com/loovien/microsvc/api/v1"
	"testing"
)

func TestTodoService_ListTodo(t *testing.T) {
	todo := NewTodoService()
	resp, err := todo.ListTodo(context.Background(), &v1.ListTodoRequest{})
	t.Log(resp, err)
}
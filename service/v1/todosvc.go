package v1

import (
	"context"
	"github.com/loovien/microsvc/api/v1"
	"github.com/loovien/microsvc/model"
	"time"
)

const apiVersion = "v1"

type todoService struct {
}

func NewTodoService() *todoService {
	return &todoService{}
}

func (todo *todoService) NewTodo(ctx context.Context, req *v1.TodoRequest) (*v1.TodoResponse, error) {
	mtodo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		Reminder:    req.Reminder.GetSeconds(),
		CreatedAt:   time.Now().Unix(),
		IsDeleted:   false,
	}
	id, err := mtodo.CreatedTodo()
	if err != nil {
		return nil, err
	}
	return &v1.TodoResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

func (todo *todoService) ListTodo(ctx context.Context, req *v1.ListTodoRequest) (*v1.ListTodoResponse, error) {
	return nil, nil
}

func (todo *todoService) UpdateTodo(ctx context.Context, req *v1.UpdatedTodoRequest) (*v1.UpdatedTodoResponse, error) {
	return nil, nil
}

func (todo *todoService) DeletedTodo(ctx context.Context, req *v1.DeleteTodoRequest) (*v1.DeleteTodoResponse, error) {
	return nil, nil
}

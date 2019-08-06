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
	criteria := model.NewCriteria()
	var bind []interface{}
	if req.StartTime.Seconds > 0 {
		bind = append(bind, req.StartTime.Seconds)
		criteria.ANDCondition = append(criteria.ANDCondition, "created_at > ?")
	}
	if req.EndTime.Seconds > 0 {
		bind = append(bind, req.EndTime.Seconds)
		criteria.ANDCondition = append(criteria.ANDCondition, "created_at <= ?")
	}

	if len(req.Title) > 0 {
		bind = append(bind, "%"+req.Title+"%")
		criteria.ANDCondition = append(criteria.ANDCondition, "title like ?")
	}

	mtodo := &model.Todo{}
	todoList, err := mtodo.GetTodoList(criteria, bind)
}

func (todo *todoService) UpdateTodo(ctx context.Context, req *v1.UpdatedTodoRequest) (*v1.UpdatedTodoResponse, error) {
	return nil, nil
}

func (todo *todoService) DeletedTodo(ctx context.Context, req *v1.DeleteTodoRequest) (*v1.DeleteTodoResponse, error) {
	return nil, nil
}

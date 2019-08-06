package model

type Todo struct {
	Id          int64
	Title       string
	Description string
	Reminder    int64
	CreatedAt   int64
	IsDeleted   bool
}

func NewTodo() *Todo {
	return &Todo{}
}

func (todo *Todo) TableName() string {
	return "todo"
}

func (todo *Todo) CreatedTodo() (int64, error) {
	_, err := x.Insert(todo)
	return todo.Id, err
}

func (todo *Todo) GetTodoList(criteria *Criteria) ([]Todo, error) {
	var resp []Todo
	session := criteria.Apply(x)
	err := session.Find(&resp)
	return resp, err
}

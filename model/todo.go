package model

type Todo struct {
	Id          int
	Title       string
	Description string
	Reminder    int64
	CreatedAt   int64 `xorm:"column(created_at)"`
	IsDeleted   bool
}

func NewTodo() *Todo {
	return &Todo{}
}

func (todo *Todo) TableName() string {
	return "todo"
}

func (todo *Todo) CreatedTodo() (int64, error) {
	return x.InsertOne(todo)
}

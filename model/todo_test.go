package model

import "testing"

func TestTodo_GetTodoList(t *testing.T) {
	var resp []Todo
	err := x.Where("title like ?", "%luowen%").Find(&resp)
	t.Log(err)
	t.Log(resp)
}
package repository

import (
	todo "todo_go/internal/domain"
	
)

type inMemoryTodoList struct{
	todos []todo.Todo
	nextID int
}

func NewinMemoryTodoRepository() *inMemoryTodoList{
	return &inMemoryTodoList{
		todos: []todo.Todo{},
		nextID: 1,
	}
}


func (r *inMemoryTodoList) Create(t todo.Todo) todo.Todo{
	t.ID = r.nextID
	r.nextID++
	r.todos = append(r.todos, t)
	return t
}

func (r *inMemoryTodoList) GetAll() []todo.Todo {
	return r.todos
}


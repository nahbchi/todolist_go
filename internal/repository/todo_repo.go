package repository

import (
	todo "todo_go/internal/domain"
	"fmt"
	
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

func (r *inMemoryTodoList) Update(id int, updated todo.Todo) (todo.Todo, error) {
	for i, t := range r.todos {
		if t.ID == id {

			r.todos[i].Done = updated.Done
			r.todos[i].Title = updated.Title 
			return r.todos[i], nil
		}
	}
	return todo.Todo{}, fmt.Errorf("todo not found")
}

func (r *inMemoryTodoList) Delete(id int) error {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

func (r *inMemoryTodoList) GetAll() []todo.Todo {
	return r.todos
}


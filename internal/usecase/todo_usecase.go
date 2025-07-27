	package usecase

	import(
	todo "todo_go/internal/domain"
	)

	type TodoRepository interface{
		GetAll() []todo.Todo
		Create(todo.Todo) todo.Todo
		Update(id int, update todo.Todo) (todo.Todo, error)
		Delete(id int) error
	}

	type TodoUseCase struct{
		repo TodoRepository
	}

	func NewUseCase( r TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		repo: r,
	}
	}

	func(uc *TodoUseCase) GetTodos() []todo.Todo {
		return uc.repo.GetAll()
	}

	func (uc *TodoUseCase) CreateTodo(t todo.Todo) todo.Todo {
		return uc.repo.Create(t)
	}

	func( uc *TodoUseCase) UpdateTodo(id int, update todo.Todo) (todo.Todo, error){
		return uc.repo.Update(id, update)
	}

	func (uc *TodoUseCase) DeleteTodo(id int) error {
	return uc.repo.Delete(id)
}

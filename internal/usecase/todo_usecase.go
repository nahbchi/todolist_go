	package usecase

	import(
	todo "todo_go/internal/domain"
	)

	type TodoRepository interface{
		GetAll() []todo.Todo
		Create(todo.Todo) todo.Todo
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


package handler

import (
	"encoding/json"
	"net/http"	
	domain "todo_go/internal/domain"
	"todo_go/internal/usecase"
)


type TodoHandler struct{
	usecase *usecase.TodoUseCase
}

func (h *TodoHandler) getTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.usecase.GetTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}


func NewTodoHandler (uc *usecase.TodoUseCase) *TodoHandler{
	return &TodoHandler{usecase: uc}
}

func(h* TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		h.getTodos(w,r)
	case http.MethodPost:
		h.createTodo(w, r)

	default:
		http.Error(w, "Methodo Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request){
	var t  domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	return 
	}

	created := h.usecase.CreateTodo(t)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}


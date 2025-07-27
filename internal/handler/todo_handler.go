package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	domain "todo_go/internal/domain"
	"todo_go/internal/usecase"
)

type TodoHandler struct {
	usecase *usecase.TodoUseCase
}

func NewTodoHandler(uc *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{usecase: uc}
}

func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.URL.Path == "/todos" {
		switch r.Method {
		case http.MethodGet:
			h.getTodos(w, r)
		case http.MethodPost:
			h.createTodo(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/todos/") {
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid todo ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodPut:
			h.updateTodo(w, r, id)
		case http.MethodDelete:
			h.deleteTodo(w, r, id)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func (h *TodoHandler) getTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.usecase.GetTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	var t domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	created := h.usecase.CreateTodo(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *TodoHandler) updateTodo(w http.ResponseWriter, r *http.Request, id int) {
	var updated domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todo, err := h.usecase.UpdateTodo(id, updated)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) deleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.usecase.DeleteTodo(id); err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

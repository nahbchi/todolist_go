package main

import (
	"fmt"
	"log"
	"net/http"
	"todo_go/internal/handler"
	"todo_go/internal/repository"
	"todo_go/internal/usecase"
)

func main(){
	repo := repository.NewinMemoryTodoRepository()
	uc := usecase.NewUseCase(repo)
	h := handler.NewTodoHandler(uc)

  http.Handle("/todos", h)
  http.Handle("/", http.FileServer(http.Dir("./web")))


  fmt.Println("server running at http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080",nil))
}



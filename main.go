package main

import (
	"fmt"

	"github.com/matheusalecksander/go-todo/src/features/todo"
	inmemory "github.com/matheusalecksander/go-todo/src/features/todo/infrastructure/repository/in-memory"
)

func main() {
	repository := inmemory.NewTodoRepository()

	fmt.Println("Iniciando")

	todoService := todo.NewTodoService(*repository)

	myTodo := todoService.CreateTodo("Aprender GO")
	todoInRepo, _ := todoService.FindAll()

	fmt.Println(todoInRepo)

	myTodo.CompleteTodo()

	todoService.UpdateTodo(myTodo)
	todoInRepo, _ = todoService.FindAll()

	fmt.Println(todoInRepo)

	myTodo.UncompleteTodo()

	todoService.UpdateTodo(myTodo)

	todoInRepo, _ = todoService.FindAll()

	fmt.Println(todoInRepo)

	err := todoService.DeleteTodo(myTodo.Uuid)

	if err != nil {
		panic(err)
	}

	todoInRepo, err = todoService.FindAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(todoInRepo)
}

package inmemory

import (
	"fmt"

	todo_structs "github.com/matheusalecksander/go-todo/src/features/todo/domain/structs"
)

const todoNotFoundMessage = "todo with id %s not found"

type TodoRepository struct {
	todos []todo_structs.STodo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: make([]todo_structs.STodo, 0),
	}
}

func (repo *TodoRepository) SaveTodo(todo *todo_structs.STodo) {
	repo.todos = append(repo.todos, *todo)
}

func (repo *TodoRepository) Create(todo todo_structs.STodo) error {
	repo.todos = append(repo.todos, todo)
	return nil
}

func (repo *TodoRepository) FindAll() ([]todo_structs.STodo, error) {
	return repo.todos, nil
}

func (repo *TodoRepository) FindById(id string) (todo_structs.STodo, error) {
	for _, t := range repo.todos {
		if t.Uuid == id {
			return t, nil
		}
	}
	return todo_structs.STodo{}, fmt.Errorf(todoNotFoundMessage, id)
}

func (repo *TodoRepository) Update(todo todo_structs.STodo) error {
	for i, t := range repo.todos {
		if t.Uuid == todo.Uuid {
			repo.todos[i] = todo
			return nil
		}
	}
	return fmt.Errorf(todoNotFoundMessage, todo.Uuid)
}

func (repo *TodoRepository) Delete(id string) error {
	for i, t := range repo.todos {
		if t.Uuid == id {
			repo.todos = append(repo.todos[:i], repo.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf(todoNotFoundMessage, id)
}

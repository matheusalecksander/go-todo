package todo

import (
	"github.com/google/uuid"
	todo_structs "github.com/matheusalecksander/go-todo/src/features/todo/domain/structs"
	inmemory "github.com/matheusalecksander/go-todo/src/features/todo/infrastructure/repository/in-memory"
)

type TodoService struct {
	todoRepository inmemory.TodoRepository
}

func NewTodoService(repository inmemory.TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: repository,
	}
}

func (service *TodoService) CreateTodo(description string) todo_structs.STodo {
	uuid := uuid.New().String()
	todo := todo_structs.STodo{
		Uuid:        uuid,
		Description: description,
		Completed:   false,
	}

	service.todoRepository.Create(todo)

	return todo
}

func (service *TodoService) UpdateTodo(todo todo_structs.STodo) {
	_, err := service.FindById(todo.Uuid)
	if err != nil {
		panic(err)
	}
	service.todoRepository.Update(todo)
}

func (service *TodoService) FindById(id string) (todo_structs.STodo, error) {
	todo, err := service.todoRepository.FindById(id)
	if err != nil {
		return todo_structs.STodo{}, err
	}

	return todo, nil
}

func (service *TodoService) FindAll() ([]todo_structs.STodo, error) {
	todo, err := service.todoRepository.FindAll()
	if err != nil {
		return []todo_structs.STodo{}, err
	}

	return todo, nil
}

func (service *TodoService) DeleteTodo(id string) error {
	return service.todoRepository.Delete(id)
}

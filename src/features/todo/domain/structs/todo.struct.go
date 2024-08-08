package todo_structs

type STodo struct {
	Uuid        string
	Description string
	Completed   bool
}

func (todo *STodo) CompleteTodo() {
	todo.Completed = true
}

func (todo *STodo) UncompleteTodo() {
	todo.Completed = false
}

package todo

import todoService "github.com/fastworkco/go-boilerplate/internal/service/todo"

type TodoHandlerDependencies struct {
	TodoService todoService.TodoService
}

type TodoHandler struct {
	todoService todoService.TodoService
}

func NewTodoHandler(dependencies TodoHandlerDependencies) *TodoHandler {
	return &TodoHandler{
		todoService: dependencies.TodoService,
	}
}

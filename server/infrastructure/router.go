package infrastructure

import (
	"server/interfaces/controllers"
)

type Router struct {
	controllers.TodoController
}

func RouterInit() (*Router, error) {
	return &Router{
		TodoController: *controllers.NewTodoController(NewSqlHandler()),
	}, nil
}

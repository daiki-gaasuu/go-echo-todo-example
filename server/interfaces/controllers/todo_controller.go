package controllers

import (
	"server/domain"
	"server/interfaces/database"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) GetTodos(c echo.Context) (err error) {
	todos, err := controller.Interactor.GetAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todos)
	return
}

func (controller *TodoController) GetTodoById(c echo.Context, id int) (err error) {
	todo, err := controller.Interactor.GetById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) PostTodo(c echo.Context) (err error) {
	t := domain.Todo{}
	c.Bind(&t)
	todo, err := controller.Interactor.Add(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) UpdateTodo(c echo.Context) (err error) {
	t := domain.Todo{}
	c.Bind(&t)
	todo, err := controller.Interactor.Update(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) DeleteTodo(c echo.Context, id int) (err error) {
	todo := domain.Todo{
		ID: id,
	}
	err = controller.Interactor.DeleteById(todo)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(204, "")
	return
}

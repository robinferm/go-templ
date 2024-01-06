package main

import (
	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	todos := GetTodos()
	component := Todos(todos)
	return component.Render(c.Request().Context(), c.Response())
}

func DeleteTodoByIdHandler(c echo.Context) error {
	id := c.Param("id")
	DeleteTodoById(id)
	return c.Redirect(302, "/")
}

package main

import (
	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	todos := GetTodos()
	component := TodoPage(todos)
	return component.Render(c.Request().Context(), c.Response())
}

func AddTodoHandler(c echo.Context) error {
	title := c.FormValue("title")
	AddTodo(title)

	todos := GetTodos()
	component := todoTable(todos)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func DeleteTodoByIdHandler(c echo.Context) error {
	id := c.Param("id")
	DeleteTodoById(id)

	todos := GetTodos()
	component := todoTable(todos)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

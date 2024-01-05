package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/robinferm/go-templ/internal/db"
	"github.com/robinferm/go-templ/views/todoPage"
)

func GetTodosHandler(c echo.Context) error {
	todos := db.GetTodos()
	component := todoPage.Todos(todos)
	return component.Render(c.Request().Context(), c.Response())
}

func DeleteTodoByIdHandler(c echo.Context) error {
	id := c.Param("id")
	db.DeleteTodoById(id)
	return c.Redirect(302, "/")
}

// func PostHandler(c echo.Context) error {
// 	c.Request().ParseForm()

// 	component := todoPage.Counter(global.Count)
// 	return component.Render(c.Request().Context(), c.Response())
// }

package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/robinferm/go-templ/views/todoPage"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func GetHandler(c echo.Context) error {
	component := todoPage.Page(global.Count)
	return component.Render(c.Request().Context(), c.Response())
}

func PostHandler(c echo.Context) error {
	c.Request().ParseForm()

	if c.FormValue("global") == "global" {
		global.Count++
	}

	component := todoPage.Counter(global.Count)
	return component.Render(c.Request().Context(), c.Response())
}

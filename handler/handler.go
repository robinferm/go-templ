package handler

import (
	"github.com/labstack/echo"
	"github.com/robinferm/go-templ/templ"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func GetHandler(c echo.Context) error {
	component := templ.Page(global.Count)
	return component.Render(c.Request().Context(), c.Response())
}

func PostHandler(c echo.Context) error {
	c.Request().ParseForm()

	if c.FormValue("global") == "global" {
		global.Count++
	}

	component := templ.Counter(global.Count)
	return component.Render(c.Request().Context(), c.Response())
}

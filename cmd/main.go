package main

import (
	"github.com/labstack/echo"
	"github.com/robinferm/go-templ/handler"
)

func main() {
	e := echo.New()

	e.GET("/", handler.GetHandler)

	e.POST("/", handler.PostHandler)

	e.Logger.Fatal(e.Start("localhost:1323"))

}

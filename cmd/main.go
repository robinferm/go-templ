package main

import (
	"github.com/labstack/echo"
	"github.com/robinferm/go-templ/internal/handlers"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.GetHandler)

	e.POST("/", handlers.PostHandler)

	e.Logger.Fatal(e.Start("localhost:1323"))

}

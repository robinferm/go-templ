package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	ConnectDB()
	defer CloseDB()

	e := echo.New()

	e.GET("/", GetTodosHandler)
	e.POST("/delete/:id", DeleteTodoByIdHandler)
	// e.POST("/", handlers.PostHandler)

	e.Logger.Fatal(e.Start("localhost:1323"))

}

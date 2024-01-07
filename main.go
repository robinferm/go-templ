package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	ConnectDB()
	defer CloseDB()

	e := echo.New()

	e.GET("/", GetTodosHandler)
	e.DELETE("/delete/:id", DeleteTodoByIdHandler)
	e.POST("/add", AddTodoHandler)
	e.POST("/toggle/:id", ToggleTodoByIdHandler)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}

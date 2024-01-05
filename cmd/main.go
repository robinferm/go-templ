package main

import (
	"github.com/labstack/echo/v4"
	"github.com/robinferm/go-templ/internal/db"
	"github.com/robinferm/go-templ/internal/handlers"
)

func main() {

	db.ConnectDB()
	defer db.CloseDB()

	e := echo.New()

	e.GET("/", handlers.GetTodosHandler)
	e.POST("/delete/:id", handlers.DeleteTodoByIdHandler)
	// e.POST("/", handlers.PostHandler)

	e.Logger.Fatal(e.Start("localhost:1323"))

}

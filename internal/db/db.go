package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/robinferm/go-templ/internal/models"
)

var conn *pgx.Conn

func ConnectDB() {
	var err error
	url := "postgres://postgres:mysecretpassword@localhost:5432/tododb"
	conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to database")
}

func CloseDB() {
	conn.Close(context.Background())
}

func GetTodos() []models.Todo {
	var todos = []models.Todo{}
	rows, err := conn.Query(context.Background(), "select * from todos")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return todos
}

func DeleteTodoById(id string) {
	fmt.Printf("Deleting todo with id: %s\n", id)
	_, err := conn.Exec(context.Background(), "delete from todos where id=$1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

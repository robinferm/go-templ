package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
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

func GetTodos() []Todo {
	var todos = []Todo{}
	rows, err := conn.Query(context.Background(), "select * from todos")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
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

func AddTodo(title string) {
	_, err := conn.Exec(context.Background(), "insert into todos (title, done) values ($1, $2)", title, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

func DeleteTodoById(id string) {
	_, err := conn.Exec(context.Background(), "delete from todos where id=$1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

func ToggleTodoById(id string) {
	_, err := conn.Exec(context.Background(), "update todos set done = not done where id=$1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("sqlite3", "db/todo_CLA.db")
	if err != nil {
		panic(err)
	}
	tableName := "todos"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			task STRING NOT NULL,
			limitDate STRING NOT NULL,
			status BOOLEAN
		)`, tableName)
	_, err = conn.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

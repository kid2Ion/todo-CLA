package infra

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("sqlite3", "todo_CLA.db")
	if err != nil {
		panic(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

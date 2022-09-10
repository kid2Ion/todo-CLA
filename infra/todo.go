package infra

import (
	"fmt"
	"todo_CLA/domain/model"
	"todo_CLA/domain/repository"
)

type TodoRepository struct {
	SqlHandler
}

func NewTodoRepository(sqlHandler SqlHandler) repository.TodoRepository {
	todoRepository := TodoRepository{sqlHandler}
	return &todoRepository
}

func (todoRepo *TodoRepository) FindAll() ([]*model.Todo, error) {
	rows, err := todoRepo.SqlHandler.Conn.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to get all todos")
	}

	todos := []*model.Todo{}
	for rows.Next() {
		todo := model.Todo{}
		rows.Scan(&todo.Id, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, &todo)
	}

	return todos, nil
}

func (todoRepo *TodoRepository) Find(word string) ([]*model.Todo, error) {
	rows, err := todoRepo.SqlHandler.Conn.Query("SELECT * FROM todos WHERE task LIKE ?", "%"+word+"%")
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("no records")
	}

	todos := []*model.Todo{}
	for rows.Next() {
		todo := model.Todo{}
		rows.Scan(&todo.Id, &todo.Task, &todo.LimitDate, &todo.Status)
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (todoRepo *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.SqlHandler.Conn.Exec("INSERT INTO todos (task,limitDate,status) VALUES (?, ?, ?)", todo.Task, todo.LimitDate, todo.Status)
	return todo, err
}

func (todoRepo *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.SqlHandler.Conn.Exec("UPDATE todos SET task = ?, limitDate = ?, status = ? WHERE id = ?", todo.Task, todo.LimitDate, todo.Status, todo.Id)
	return todo, err
}

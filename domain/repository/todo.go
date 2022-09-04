package repository

import "todo_CLA/domain/model"

type TodoRepository interface {
	FindAll() ([]*model.Todo, error)
	Find(word string) ([]*model.Todo, error)
	Create(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
}

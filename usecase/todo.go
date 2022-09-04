package usecase

import (
	"todo_CLA/domain/model"
	"todo_CLA/domain/repository"
)

type TodoUsecase interface {
	Search(string) (todo []*model.Todo, err error)
	View() (todo []*model.Todo, err error)
	Add(*model.Todo) (err error)
	Edit(*model.Todo) (err error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo: todoRepo}
	return &todoUsecase
}

func (usecase *todoUsecase) Search(word string) (todo []*model.Todo, err error) {
	todo, err = usecase.todoRepo.Find(word)
	return
}

func (usecase *todoUsecase) View() (todo []*model.Todo, err error) {
	todo, err = usecase.todoRepo.FindAll()
	return
}

func (usecase *todoUsecase) Add(todo *model.Todo) (err error) {
	_, err = usecase.todoRepo.Create(todo)
	return
}

func (usecase *todoUsecase) Edit(todo *model.Todo) (err error) {
	_, err = usecase.todoRepo.Update(todo)
	return
}

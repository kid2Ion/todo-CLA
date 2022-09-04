package injector

import (
	"todo_CLA/domain/repository"
	"todo_CLA/handler"
	"todo_CLA/infra"
	"todo_CLA/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	return *sqlHandler
}

func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return infra.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	todoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(todoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return *handler.NewTodoHandler(InjectTodoUsecase())
}

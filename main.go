package main

import (
	"fmt"
	"todo_CLA/handler"
	"todo_CLA/injector"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("start server")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

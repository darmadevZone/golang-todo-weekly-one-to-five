package main

import (
	"lets-golang/controller"
	"lets-golang/docs"
	"lets-golang/models"
	"lets-golang/repositories"
	"lets-golang/services"
	"log"

	"github.com/gin-gonic/gin"
)

// mockTodo
func main()  {
	todos := [] models.Todo{
		{ID: 1, Name: "Todo 1"},
		{ID: 2, Name: "Todo 2"},
		{ID: 3, Name: "Todo 3"},
		{ID: 4, Name: "Todo 4"},			
		{ID: 5, Name: "Todo 5"},
		{ID: 6, Name: "Todo 6"},
		{ID: 7, Name: "Todo 7"},
		{ID: 8, Name: "Todo 8"},
		{ID: 9, Name: "Todo 9"},
		{ID: 10, Name: "Todo 10"},
	}

	todoRepository := repositories.NewTodoRepository(todos)
	todoService:= services.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)
	docs.PointerFunc()
	engine := gin.Default()
	
	engine.GET("/", todoController.GetTodoAll)
	engine.GET("/:id",todoController.GetTodoById)
	// TODO impl: Update todo
	engine.PUT("/:id",todoController.UpdateTodo)
	engine.DELETE("/:id", todoController.DeleteTodo)
	engine.POST("/create", todoController.CreateTodo)
	
	
	log.Fatal(engine.Run(":8080"))
}

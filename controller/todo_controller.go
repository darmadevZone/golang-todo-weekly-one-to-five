package controller

import (
	"lets-golang/dto"
	"lets-golang/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	GetTodoAll(ctx *gin.Context)
	GetTodoById(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	CreateTodo(ctx *gin.Context)
}

type TodoController struct {
	todoService services.ITodoService
}


// UpdateTodo implements ITodoController.
func (t *TodoController) UpdateTodo(ctx *gin.Context) {
	todoId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var input dto.Todo
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateTodo, err := t.todoService.UpdateTodo(uint(todoId), &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": updateTodo,
	})
}

func NewTodoController(todoService services.ITodoService) ITodoController {
	return &TodoController{todoService}
}

// GetTodoAll implements ITodoController.
func (t *TodoController) GetTodoAll(ctx *gin.Context) {
	todos, err := t.todoService.GetTodoAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// GetTodoById implements ITodoController.
func (t *TodoController) GetTodoById(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	item, err := t.todoService.GetTodoById(uint(itemId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

// DeleteTodo implements ITodoController.
func (t *TodoController) DeleteTodo(ctx *gin.Context) {
	todoId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	todo, err := t.todoService.DeleteTodo(uint(todoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

// CreateTodo implements ITodoController.
func (t *TodoController) CreateTodo(ctx *gin.Context) {
	var input dto.CreatTodo
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := t.todoService.CreateTodo(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

package services

import (
	"lets-golang/dto"
	"lets-golang/models"
	"lets-golang/repositories"
)

type ITodoService interface {
	GetTodoAll() (*[]models.Todo, error)
	GetTodoById(id uint) (*models.Todo, error)
	UpdateTodo(id uint, input *dto.Todo) (*models.Todo, error)
	DeleteTodo(todoId uint) (*models.Todo, error)
	CreateTodo( *dto.CreatTodo) (*models.Todo, error)
}

type TodoService struct {
	todoRepository repositories.ITodoRepository
}


func NewTodoService(todoRepository repositories.ITodoRepository) ITodoService {
	return &TodoService{todoRepository}
}

// GetTodoAll implements ITodoService.
func (t *TodoService) GetTodoAll() (*[]models.Todo, error) {
	return t.todoRepository.GetTodoAll()
}

// GetTodoById implements ITodoService.
func (t *TodoService) GetTodoById(id uint) (*models.Todo, error) {
	return t.todoRepository.GetTodoById(id)
}

// UpdateTodo implements ITodoService.
func (t *TodoService) UpdateTodo(id uint, input *dto.Todo) (*models.Todo, error) {
	todo, err := t.todoRepository.GetTodoById(id)
	if err != nil {
		return nil, err
	}
	if input.Name != nil {
		todo.Name = *input.Name
	}
	return t.todoRepository.UpdateTodo(id, todo)
}

// DeleteTodo implements ITodoService.
func (t *TodoService) DeleteTodo(todoId uint) (*models.Todo, error) {
	todo, err := t.todoRepository.DeleteTodo(todoId)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// CreateTodo implements ITodoService.
func (t *TodoService) CreateTodo(input *dto.CreatTodo) (*models.Todo, error) {
	var todo models.Todo
	todos,err := t.todoRepository.GetTodoAll()
	if err != nil {
		return nil, err
	}
	todo.ID = uint(len(*todos) + 1)
	todo.Name = *input.Name
	return t.todoRepository.CreateTodo(&todo)
}

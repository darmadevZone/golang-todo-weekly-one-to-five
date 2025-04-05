package repositories

import (
	"errors"
	"lets-golang/models"
)

type ITodoRepository interface {
	GetTodoAll() (*[]models.Todo, error)
	GetTodoById(id uint) (*models.Todo, error)
	UpdateTodo(id uint, updateTodo *models.Todo) (*models.Todo, error)
	DeleteTodo(todoId uint) (*models.Todo, error)
	CreateTodo(newTodo *models.Todo) (*models.Todo, error)
}

type TodoRepository struct {
	Todos []models.Todo
}


func NewTodoRepository(todos []models.Todo) ITodoRepository {
	return &TodoRepository{todos}
}

// GetTodoAll implements ITodoRepository.
func (t *TodoRepository) GetTodoAll() (*[]models.Todo, error) {
	todos := t.Todos
	if len(todos) == 0 {
		return nil, errors.New("todo が見つかりませんでした。")
	}
	return &t.Todos,nil 
}

// GetTodoById implements ITodoRepository.
func (t *TodoRepository) GetTodoById(id uint) (*models.Todo, error) {
	for _, v := range t.Todos {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("todo not found")
}

// UpdateTodo implements ITodoRepository.
func (t *TodoRepository) UpdateTodo(id uint, updateTodo *models.Todo) (*models.Todo, error) {
	for i, v := range t.Todos {
		if v.ID == updateTodo.ID {
			t.Todos[i] = *updateTodo
			return &t.Todos[i], nil
		}
	}
	return nil, errors.New("todoが更新できませんでした。")
}

// DeleteTodo implements ITodoRepository.
func (t *TodoRepository) DeleteTodo(todoId uint) (*models.Todo, error) {
	for i, v := range t.Todos {
		if v.ID == todoId {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			return &v, nil
		}
	}
	return nil, errors.New("todo not found")
}

// CreateTodo implements ITodoRepository.
func (t *TodoRepository) CreateTodo(newTodo *models.Todo) (*models.Todo, error) {
	t.Todos = append(t.Todos, *newTodo)
	return newTodo, nil 
}

package services

import (
	"errors"
	"todo-api/models"
	"todo-api/repositories"
)

type TodoService struct {
	Repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

func (s *TodoService) GetTodos(page, limit int, done *bool, search string) ([]models.Todo, error) {
	return s.Repo.GetAllTodos(page, limit, done, search)
}

func (s *TodoService) FindById(id int) (models.Todo, error) {
	return s.Repo.FindById(id)
}

func (s *TodoService) CreateTodo(todo models.Todo) (models.Todo, error) {

	if todo.Title == "" {
		return todo, errors.New("title is required")
	}

	return s.Repo.CreateTodo(todo)
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.Repo.DeleteTodo(id)
}

func (s *TodoService) UpdateTodo(id int, todo models.Todo) (models.Todo, error) {

	if todo.Title == "" {
		return todo, errors.New("title is required")
	}

	return s.Repo.UpdateTodo(id, todo)
}

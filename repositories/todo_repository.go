package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"todo-api/models"
)

type TodoRepository interface {
	GetAllTodos(page, limit int, done *bool, search string) ([]models.Todo, error)
	FindById(id int) (models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(id int) error
	UpdateTodo(id int, todo models.Todo) (models.Todo, error)
}

type PostgresTodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &PostgresTodoRepository{DB: db}
}

func (r *PostgresTodoRepository) GetAllTodos(page, limit int, done *bool, search string) ([]models.Todo, error) {
	offset := (page - 1) * limit

	query := "SELECT id, title, done FROM todos WHERE 1=1"
	args := []interface{}{}
	argIndex := 1

	// filter done
	if done != nil {
		query += fmt.Sprintf(" AND done = $%d", argIndex)
		args = append(args, *done)
		argIndex++
	}

	// search title
	if search != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argIndex)
		args = append(args, "%"+search+"%")
		argIndex++
	}

	// pagination
	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *PostgresTodoRepository) FindById(id int) (models.Todo, error) {
	query := "SELECT id,title,done FROM todos WHERE id=$1"

	var todo models.Todo

	err := r.DB.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Done)

	if err == sql.ErrNoRows {
		return todo, errors.New("invalid id")
	}

	if err != nil {
		return todo, err
	}

	return todo, nil

}

func (r *PostgresTodoRepository) CreateTodo(todo models.Todo) (models.Todo, error) {

	query := "INSERT INTO todos (title,done) VALUES ($1, $2) RETURNING id"

	err := r.DB.QueryRow(query, todo.Title, todo.Done).Scan(&todo.ID)

	if err != nil {
		return todo, nil
	}

	return todo, nil
}

func (r *PostgresTodoRepository) DeleteTodo(id int) error {
	_, err := r.DB.Exec("DELETE FROM todos WHERE id=$1", id)

	return err
}

func (r *PostgresTodoRepository) UpdateTodo(id int, todo models.Todo) (models.Todo, error) {
	query := "UPDATE todos SET title=$1, done=$2 WHERE id=$3"

	result, err := r.DB.Exec(query, todo.Title, todo.Done, id)
	if err != nil {
		return todo, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return todo, err
	}

	if rowsAffected == 0 {
		return todo, errors.New("todo not found")
	}

	todo.ID = id
	return todo, nil

}

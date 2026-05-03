package repositories

import (
	"database/sql"
	"todo-api/models"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) CreateUser(user models.User) (models.User, error) {
	query := "INSERT INTO users (email,password) VALUES ($1, $2) RETURNING id,role"

	err := r.DB.QueryRow(query, user.Email, user.Password).Scan(&user.ID, &user.Role)

	return user, err

}

func (r *PostgresUserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	query := "SELECT id, email , password, role FROM users WHERE email=$1"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	return user, err

}

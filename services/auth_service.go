package services

import (
	"errors"
	"todo-api/models"
	"todo-api/repositories"
	"todo-api/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) Register(email, password string) (models.User, error) {

	if email == "" || password == "" {
		return models.User{}, errors.New("email and password required")
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:    email,
		Password: string(hashed),
	}

	return s.UserRepo.CreateUser(user)

}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// generate JWT
	token, err := utils.GenerateToken(user.ID, user.Role)

	if err != nil {
		return "", err
	}

	return token, nil

}

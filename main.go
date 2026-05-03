package main

import (
	"fmt"
	"net/http"
	"todo-api/config"
	"todo-api/handlers"
	"todo-api/middlewares"
	"todo-api/repositories"
	"todo-api/services"

	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	// todo repo
	repo := repositories.NewTodoRepository(config.DB)
	service := services.NewTodoService(repo)
	handler := handlers.NewTodoHandler(service)

	// user repo
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewAuthService(userRepo)
	userHandler := handlers.NewAuthHandler(userService)

	r := chi.NewRouter()
	r.Use(middlewares.Logger)

	// public routes
	// Router Auth
	r.Post("/auth/login", userHandler.Login)
	r.Post("/auth/register", userHandler.Register)

	// private routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.Auth)
		// Router Todos
		r.Get("/todos", handler.GetTodos)
		r.Get("/todos/{id}", handler.GetTodoById)
		r.Post("/todos", handler.CreateTodo)
		r.Delete("/todos/{id}", middlewares.AdminOnly(handler.DeleteTodo))
		r.Put("/todos/{id}", handler.UpdateTodo)

	})

	fmt.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", r)

}

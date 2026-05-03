package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/models"
	"todo-api/services"
	"todo-api/utils"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	Service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{Service: service}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	pageStr := query.Get("page")
	limitStr := query.Get("limit")
	doneStr := query.Get("done")
	search := query.Get("q")

	page := 1
	limit := 10

	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			utils.Error(w, http.StatusBadRequest, "invalid page")
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			utils.Error(w, http.StatusBadRequest, "invalid limit")
			return
		}

	}

	var done *bool
	if doneStr != "" {
		val, err := strconv.ParseBool(doneStr)
		if err != nil {
			utils.Error(w, http.StatusBadRequest, "invalid done value")
			return
		}
		done = &val
	}

	todos, err := h.Service.GetTodos(page, limit, done, search)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid format id")
		return
	}

	todo, err := h.Service.FindById(id)

	if err != nil {
		utils.Error(w, http.StatusNotFound, "todo not found")
		return
	}

	utils.JSON(w, http.StatusOK, todo)

}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.Service.CreateTodo(todo)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusCreated, result)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.Service.DeleteTodo(id); err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"message": "Deleted",
	})

}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.Service.UpdateTodo(id, todo)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, result)

}

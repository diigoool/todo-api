package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/services"
	"todo-api/utils"
)

type AuthHandlder struct {
	Service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandlder {
	return &AuthHandlder{Service: s}
}

func (h *AuthHandlder) Register(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	user, err := h.Service.Register(req.Email, req.Password)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(w, http.StatusCreated, user)
}

func (h *AuthHandlder) Login(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})

}

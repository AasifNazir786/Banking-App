package handlers

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHODS ARE ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	var user models.AccountHolder
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding user: %v", err), http.StatusBadRequest)
		return
	}

	err = h.service.RegisterUser(user.UserName, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "user created successfully")
}

func (h *UserHandler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHODS ARE ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	var user models.AccountHolder
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding user: %v", err), http.StatusBadRequest)
		return
	}

	token, err := h.service.AuthenticateUser(user.UserName, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}

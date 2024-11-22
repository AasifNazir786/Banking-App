package handlers

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	service *services.AccountService
}

func NewAccountHandler(service *services.AccountService) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	var accountRequest models.Account
	if err := json.NewDecoder(r.Body).Decode(accountRequest); err != nil {

		http.Error(w, "INVALID INPUT", http.StatusBadRequest)
		return
	}

	account, err := h.service.CreateAccount(accountRequest.Name, accountRequest.Balance, accountRequest.AccountType)

	if err != nil {

		http.Error(w, "FAILED TO CREATE ACCOUNT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(w, "ONLY GET REQUEST IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {

		http.Error(w, "NOT VALID ID", http.StatusNotAcceptable)
		return
	}

	account, err := h.service.RetrieveAccount(id)

	if err != nil {
		http.Error(w, "FAILED TO RETRIEVE ACCOUNT", http.StatusNoContent)
		return
	}

	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(w, "ONLY GET REQUEST IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	accounts, err := h.service.RetrieveAllAccounts()

	if err != nil {
		http.Error(w, "FAILED TO RETRIEVE ACCOUNT", http.StatusNoContent)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

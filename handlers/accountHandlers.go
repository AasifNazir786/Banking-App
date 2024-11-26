package handlers

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"fmt"
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
	if err := json.NewDecoder(r.Body).Decode(&accountRequest); err != nil {

		http.Error(w, "INVALID INPUT", http.StatusBadRequest)
		return
	}

	account, err := h.service.CreateAccount(accountRequest.Name, accountRequest.Balance, accountRequest.AccountType)

	if err != nil {

		http.Error(w, "FAILED TO CREATE ACCOUNT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

	w.Header().Add("Content-Type", "application/json")
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func (h *AccountHandler) Transfer(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	toId, err := strconv.Atoi(r.URL.Query().Get("toId"))

	if err != nil {
		http.Error(w, "not valid toId", http.StatusNotAcceptable)
		return
	}

	fromId, err := strconv.Atoi(r.URL.Query().Get("fromId"))

	if err != nil {
		http.Error(w, "not valid fromId", http.StatusNotAcceptable)
		return
	}

	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))

	if amount <= 0 {
		http.Error(w, "amount should be positive", http.StatusNotAcceptable)
		return
	}

	err = h.service.TransferFrom_To(fromId, toId, float64(amount))

	if err != nil {
		http.Error(w, "Can't transfer money", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "successful",
	})
}

func (h *AccountHandler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {

		http.Error(w, "error parsing id", http.StatusBadRequest)
		return
	}

	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))

	if err != nil {

		http.Error(w, "error parsing amount", http.StatusBadRequest)
		return
	}

	err = h.service.Withdraw_(id, float64(amount))

	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := fmt.Sprintf("Successfully withdrew from account ID: %d", id)

	json.NewEncoder(w).Encode(map[string]string{
		"message ": message,
	})
}

func (h *AccountHandler) DepositHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {

		http.Error(w, "error parsing id", http.StatusBadRequest)
		return
	}

	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))

	if err != nil {

		http.Error(w, "error parsing amount", http.StatusBadRequest)
		return
	}

	err = h.service.Deposit(id, float64(amount))

	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := fmt.Sprintf("Successfully Deposited from Account Id: %d", id)

	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})

}

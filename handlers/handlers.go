package handlers

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	// var account models.Account
	// type accountRequest struct {
	// 	Name        string             `json:"name"`
	// 	Balance     float64            `json:"balance"`
	// 	AccountType models.AccountType `json:"account_type"`
	// }
	var accountRequest models.Account

	if err := json.NewDecoder(r.Body).Decode(&accountRequest); err != nil {
		http.Error(w, "INVALID INPUT", http.StatusNotFound)
		return
	}

	account := services.CreateAccount(accountRequest.Name, accountRequest.Balance, accountRequest.AccountType)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))

	account, err := services.Deposit(id, float64(amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	transaction := services.RecordTransaction(id, "deposit", "deposited funds: "+strconv.Itoa(amount), float64(amount))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"account":     account,
		"transaction": transaction,
	})
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))

	account, err := services.Withdraw(id, float64(amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction := services.RecordTransaction(id, "withdraw", "Credited funds: "+strconv.Itoa(amount), float64(amount))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"account":     account,
		"transaction": transaction,
	})
}

func CheckBalanceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}

	account, err := services.CheckBalance(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account)
}

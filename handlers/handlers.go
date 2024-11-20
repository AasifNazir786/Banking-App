package handlers

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/services"
	"Go-GitHub-Projects/Banking-App/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

var accounts = storage.GetAccounts()

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
	}

	transaction := services.RecordTransaction(id, "deposit", "deposited funds", float64(amount))

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

	for i := range accounts {
		if accounts[i].Id == id {
			if accounts[i].Balance < float64(amount) {
				http.Error(w, "Insufficient Balance", http.StatusBadRequest)
				return
			}
			accounts[i].Balance -= float64(amount)
			json.NewEncoder(w).Encode(accounts[i])
			return
		}
	}
	http.Error(w, "Account Not Found", http.StatusNotFound)
}

func CheckBalanceHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	for _, account := range accounts {
		if account.Id == id {
			json.NewEncoder(w).Encode(account)
			return
		}
	}
	http.Error(w, "Account Not Found", http.StatusNotFound)
}

package handlers

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
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

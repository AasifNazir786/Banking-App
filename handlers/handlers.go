package handlers

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "INVALID INPUT", http.StatusNotFound)
		return
	}

	account.Id = len(storage.Accounts) + 1
	storage.Accounts = append(storage.Accounts, account)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)

}

func Deposite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))

	for i := range storage.Accounts {
		if storage.Accounts[i].Id == id {
			storage.Accounts[i].Balance += float64(amount)
			json.NewEncoder(w).Encode(storage.Accounts[i])
			return
		}
	}
	http.Error(w, "Account Not Found", http.StatusNotFound)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHOD IS ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))

	for i := range storage.Accounts {
		if storage.Accounts[i].Id == id {
			if storage.Accounts[i].Balance < float64(amount) {
				http.Error(w, "Insufficient Balance", http.StatusBadRequest)
				return
			}
			storage.Accounts[i].Balance -= float64(amount)
			json.NewEncoder(w).Encode(storage.Accounts[i])
			return
		}
	}
	http.Error(w, "Account Not Found", http.StatusNotFound)
}

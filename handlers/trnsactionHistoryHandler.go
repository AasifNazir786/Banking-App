package handlers

import (
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "ONLY GET REQUEST IS ALLOWED", http.StatusBadRequest)
		return
	}

	var err error

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Please enter valid id", http.StatusBadRequest)
	}

	transactions, err := services.GetTransactions(id)
	if err != nil {
		http.Error(w, "No Content Found", http.StatusNoContent)
	}

	json.NewEncoder(w).Encode(transactions)
}

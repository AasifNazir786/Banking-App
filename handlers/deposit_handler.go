package handlers

import (
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
	"strconv"
)

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

package handlers

import (
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
	"strconv"
)

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

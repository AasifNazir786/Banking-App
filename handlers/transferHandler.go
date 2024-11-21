package handlers

import (
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST Method IS ALLOWED", http.StatusBadRequest)
		return
	}

	var err error

	toId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Not valid Id", http.StatusBadRequest)
		return
	}
	fromId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Not valid Id", http.StatusBadRequest)
		return
	}
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		http.Error(w, "Not valid Id", http.StatusBadRequest)
		return
	}

	accountsMap, err := services.TransferFrom_To(toId, fromId, float64(amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(accountsMap)
}

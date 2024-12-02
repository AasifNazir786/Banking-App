package handlers

import (
	"Go-GitHub-Projects/Banking-App/services"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {

	return &TransactionHandler{
		transactionService: transactionService,
	}
}

// func(h *TransactionHandler)SaveTransactionHandler(w http.ResponseWriter, r *http.Request){

// 	if r.Method != http.MethodPost{
// 		http.Error(w, "Only Post methods are Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	var transaction models.Transaction
// 	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil{
// 		http.Error(w, err.Error(), http.StatusNotAcceptable)
// 		return
// 	}

// 	h.transactionService.SaveTransaction(transaction.AccountId, )
// }

// func(h *TransactionHandler)GetAllTransactionsHandler(w http.ResponseWriter, r *http.Request){

// }

func (h *TransactionHandler) GetAllByAccountIdHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only Get Methods Are Allowed", http.StatusMethodNotAllowed)
		return
	}

	accountId, err := strconv.Atoi(r.URL.Query().Get("accountId"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	transactions, err := h.transactionService.RetrieveAllByAccountId(accountId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) GetAllBetweenDatesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only Get Methods are Allowed", http.StatusMethodNotAllowed)
		return
	}

	accountId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	startDate, err := time.Parse("2006-01-02", r.URL.Query().Get("sDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	endDate, err := time.Parse("2006-01-02", r.URL.Query().Get("eDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	transactions, err := h.transactionService.GetAllByDates(accountId, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(transactions)
}

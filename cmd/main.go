package main

import (
	"Go-GitHub-Projects/Banking-App/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/create-account", handlers.CreateAccountHandler)
	http.HandleFunc("/deposit", handlers.DepositHandler)
	http.HandleFunc("/withdraw", handlers.WithdrawHandler)
	http.HandleFunc("/balance", handlers.CheckBalanceHandler)
	http.HandleFunc("/transfer", handlers.TransferHandler)
	http.HandleFunc("/get_transactions", handlers.TransactionsHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

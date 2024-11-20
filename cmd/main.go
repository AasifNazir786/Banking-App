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

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

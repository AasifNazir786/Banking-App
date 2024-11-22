package main

import (
	"Go-GitHub-Projects/Banking-App/handlers"
	"Go-GitHub-Projects/Banking-App/services"
	"Go-GitHub-Projects/Banking-App/storage"
	"log"
	"net/http"
)

func main() {

	if err := storage.InitDB(); err != nil {

		log.Fatalf("Failed to initialize the database: %v", err)
	}

	defer storage.CloseDB()

	db := storage.GetDB()

	accountStorage := storage.NewAccountStorage(db)
	accountService := services.NewAccountService(accountStorage)
	accountHandler := handlers.NewAccountHandler(accountService)

	http.HandleFunc("/create-account", accountHandler.CreateAccountHandler)
	http.HandleFunc("/deposit", handlers.DepositHandler)
	http.HandleFunc("/withdraw", handlers.WithdrawHandler)
	http.HandleFunc("/balance", handlers.CheckBalanceHandler)
	http.HandleFunc("/transfer", handlers.TransferHandler)
	http.HandleFunc("/get_transactions", handlers.TransactionsHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

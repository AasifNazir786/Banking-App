package main

import (
	"Go-GitHub-Projects/Banking-App/handlers"
	"Go-GitHub-Projects/Banking-App/middleware"
	"Go-GitHub-Projects/Banking-App/services"
	"Go-GitHub-Projects/Banking-App/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {

	if err := storage.InitDB(); err != nil {

		log.Fatalf("Failed to initialize the database: %v", err)
	}

	defer storage.CloseDB()

	s := "abcdefghi"

	x := s[len("abc"):]
	fmt.Println(x)

	db := storage.GetDB()

	userRepo := storage.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	http.HandleFunc("/register", userHandler.UserRegisterHandler)
	http.HandleFunc("/login", userHandler.UserLoginHandler)

	transactionStorage := storage.NewTransactionStorage(db)
	transactionService := services.NewTransactionService(transactionStorage)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	accountStorage := storage.NewAccountStorage(db)
	accountService := services.NewAccountService(accountStorage, transactionService)
	accountHandler := handlers.NewAccountHandler(accountService)

	http.Handle("/createAccount", middleware.LoggerMiddleware(middleware.AuthMiddleWare(http.HandlerFunc(accountHandler.CreateAccountHandler))))
	http.HandleFunc("/getAccInfo", accountHandler.GetAccountById)
	http.HandleFunc("/getHistoryBWDates", transactionHandler.GetAllBetweenDatesHandler)
	http.HandleFunc("/getAllTransactions", transactionHandler.GetAllByAccountIdHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/storage"
)

var accounts = storage.GetAccounts()

// CreateAccount creates a new account and returns the created account
func CreateAccount(name string, balance float64, accountType models.AccountType) models.Account {
	account := models.Account{
		Id:          len(accounts) + 1,
		Name:        name,
		Balance:     balance,
		AccountType: accountType,
	}
	accounts = append(accounts, account)
	return account
}

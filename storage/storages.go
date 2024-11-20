package storage

import "Go-GitHub-Projects/Banking-App/models"

// var Accounts []models.Account

var accounts = []models.Account{}
var transactions = []models.Transaction{}

func GetAccounts() []models.Account {
	return accounts
}

func GetTransactions() []models.Transaction {
	return transactions
}

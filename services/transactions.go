package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/storage"
	"time"
)

var transactions = storage.GetTransactions()

func RecordTransaction(accountId int, tType, description string, amount float64) models.Transaction {
	transaction := models.Transaction{
		Id:          len(transactions) + 1,
		AccountId:   accountId,
		Type:        tType,
		Amount:      amount,
		TimeStamp:   time.Now().Format(time.RFC3339),
		Description: description,
	}

	transactions = append(transactions, transaction)
	return transaction
}

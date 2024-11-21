package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"errors"
)

func GetTransactions(id int) ([]models.Transaction, error) {

	if transactions == nil {
		return nil, errors.New("no transactions available")
	}

	filteredTransactions := []models.Transaction{}

	for i := range transactions {
		if transactions[i].AccountId == id {
			filteredTransactions = append(filteredTransactions, transactions[i])
		}
	}

	if len(filteredTransactions) == 0 {
		return nil, errors.New("no transaction of given account_id")
	}

	return filteredTransactions, nil
}

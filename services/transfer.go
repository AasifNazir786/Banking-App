package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"errors"
	"strconv"
)

func TransferFrom_To(toId, fromId int, amount float64) (map[string]interface{}, error) {

	var fromAccount, toAccount *models.Account

	for i := range accounts {
		if accounts[i].Id == toId {
			toAccount = &accounts[i]
		}
		if accounts[i].Id == fromId {
			fromAccount = &accounts[i]
		}
	}

	if fromAccount == nil || toAccount == nil {
		return map[string]interface{}{}, errors.New("one of the two accounts is nil")
	}

	if fromAccount.Balance < amount {
		return map[string]interface{}{}, errors.New("Insufficient Balance...")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	RecordTransaction(fromAccount.Id, "transfer", "Transferred funds to account "+strconv.Itoa(toId), amount)
	RecordTransaction(toAccount.Id, "transfer", "Received funds from account "+strconv.Itoa(fromId), amount)

	return map[string]interface{}{
		"fromAccount": fromAccount,
		"toAccount":   toAccount,
	}, nil

}

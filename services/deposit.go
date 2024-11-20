package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"errors"
)

func Deposit(id int, amount float64) (models.Account, error) {

	for i := range accounts {
		if accounts[i].Id == id {
			accounts[i].Balance += amount
			return accounts[i], nil
		}
	}
	return models.Account{}, errors.New("account not found")
}

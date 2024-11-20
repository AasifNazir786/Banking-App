package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"errors"
)

func Withdraw(id int, amount float64) (models.Account, error) {

	for i := range accounts {
		if accounts[i].Id == id {
			if accounts[i].Balance < amount {
				return models.Account{}, errors.New("Insufficient Balance...")
			}
			accounts[i].Balance -= amount
			return accounts[i], nil
		}
	}
	return models.Account{}, errors.New("please enter valid id...")
}

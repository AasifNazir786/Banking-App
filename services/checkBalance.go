package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"errors"
)

func CheckBalance(id int) (models.Account, error) {

	for i := range accounts {
		if accounts[i].Id == id {
			return accounts[i], nil
		}
	}
	return models.Account{}, errors.New("please enter valid id")
}

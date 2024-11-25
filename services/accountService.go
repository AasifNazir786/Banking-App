package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/storage"
	"fmt"
)

type AccountService struct {
	storage *storage.AccountStorage
}

func NewAccountService(storage *storage.AccountStorage) *AccountService {
	return &AccountService{
		storage: storage,
	}
}

func (s *AccountService) CreateAccount(name string, balance float64, accountType models.AccountType) (models.Account, error) {

	account := models.Account{
		Name:        name,
		Balance:     balance,
		AccountType: accountType,
	}

	id, err := s.storage.AddAccount(account)

	if err != nil {

		return models.Account{}, fmt.Errorf("failed to create account: %w", err)
	}
	account.Id = id
	return account, nil
}

func (s *AccountService) RetrieveAccount(id int) (models.Account, error) {

	account, err := s.storage.GetAccountById(id)
	if err != nil {
		return models.Account{}, err
	}
	return account, nil
}

func (s *AccountService) RetrieveAllAccounts() ([]models.Account, error) {

	accounts, err := s.storage.GetAllAccounts()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountService) TransferFrom_To(fromId, toId int, amount float64) error {

	fromAccount, err := s.storage.GetAccountById(fromId)

	if err != nil {
		return err
	}

	toAccount, err := s.storage.GetAccountById(toId)

	if err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return fmt.Errorf("insufficient balance of accountId: %d", fromId)
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	err = s.storage.UpdateAccount(toId, toAccount)

	if err != nil {
		return fmt.Errorf("failed to update to account: %w", err)
	}

	err = s.storage.UpdateAccount(fromId, fromAccount)

	if err != nil {
		return fmt.Errorf("failed to update from account: %w", err)
	}

	return nil
}

func (s *AccountService) Withdraw_(id int, amount float64) (models.Account, error) {

	if amount <= 0 {
		return models.Account{}, fmt.Errorf("invalid withdrawal amount: %.2f", amount)
	}

	account, err := s.storage.GetAccountById(id)

	if err != nil {

		return models.Account{}, err
	}
	if account.Balance < amount {

		return models.Account{}, fmt.Errorf("insufficient balance")
	}
	account.Balance -= amount

	if err := s.storage.UpdateAccount(id, account); err != nil {

		return models.Account{}, fmt.Errorf("could not update account: %w", err)
	}

	return account, nil
}

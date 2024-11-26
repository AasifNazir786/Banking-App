package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/storage"
	"fmt"
)

type AccountService struct {
	storage *storage.AccountStorage
	service *TransactionService
}

func NewAccountService(storage *storage.AccountStorage, service *TransactionService) *AccountService {
	return &AccountService{
		storage: storage,
		service: service,
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

	err = s.storage.UpdateAccount(fromAccount)

	if err != nil {
		return fmt.Errorf("failed to update from account: %w", err)
	}

	toAccount.Balance += amount

	err = s.storage.UpdateAccount(toAccount)

	if err != nil {
		return fmt.Errorf("failed to update to account: %w", err)
	}

	_, err = s.service.SaveTransaction(toId, amount, "transfer", fmt.Sprintf("Received %.2f from account %d", amount, fromId))

	if err != nil {

		return fmt.Errorf("failed to record transaction for reciever: %w", err)
	}
	_, err = s.service.SaveTransaction(fromId, amount, "recieve", fmt.Sprintf("Transfers %.2f from account %d", amount, toId))

	if err != nil {

		return fmt.Errorf("failed to record transaction for sender: %w", err)
	}

	return nil
}

func (s *AccountService) Withdraw_(id int, amount float64) error {

	if amount <= 0 {
		return fmt.Errorf("invalid withdrawal amount: %.2f", amount)
	}

	account, err := s.storage.GetAccountById(id)

	if err != nil {

		return fmt.Errorf("can't retrieve an account")
	}
	if account.Balance < amount {

		return fmt.Errorf("insufficient balance")
	}
	account.Balance -= amount

	if err := s.storage.UpdateAccount(account); err != nil {

		return fmt.Errorf("could not update account: %w", err)
	}

	_, err = s.service.SaveTransaction(id, amount, "withdraw's", fmt.Sprintf("Withdraw's %.2f from account %d", amount, id))

	if err != nil {

		return fmt.Errorf("failed to record withdrawal transaction: %w", err)
	}

	return nil
}

func (s *AccountService) Deposit(id int, amount float64) error {

	account, err := s.storage.GetAccountById(id)
	if err != nil {
		return fmt.Errorf("error retrieving account")
	}

	account.Balance += amount

	if err := s.storage.UpdateAccount(account); err != nil {
		return fmt.Errorf("can't update account")
	}

	_, err = s.service.SaveTransaction(id, amount, "Deposit", fmt.Sprintf("Deposites %.2f from account %d", amount, id))

	if err != nil {

		return fmt.Errorf("failed to record deposit transaction: %w", err)
	}

	return nil
}

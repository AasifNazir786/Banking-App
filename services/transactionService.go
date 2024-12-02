package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/storage"
	"time"
)

type TransactionService struct {
	storage *storage.TransactionStorage
}

func NewTransactionService(storage *storage.TransactionStorage) *TransactionService {

	return &TransactionService{
		storage: storage,
	}
}

func (s *TransactionService) SaveTransaction(accountId int, amount float64,
	tType models.TransactionType, description string) (models.Transaction, error) {

	transaction := models.Transaction{
		AccountId:   accountId,
		Amount:      amount,
		Type:        tType,
		Date:        time.Now(),
		Description: description,
	}

	id, err := s.storage.AddTransaction(transaction)

	if err != nil {

		return models.Transaction{}, err
	}
	transaction.Id = id

	return transaction, nil
}

func (s *TransactionService) GetAllByDates(accountId int, start, end time.Time) ([]models.Transaction, error) {

	transactions, err := s.storage.GetTransactionsFromDB(accountId, start, end)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// func (s *TransactionService) UpdateTransaction(transaction models.Transaction) error {

// 	err := s.storage.UpdateTransaction(transaction)

// 	if err != nil {

// 		return err
// 	}
// 	return nil
// }

func (s *TransactionService) RetrieveAllByAccountId(id int) ([]models.Transaction, error) {

	transactions, err := s.storage.GetAllByAccountId(id)

	if err != nil {

		return nil, err
	}
	return transactions, nil

}

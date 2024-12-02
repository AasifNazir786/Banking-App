package storage

import (
	"Go-GitHub-Projects/Banking-App/models"
	"database/sql"
	"fmt"
	"time"
)

type TransactionStorage struct {
	db *sql.DB
}

func NewTransactionStorage(db *sql.DB) *TransactionStorage {
	return &TransactionStorage{
		db: db,
	}
}

func (s *TransactionStorage) AddTransaction(transaction models.Transaction) (int, error) {

	query := `INSERT INTO transactions (account_id, amount, type, time_stamp, description)
			VALUES ($1, $2, $3, $4, $5) RETURNING $6`

	var id int

	err := s.db.QueryRow(query, transaction.AccountId, transaction.Amount,
		transaction.Type, transaction.Date, transaction.Description).Scan(&id)
	if err != nil {

		return 0, fmt.Errorf("failed to insert transaction %w", err)
	}

	return id, nil
}

func (s *TransactionStorage) GetTransactionsFromDB(accountID int, startDate, endDate time.Time) ([]models.Transaction, error) {

	query := `SELECT * FROM transactions
				WHERE account_id = $1 AND date BETWEEN $2 AND $3`

	rows, err := s.db.Query(query, accountID, startDate, endDate)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.Id, &t.AccountId, &t.Amount, &t.Date, &t.Description); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (s *TransactionStorage) GetTransactionById(id int) (models.Transaction, error) {

	query := `SELECT * FROM transactions
									WHERE id = $1`

	var transaction models.Transaction

	err := s.db.QueryRow(query, id).Scan(&transaction.Id, &transaction.AccountId,
		&transaction.Amount, &transaction.Type, &transaction.Date, &transaction.Description)

	if err != nil {
		return models.Transaction{}, fmt.Errorf("can't retrieve transaction")
	}

	return transaction, nil
}

func (s *TransactionStorage) GetAllTransactions() ([]models.Transaction, error) {

	query := `SELECT * FROM transactions`

	rows, err := s.db.Query(query)

	if err != nil {
		return []models.Transaction{}, fmt.Errorf("can't retrieve transactions")
	}

	transactions := []models.Transaction{}

	for rows.Next() {
		var transaction models.Transaction

		if err := rows.Scan(&transaction.Id, &transaction.AccountId, &transaction.Amount,
			&transaction.Type, &transaction.Date, &transaction.Description); err != nil {

			return []models.Transaction{}, fmt.Errorf("can't retrieve transaction")
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *TransactionStorage) GetAllByAccountId(accountId int) ([]models.Transaction, error) {

	query := `SELECT * FROM transactions
							WHERE account_id = $1`

	rows, err := s.db.Query(query, accountId)

	if err != nil {

		return []models.Transaction{}, fmt.Errorf("can't retrieve transactions")
	}

	transactions := []models.Transaction{}

	for rows.Next() {
		var transaction models.Transaction

		if err := rows.Scan(&transaction.Id, &transaction.AccountId, &transaction.Amount,
			&transaction.Type, &transaction.Date, &transaction.Description); err != nil {

			return []models.Transaction{}, fmt.Errorf("can't retrieve transaction")
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *TransactionStorage) UpdateTransaction(updatedTransaction models.Transaction) error {

	query := `UPDATE transactions
				SET account_id = &1, amount = $2, type = $3, time_stamp = $4, description = $5
				WHERE id = $6`

	_, err := s.db.Exec(query, updatedTransaction.AccountId, updatedTransaction.Amount,
		updatedTransaction.Type, updatedTransaction.Date, updatedTransaction.Description, updatedTransaction.Id)

	if err != nil {

		return fmt.Errorf("not updated transaction")
	}
	return nil
}

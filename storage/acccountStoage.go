package storage

import (
	"Go-GitHub-Projects/Banking-App/models"
	"database/sql"
	"errors"
	"fmt"
)

type AccountStorage struct {
	db *sql.DB
}

func NewAccountStorage(db *sql.DB) *AccountStorage {
	return &AccountStorage{
		db: db,
	}
}

func (d *AccountStorage) AddAccount(account models.Account) (int, error) {

	query := `INSERT INTO accounts (name, balance, account_type)
	VALUES($1, $2, $3) RETURNING id`

	var id int
	err := d.db.QueryRow(query, account.Name, account.Balance, account.AccountType).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert account: %w", err)
	}
	return id, nil
}

func (d *AccountStorage) GetAccountById(id int) (models.Account, error) {
	query := `SELECT * FROM accounts
	WHERE id = $1`

	var account models.Account
	err := d.db.QueryRow(query, id).Scan(&account.Id, &account.Name, &account.Balance, &account.AccountType)

	if err != nil {

		return models.Account{}, err
	}
	return account, nil
}

func (d *AccountStorage) GetAllAccounts() ([]models.Account, error) {
	query := `SELECT * FROM accounts`

	rows, err := d.db.Query(query)
	if err != nil {
		return []models.Account{}, errors.New("unable to fetch the accounts")
	}
	accounts := []models.Account{}
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.Id, &account.Name, &account.Balance, &account.AccountType); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (d *AccountStorage) UpdateAccount(id int, updatedAccount models.Account) error {
	query := `UPDATE accounts
	SET name = $1, balance = $2, account_type = $3
	WHERE id = $4`

	_, err := d.db.Exec(query, updatedAccount.Name, updatedAccount.Balance, updatedAccount.AccountType, id)

	if err != nil {
		return err
	}

	return nil
}

package models

import "time"

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
	Transfer TransactionType = "transfer"
)

type Transaction struct {
	Id          int             `json:"id"`
	AccountId   int             `json:"account_id"`
	Amount      float64         `json:"amount"`
	Type        TransactionType `json:"type"`
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
}

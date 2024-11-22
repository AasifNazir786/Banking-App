package models

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
	TimeStamp   string          `json:"time_stamp"`
	Description string          `json:"description"`
}

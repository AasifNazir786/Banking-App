package models

// enum type
type AccountType string

const (
	CD      AccountType = "cd"
	Savings AccountType = "savings"
	Current AccountType = "current"
)

type Account struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Balance     float64     `json:"balance"`
	AccountType AccountType `json:"account_type"`
}

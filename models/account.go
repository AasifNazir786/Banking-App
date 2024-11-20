package models

type AccountType string

const (
	Checking AccountType = "checking"
	Savings  AccountType = "savings"
	// Current  AccountType = "current"
)

type Account struct {
	Id          int         `json:id`
	Name        string      `json:"name"`
	Balance     float64     `json:"balance"`
	AccountType AccountType `json:"account_type"`
}

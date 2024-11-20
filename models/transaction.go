package models

type Transaction struct {
	Id          int     `json:"id"`
	AccountId   int     `"json:account_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	TimeStamp   string  `json:"time_stamp"`
	Description string  `json:"description"`
}

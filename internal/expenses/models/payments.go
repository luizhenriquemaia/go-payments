package models

import "time"

type PaymentModel struct {
	Id         int64      `json:"id"`
	Expense_id int64      `json:"expense_id"`
	Receipt    string     `json:"receipt"`
	Method     int        `json:"method"`
	Account    int        `json:"account"`
	Paid_at    *time.Time `json:"paid_at"`
	Updated_at time.Time  `json:"updated_at"`
	Created_at time.Time  `json:"created_at"`
}

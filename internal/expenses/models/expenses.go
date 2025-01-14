package models

import "time"

type ExpenseModel struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Cost_center int       `json:"cost_center"`
	Status      int       `json:"status"`
	Bar_code    string    `json:"bar_code"`
	Document    string    `json:"document"`
	Due_date    time.Time `json:"due_date"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}

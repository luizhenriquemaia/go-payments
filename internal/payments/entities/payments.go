package entities

import (
	"go-payments/internal/payments/enums"
	"time"
)

type PaymentEntity struct {
	Id         int64                `json:"id"`
	Expense_id int64                `json:"expense_id"`
	Receipt    string               `json:"receipt"`
	Method     enums.PaymentMethod  `json:"method"`
	Account    enums.PaymentAccount `json:"account"`
	Paid_at    *time.Time           `json:"paid_at"`
	Updated_at time.Time            `json:"updated_at"`
	Created_at time.Time            `json:"created_at"`
}

type PaymentRespEntity struct {
	Id         int64      `json:"id"`
	Expense_id int64      `json:"expense_id"`
	Receipt    string     `json:"receipt"`
	Method     string     `json:"method"`
	Account    string     `json:"account"`
	Paid_at    *time.Time `json:"paid_at"`
	Updated_at time.Time  `json:"updated_at"`
	Created_at time.Time  `json:"created_at"`
}

type PayExpenseEntity struct {
	Expense_id int                  `binding:"required"`
	Method     enums.PaymentMethod  `binding:"required,enum"`
	Account    enums.PaymentAccount `binding:"required,enum"`
}

type AddPaymentDbEntity struct {
	Expense_id int
	Receipt    string
	Method     int
	Account    int
	Paid_at    time.Time
	Updated_at time.Time
	Created_at time.Time
}

func (entity *PayExpenseEntity) GetToDb() *AddPaymentDbEntity {
	now := time.Now().UTC()
	return &AddPaymentDbEntity{
		Expense_id: entity.Expense_id,
		Receipt:    now.Format("200601021504"),
		Method:     int(entity.Method),
		Account:    int(entity.Account),
		Paid_at:    now,
		Updated_at: now,
		Created_at: now,
	}
}

// Adapter method to get payment entity to response
func (entity *PaymentEntity) GetToResp() *PaymentRespEntity {
	return &PaymentRespEntity{
		Id:         entity.Id,
		Expense_id: entity.Expense_id,
		Receipt:    entity.Receipt,
		Method:     entity.Method.String(),
		Account:    entity.Account.String(),
		Paid_at:    entity.Paid_at,
		Updated_at: entity.Updated_at,
		Created_at: entity.Created_at,
	}
}

package factories

import (
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/enums"
	"log"
	"time"
)

type PaymentFactory struct{}

func (factory *PaymentFactory) GetFromDb(
	id int64,
	expense_id int64,
	receipt string,
	method int,
	account int,
	paid_at *time.Time,
	updated_at time.Time,
	created_at time.Time,
) entities.PaymentEntity {
	brazil_tz, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatal("Erro on load America/Sao_Paulo timezone for convert payment times")
	}

	method_enum, _ := enums.GetPaymentMethodByValue(method)
	account_enum, _ := enums.GetPaymentAccountByValue(account)
	var localized_paid_at time.Time
	if paid_at != nil {
		localized_paid_at = paid_at.In(brazil_tz)
	}
	return entities.PaymentEntity{
		Id:         id,
		Expense_id: expense_id,
		Receipt:    receipt,
		Method:     method_enum,
		Account:    account_enum,
		Paid_at:    &localized_paid_at,
		Updated_at: updated_at.In(brazil_tz),
		Created_at: created_at.In(brazil_tz),
	}
}

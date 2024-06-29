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
	description string,
	cost_center int,
	status int,
	bar_code string,
	document string,
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
	status_enum, _ := enums.GetPaymentStatusByValue(status)
	cc_enum, _ := enums.GetCostCenterByValue(cost_center)
	method_enum, _ := enums.GetPaymentMethodByValue(method)
	account_enum, _ := enums.GetPaymentAccountByValue(account)
	return entities.PaymentEntity{
		Id:          id,
		Description: description,
		Cost_center: cc_enum,
		Status:      status_enum,
		Bar_code:    bar_code,
		Document:    document,
		Receipt:     receipt,
		Paid_at:     paid_at,
		Method:      method_enum,
		Account:     account_enum,
		Updated_at:  updated_at.In(brazil_tz),
		Created_at:  created_at.In(brazil_tz),
	}
}

func (factory *PaymentFactory) GetToResp(entity *entities.PaymentEntity) *entities.PaymentEntityResponse {
	return &entities.PaymentEntityResponse{
		Id:          entity.Id,
		Description: entity.Description,
		Cost_center: entity.Cost_center.String(),
		Status:      entity.Status.String(),
		Bar_code:    entity.Bar_code,
		Document:    entity.Document,
		Receipt:     entity.Receipt,
		Paid_at:     entity.Paid_at,
		Method:      entity.Method.String(),
		Account:     entity.Account.String(),
		Updated_at:  entity.Updated_at,
		Created_at:  entity.Created_at,
	}
}

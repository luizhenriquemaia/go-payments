package payments

import "time"

type AbstractPaymentFactory interface {
	Get_to_add() AddPaymentEntity
}

type PaymentFactory struct{}

func (factory *PaymentFactory) Get_to_add(
	description string,
	cost_center Cost_center,
	bar_code string,
) AddPaymentEntity {
	return AddPaymentEntity{
		Description: description,
		Cost_center: cost_center,
		Bar_code:    bar_code,
	}
}

func (factory *PaymentFactory) Get_from_db(
	id int64,
	description string,
	cost_center int,
	status int,
	bar_code string,
	updated_at time.Time,
	created_at time.Time,
) PaymentEntity {
	return PaymentEntity{
		Id:          id,
		Description: description,
		Cost_center: Cost_center(cost_center),
		Status:      Payment_status(status),
		Bar_code:    bar_code,
		Updated_at:  updated_at,
		Created_at:  created_at,
	}
}

func (factory *PaymentFactory) Get_to_resp(entity *PaymentEntity) *PaymentEntityResponse {
	return &PaymentEntityResponse{
		Id:          entity.Id,
		Description: entity.Description,
		Cost_center: entity.Cost_center.String(),
		Status:      entity.Status.String(),
		Bar_code:    entity.Bar_code,
		Updated_at:  entity.Updated_at,
		Created_at:  entity.Created_at,
	}
}

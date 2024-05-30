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
		Status:      status,
		Bar_code:    bar_code,
		Updated_at:  updated_at,
		Created_at:  created_at,
	}
}

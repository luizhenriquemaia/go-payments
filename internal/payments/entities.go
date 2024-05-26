package payments

import "time"

type PaymentEntity struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Cost_center int       `json:"cost_center"`
	Status      int       `json:"status"`
	Bar_code    string    `json:"bar_code"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}

type AddPaymentEntity struct {
	Description string `validate:"min=3,max=150"`
	Cost_center int    `validate:"min=0"`
	Bar_code    string `validate:"regexp=([0-9[]])"`
}

type addPaymentDb struct {
	description string
	cost_center int
	bar_code    string
	updated_at  time.Time
	created_at  time.Time
}

func (entity *AddPaymentEntity) Get_to_db() *addPaymentDb {
	now := time.Now().UTC()
	return &addPaymentDb{
		description: entity.Description,
		cost_center: entity.Cost_center,
		bar_code:    entity.Bar_code,
		updated_at:  now,
		created_at:  now,
	}
}

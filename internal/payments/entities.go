package payments

import "time"

type PaymentEntity struct {
	Id          int64          `json:"id"`
	Description string         `json:"description"`
	Cost_center Cost_center    `json:"cost_center"`
	Status      Payment_status `json:"status"`
	Bar_code    string         `json:"bar_code"`
	Updated_at  time.Time      `json:"updated_at"`
	Created_at  time.Time      `json:"created_at"`
}

type PaymentEntityResponse struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Cost_center string    `json:"cost_center"`
	Status      string    `json:"status"`
	Bar_code    string    `json:"bar_code"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}

type AddPaymentEntity struct {
	Description string      `binding:"required" validate:"min=3,max=150"`
	Cost_center Cost_center `binding:"required,enum" validate:"min=0"`
	Bar_code    string      `binding:"required" validate:"regexp=([0-9[]])"`
}

type addPaymentDb struct {
	description string
	cost_center Cost_center
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

//Adapter method to get payment entity to response
func (entity *PaymentEntity) Get_to_resp() *PaymentEntityResponse {
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

package payments

import (
	"strconv"
	"time"
)

type PaymentEntity struct {
	Id          int64          `json:"id"`
	Description string         `json:"description"`
	Cost_center CostCenter     `json:"cost_center"`
	Status      PaymentStatus  `json:"status"`
	Bar_code    string         `json:"bar_code"`
	Document    string         `json:"document"`
	Receipt     string         `json:"receipt"`
	Method      PaymentMethod  `json:"method"`
	Account     PaymentAccount `json:"account"`
	Paid_at     *time.Time     `json:"paid_at"`
	Updated_at  time.Time      `json:"updated_at"`
	Created_at  time.Time      `json:"created_at"`
}

type PaymentEntityResponse struct {
	Id          int64      `json:"id"`
	Description string     `json:"description"`
	Cost_center string     `json:"cost_center"`
	Status      string     `json:"status"`
	Bar_code    string     `json:"bar_code"`
	Document    string     `json:"document"`
	Receipt     string     `json:"receipt"`
	Method      string     `json:"method"`
	Account     string     `json:"account"`
	Paid_at     *time.Time `json:"paid_at"`
	Updated_at  time.Time  `json:"updated_at"`
	Created_at  time.Time  `json:"created_at"`
}

type PaymentReqQuery struct {
	Cost_center CostCenter
	Status      PaymentStatus
}

type AddPaymentEntity struct {
	Description string     `binding:"required,min_length=3,max_length=150"`
	Cost_center CostCenter `binding:"required,enum"`
	Bar_code    string     `binding:"required,only_digits,equal_length=47"`
}

type addPaymentDb struct {
	description string
	cost_center CostCenter
	bar_code    string
	document    string
	updated_at  time.Time
	created_at  time.Time
}

func (entity *AddPaymentEntity) getToDb() *addPaymentDb {
	now := time.Now().UTC()
	return &addPaymentDb{
		description: entity.Description,
		cost_center: entity.Cost_center,
		bar_code:    entity.Bar_code,
		document:    now.Format("200601021504") + strconv.Itoa(int(entity.Cost_center)),
		updated_at:  now,
		created_at:  now,
	}
}

// Adapter method to get payment entity to response
func (entity *PaymentEntity) getToResp() *PaymentEntityResponse {
	return &PaymentEntityResponse{
		Id:          entity.Id,
		Description: entity.Description,
		Cost_center: entity.Cost_center.String(),
		Status:      entity.Status.String(),
		Bar_code:    entity.Bar_code,
		Document:    entity.Document,
		Receipt:     entity.Receipt,
		Method:      entity.Method.String(),
		Account:     entity.Account.String(),
		Paid_at:     entity.Paid_at,
		Updated_at:  entity.Updated_at,
		Created_at:  entity.Created_at,
	}
}

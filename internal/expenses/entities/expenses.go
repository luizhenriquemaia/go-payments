package entities

import (
	"errors"
	"go-payments/internal/expenses/enums"
	"go-payments/utils"
	"strconv"
	"time"
)

type ExpenseEntity struct {
	Id          int64               `json:"id"`
	Description string              `json:"description"`
	Cost_center enums.CostCenter    `json:"cost_center"`
	Status      enums.ExpenseStatus `json:"status"`
	Bar_code    string              `json:"bar_code"`
	Document    string              `json:"document"`
	Due_date    time.Time           `json:"due_date"`
	Updated_at  time.Time           `json:"updated_at"`
	Created_at  time.Time           `json:"created_at"`
}

type ExpenseEntityResponse struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Cost_center string    `json:"cost_center"`
	Status      string    `json:"status"`
	Bar_code    string    `json:"bar_code"`
	Document    string    `json:"document"`
	Due_date    time.Time `json:"due_date"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}

type ExpenseReqQuery struct {
	Cost_center enums.CostCenter
	Status      enums.ExpenseStatus
}

type AddExpenseEntity struct {
	Description string           `binding:"required,min_length=3,max_length=150"`
	Cost_center enums.CostCenter `binding:"required,enum"`
	Bar_code    string           `binding:"required,only_digits,equal_length=47"`
	Due_date    string           `binding:"required"`
}

type AddExpenseDb struct {
	Description string
	Cost_center enums.CostCenter
	Bar_code    string
	Document    string
	Due_date    time.Time
	Updated_at  time.Time
	Created_at  time.Time
}

func (entity *AddExpenseEntity) GetToDb() (*AddExpenseDb, error) {
	now := time.Now().UTC()
	due_datetime, err := utils.ParseDateStrToTime(entity.Due_date)
	if err != nil {
		return nil, errors.New("data de vencimento inválida, por favor certifique-se que ela esteja no formato YYYY-MM-DD")
	}
	return &AddExpenseDb{
		Description: entity.Description,
		Cost_center: entity.Cost_center,
		Bar_code:    entity.Bar_code,
		Document:    now.Format("200601021504") + strconv.Itoa(int(entity.Cost_center)),
		Due_date:    *due_datetime,
		Updated_at:  now,
		Created_at:  now,
	}, nil
}

// Adapter method to get payment entity to response
func (entity *ExpenseEntity) GetToResp() *ExpenseEntityResponse {
	return &ExpenseEntityResponse{
		Id:          entity.Id,
		Description: entity.Description,
		Cost_center: entity.Cost_center.String(),
		Status:      entity.Status.String(),
		Bar_code:    entity.Bar_code,
		Document:    entity.Document,
		Due_date:    entity.Due_date,
		Updated_at:  entity.Updated_at,
		Created_at:  entity.Created_at,
	}
}

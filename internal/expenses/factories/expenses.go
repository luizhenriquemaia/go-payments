package factories

import (
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/enums"
	"log"
	"time"
)

type ExpenseFactory struct{}

func (factory *ExpenseFactory) GetFromDb(
	id int64,
	description string,
	cost_center int,
	status int,
	bar_code string,
	document string,
	due_date time.Time,
	updated_at time.Time,
	created_at time.Time,
) entities.ExpenseEntity {
	brazil_tz, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatal("Erro on load America/Sao_Paulo timezone for convert payment times")
	}
	status_enum, _ := enums.GetExpenseStatusByValue(status)
	cc_enum, _ := enums.GetCostCenterByValue(cost_center)
	return entities.ExpenseEntity{
		Id:          id,
		Description: description,
		Cost_center: cc_enum,
		Status:      status_enum,
		Bar_code:    bar_code,
		Document:    document,
		Due_date:    due_date.In(brazil_tz),
		Updated_at:  updated_at.In(brazil_tz),
		Created_at:  created_at.In(brazil_tz),
	}
}

func (factory *ExpenseFactory) GetToResp(entity *entities.ExpenseEntity) *entities.ExpenseEntityResponse {
	return &entities.ExpenseEntityResponse{
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

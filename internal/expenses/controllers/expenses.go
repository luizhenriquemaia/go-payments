package controllers

import (
	"go-payments/configs/database"
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/enums"
	"go-payments/internal/expenses/factories"
	"go-payments/internal/expenses/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func getExpensesFilterQuery(context *gin.Context) (*entities.ExpenseReqQuery, error) {
	q_status := context.DefaultQuery("status", "all")
	q_cost_center := context.DefaultQuery("cost_center", "all")
	cost_center, err := enums.GetCostCenterByName(q_cost_center)
	if err != nil && q_cost_center != "all" {
		return nil, err
	}
	status, err := enums.GetExpenseStatusByName(q_status)
	if err != nil && q_status != "all" {
		return nil, err
	}
	return &entities.ExpenseReqQuery{Cost_center: cost_center, Status: status}, nil
}

func GetExpensesController(context *gin.Context) (*[]entities.ExpenseEntityResponse, error) {
	query, err := getExpensesFilterQuery(context)
	if err != nil {
		return nil, err
	}
	repo := repositories.GetExpensesRepository(database.Get_db())
	e_payments, err := repo.FetchByStatusCC(query.Status, query.Cost_center)
	if err != nil {
		return nil, err
	}
	factory := factories.ExpenseFactory{}
	resp_entities := make([]entities.ExpenseEntityResponse, len(*e_payments))
	for i, entity := range *e_payments {
		resp_entities[i] = *factory.GetToResp(&entity)
	}
	return &resp_entities, nil
}

func AddExpensesController(context *gin.Context) (*entities.ExpenseEntityResponse, error) {
	var new_expense entities.AddExpenseEntity
	log.Printf("controller 13 %+v", new_expense)
	if err := context.ShouldBind(&new_expense); err != nil {
		return nil, err
	}
	repo := repositories.GetExpensesRepository(database.Get_db())
	expense_entity, err := repo.Add(&new_expense)
	if err != nil {
		return nil, err
	}
	response_entity := expense_entity.GetToResp()
	return response_entity, nil
}
package controllers

import (
	"go-payments/configs/database"
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/enums"
	"go-payments/internal/expenses/factories"
	"go-payments/internal/expenses/repositories"

	"github.com/gin-gonic/gin"
)

func getPaymentFilterQuery(context *gin.Context) (*entities.PaymentReqQuery, error) {
	q_status := context.DefaultQuery("status", "all")
	q_cost_center := context.DefaultQuery("cost_center", "all")
	cost_center, err := enums.GetCostCenterByName(q_cost_center)
	if err != nil && q_cost_center != "all" {
		return nil, err
	}
	status, err := enums.GetPaymentStatusByName(q_status)
	if err != nil && q_status != "all" {
		return nil, err
	}
	return &entities.PaymentReqQuery{Cost_center: cost_center, Status: status}, nil
}

func GetPaymentsController(context *gin.Context) (*[]entities.PaymentEntityResponse, error) {
	query, err := getPaymentFilterQuery(context)
	if err != nil {
		return nil, err
	}
	repo := repositories.PaymentsRepository(database.Get_db())
	e_payments, err := repo.FetchByStatusCC(query.Status, query.Cost_center)
	if err != nil {
		return nil, err
	}
	factory := factories.PaymentFactory{}
	resp_entities := make([]entities.PaymentEntityResponse, len(*e_payments))
	for i, entity := range *e_payments {
		resp_entities[i] = *factory.GetToResp(&entity)
	}
	return &resp_entities, nil
}

func AddPaymentController(context *gin.Context) (*entities.PaymentEntityResponse, error) {
	var new_payment entities.AddPaymentEntity
	if err := context.ShouldBind(&new_payment); err != nil {
		return nil, err
	}
	repo := repositories.PaymentsRepository(database.Get_db())
	payment_entity, err := repo.Add(&new_payment)
	if err != nil {
		return nil, err
	}
	response_entity := payment_entity.GetToResp()
	return response_entity, nil
}

// func payPaymentController(context *gin.Context) (*entities.PaymentEntityResponse, error) {

// }

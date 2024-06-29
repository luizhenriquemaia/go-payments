package payments

import (
	"go-payments/configs/database"

	"github.com/gin-gonic/gin"
)

func getPaymentFilterQuery(context *gin.Context) (*PaymentReqQuery, error) {
	q_status := context.DefaultQuery("status", "all")
	q_cost_center := context.DefaultQuery("cost_center", "all")
	cost_center, err := getCostCenterByName(q_cost_center)
	if err != nil && q_cost_center != "all" {
		return nil, err
	}
	status, err := getPaymentStatusByName(q_status)
	if err != nil && q_status != "all" {
		return nil, err
	}
	return &PaymentReqQuery{cost_center, status}, nil
}

func getPaymentsController(context *gin.Context) (*[]PaymentEntityResponse, error) {
	query, err := getPaymentFilterQuery(context)
	if err != nil {
		return nil, err
	}
	repo := PaymentsRepository(database.Get_db())
	entities, err := repo.fetchByStatusCC(query.Status, query.Cost_center)
	if err != nil {
		return nil, err
	}
	factory := PaymentFactory{}
	resp_entities := make([]PaymentEntityResponse, len(*entities))
	for i, entity := range *entities {
		resp_entities[i] = *factory.getToResp(&entity)
	}
	return &resp_entities, nil
}

func addPaymentController(context *gin.Context) (*PaymentEntityResponse, error) {
	var new_payment AddPaymentEntity
	if err := context.ShouldBind(&new_payment); err != nil {
		return nil, err
	}
	repo := PaymentsRepository(database.Get_db())
	payment_entity, err := repo.add(&new_payment)
	if err != nil {
		return nil, err
	}
	response_entity := payment_entity.getToResp()
	return response_entity, nil
}

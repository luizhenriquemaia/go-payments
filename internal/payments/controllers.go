package payments

import (
	"go-payments/configs/database"
	"log"

	"github.com/gin-gonic/gin"
)

func get_payment_filter_query(context *gin.Context) (*PaymentReqQuery, error) {
	q_status := context.DefaultQuery("status", "all")
	q_cost_center := context.DefaultQuery("cost_center", "all")
	cost_center, err := Get_cost_center_by_name(q_cost_center)
	if err != nil && q_cost_center != "all" {
		return nil, err
	}
	status, err := Get_payment_status_by_name(q_status)
	if err != nil && q_status != "all" {
		return nil, err
	}
	log.Printf("Status = %+v, Cost center = %+v", status, cost_center)
	log.Printf("entity = %+v", &PaymentReqQuery{cost_center, status})
	return &PaymentReqQuery{cost_center, status}, nil
}

func getPaymentsController(context *gin.Context) (*[]PaymentEntityResponse, error) {
	query, err := get_payment_filter_query(context)
	if err != nil {
		return nil, err
	}
	repo := PaymentsRepository(database.Get_db())
	entities, err := repo.Fetch_by_status_cost_center(query.Status, query.Cost_center)
	if err != nil {
		return nil, err
	}
	factory := PaymentFactory{}
	resp_entities := make([]PaymentEntityResponse, len(*entities))
	for i, entity := range *entities {
		resp_entities[i] = *factory.Get_to_resp(&entity)
	}
	return &resp_entities, nil
}

func addPaymentController(context *gin.Context) (*PaymentEntityResponse, error) {
	var new_payment AddPaymentEntity
	if err := context.ShouldBind(&new_payment); err != nil {
		return nil, err
	}
	repo := PaymentsRepository(database.Get_db())
	payment_entity, err := repo.Add(&new_payment)
	if err != nil {
		return nil, err
	}
	response_entity := payment_entity.Get_to_resp()
	return response_entity, nil
}

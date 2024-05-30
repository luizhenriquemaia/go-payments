package payments

import (
	"go-payments/configs/database"

	"github.com/gin-gonic/gin"
)

func getPaymentsController() string {
	return "test"
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

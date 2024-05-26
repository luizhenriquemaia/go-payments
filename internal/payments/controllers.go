package payments

import (
	"go-payments/configs/database"

	"github.com/gin-gonic/gin"
)

func getPaymentsController() string {
	return "test"
}

func AddPaymentController(context *gin.Context) (*PaymentEntity, error) {
	var new_payment AddPaymentEntity
	if err := context.BindJSON(&new_payment); err != nil {
		return nil, err
	}
	repo := PaymentsRepository(database.Get_db())
	payment_entity, err := repo.Add(&new_payment)
	if err != nil {
		return nil, err
	}
	return payment_entity, nil
}

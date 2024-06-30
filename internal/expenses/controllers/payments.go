package controllers

import (
	"go-payments/configs/database"
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func PayExpenseController(context *gin.Context) (*entities.PaymentRespEntity, error) {
	var new_payment entities.PayExpenseEntity
	new_payment.Expense_id = 52
	new_payment.Method = 1
	new_payment.Account = 1
	log.Printf("controller 13 %+v", new_payment)
	if err := context.ShouldBind(&new_payment); err != nil {
		log.Print("error binding")
		return nil, err
	}
	log.Print("controller 18")
	repo := repositories.GetPaymentsRepository(database.Get_db())
	payment_entity, err := repo.Add(&new_payment)
	if err != nil {
		return nil, err
	}
	result := payment_entity.GetToResp()
	return result, nil
}

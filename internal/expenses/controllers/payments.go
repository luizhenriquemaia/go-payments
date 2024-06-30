package controllers

import (
	"go-payments/configs/database"
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/repositories"

	"github.com/gin-gonic/gin"
)

func PayExpenseController(context *gin.Context) (*entities.PaymentRespEntity, error) {
	var new_payment entities.PayExpenseEntity
	if err := context.ShouldBind(&new_payment); err != nil {
		return nil, err
	}
	repo := repositories.GetPaymentsRepository(database.Get_db())
	payment_entity, err := repo.Add(&new_payment)
	if err != nil {
		return nil, err
	}
	result := payment_entity.GetToResp()
	return result, nil
}

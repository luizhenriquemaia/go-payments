package controllers

import (
	"errors"
	"go-payments/configs/database"
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/enums"
	"go-payments/internal/expenses/repositories"

	"github.com/gin-gonic/gin"
)

func PayExpenseController(context *gin.Context) (*entities.PaymentRespEntity, error) {
	var new_payment entities.PayExpenseEntity
	if err := context.ShouldBind(&new_payment); err != nil {
		return nil, err
	}

	repo_expenses := repositories.GetExpensesRepository(database.Get_db())
	expense, err := repo_expenses.FetchId(int64(new_payment.Expense_id))
	if err != nil {
		return nil, err
	}
	if !expense.Status.CheckCanPay() {
		return nil, errors.New("essa despesa já foi paga")
	}

	repo_payments := repositories.GetPaymentsRepository(database.Get_db())
	payment_entity, err := repo_payments.Add(&new_payment)
	if err != nil {
		return nil, err
	}

	err = repo_expenses.UpdateStatus(expense.Id, enums.StatusPaid)
	if err != nil {
		return nil, err
	}
	result := payment_entity.GetToResp()
	return result, nil
}

package services

import (
	"errors"
	"go-payments/configs/database"
	exp_enums "go-payments/internal/expenses/enums"
	exp_repo "go-payments/internal/expenses/repositories"
	"go-payments/internal/payments/entities"
	"go-payments/internal/payments/repositories"

	"github.com/gin-gonic/gin"
)

func PayExpenseService(context *gin.Context) (*entities.PaymentRespEntity, error) {
	var new_payment entities.PayExpenseEntity
	if err := context.ShouldBind(&new_payment); err != nil {
		return nil, err
	}

	repo_expenses := exp_repo.GetExpensesRepository(database.GetDb())
	expense, err := repo_expenses.FetchId(int64(new_payment.Expense_id))
	if err != nil {
		return nil, err
	}
	if !expense.Status.CheckCanPay() {
		return nil, errors.New("essa despesa já foi paga")
	}

	repo_payments := repositories.GetPaymentsRepository(database.GetDb())
	payment_entity, err := repo_payments.Add(&new_payment)
	if err != nil {
		return nil, err
	}

	err = repo_expenses.UpdateStatus(expense.Id, exp_enums.StatusPaid)
	if err != nil {
		return nil, err
	}
	result := payment_entity.GetToResp()
	return result, nil
}

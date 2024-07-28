package repositories

import (
	"database/sql"
	"errors"
	"go-payments/internal/payments/entities"
	"go-payments/internal/payments/factories"
	"log"
	"time"
)

type PaymentsRepository struct {
	db *sql.DB
}

func GetPaymentsRepository(db *sql.DB) *PaymentsRepository {
	return &PaymentsRepository{db}
}

func (repo *PaymentsRepository) setReceipt(id int) (*string, error) {
	new_receipt := ""
	err := repo.db.QueryRow(`
		UPDATE payment 
		SET receipt=CONCAT(receipt, CAST(id AS VARCHAR(20)))
		WHERE id=$1
		RETURNING receipt
	`, id).Scan(&new_receipt)
	if err != nil {
		log.Printf("update for set receipt error = %v", err)
		return nil, errors.New("não foi possível retornar um nome de comprovante para a nova despesa adicionada")
	}
	return &new_receipt, nil
}

func (repo *PaymentsRepository) Add(add_entity *entities.PayExpenseEntity) (*entities.PaymentEntity, error) {
	to_db := add_entity.GetToDb()
	new_id := -1
	err := repo.db.QueryRow(`
		INSERT INTO payment
			(expense_id, receipt, method, account, paid_at, updated_at, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		to_db.Expense_id,
		to_db.Receipt,
		to_db.Method,
		to_db.Account,
		to_db.Paid_at,
		to_db.Updated_at,
		to_db.Created_at,
	).Scan(&new_id)

	if err != nil {
		log.Printf("pay expense %v error = %v | values = %+v | now = %v", to_db.Expense_id, err, to_db, to_db.Updated_at.Format(time.RFC3339))
		return nil, errors.New("não foi possível pagar a despesa com os dados informados")
	}

	new_receipt, err := repo.setReceipt(new_id)
	if err != nil {
		return nil, err
	}

	factory := &factories.PaymentFactory{}
	retrieve_entity := factory.GetFromDb(
		int64(new_id),
		int64(to_db.Expense_id),
		*new_receipt,
		to_db.Method,
		to_db.Account,
		&to_db.Paid_at,
		to_db.Updated_at,
		to_db.Created_at,
	)
	return &retrieve_entity, nil
}

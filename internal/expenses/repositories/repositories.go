package repositories

import (
	"database/sql"
	"errors"
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/enums"
	"go-payments/internal/expenses/factories"
	"go-payments/internal/expenses/models"
	"log"
	"time"
)

var (
	ErrCreateFailed = errors.New("não foi possível criar pagamento com os dados informados")
)

type SqlRepository struct {
	db *sql.DB
}

func PaymentsRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{db}
}

func (repo *SqlRepository) GetFilteredRows(
	status enums.PaymentStatus,
	cost_center enums.CostCenter,
) (*sql.Rows, error) {
	if status == enums.StatusNotInformed && cost_center == enums.CcNotInformed {
		return repo.db.Query(`
			SELECT * FROM payment ORDER BY id DESC
		`)
	} else if status == enums.StatusNotInformed && cost_center != enums.CcNotInformed {
		return repo.db.Query(`
			SELECT * FROM payment WHERE cost_center=$1 ORDER BY id DESC
			`,
			cost_center,
		)
	} else if status != enums.StatusNotInformed && cost_center == enums.CcNotInformed {
		return repo.db.Query(`
		SELECT * FROM payment WHERE status=$1 ORDER BY id DESC
			`,
			status,
		)
	} else {
		return repo.db.Query(`
		SELECT * FROM payment WHERE cost_center=$1 AND status=$2 ORDER BY id DESC
			`,
			cost_center,
			status,
		)
	}
}

func (repo *SqlRepository) FetchByStatusCC(
	status enums.PaymentStatus,
	cost_center enums.CostCenter,
) (*[]entities.PaymentEntity, error) {
	rows, err := repo.GetFilteredRows(status, cost_center)
	if err != nil {
		log.Printf("get filtered payments error = %v", err)
		return nil, errors.New("não foi possível retornar os pagamentos filtrados")
	}
	defer rows.Close()

	var payments_db []models.PaymentModel
	for rows.Next() {
		var payment models.PaymentModel
		if err := rows.Scan(
			&payment.Id,
			&payment.Description,
			&payment.Cost_center,
			&payment.Status,
			&payment.Bar_code,
			&payment.Updated_at,
			&payment.Created_at,
			&payment.Document,
			&payment.Receipt,
			&payment.Paid_at,
			&payment.Method,
			&payment.Account,
		); err != nil {
			log.Printf("parsing filtered payment to entity in get payments error = %v", err)
			return nil, errors.New("não foi possível retornar os pagamentos filtrados")
		}
		payments_db = append(payments_db, payment)
	}

	factory := factories.PaymentFactory{}
	entities := make([]entities.PaymentEntity, len(payments_db))
	for i, payment := range payments_db {
		entities[i] = factory.GetFromDb(
			payment.Id,
			payment.Description,
			payment.Cost_center,
			payment.Status,
			payment.Bar_code,
			payment.Document,
			payment.Receipt,
			payment.Method,
			payment.Account,
			payment.Paid_at,
			payment.Updated_at,
			payment.Created_at,
		)
	}
	return &entities, nil
}

func (repo *SqlRepository) setDocument(id int) (*string, error) {
	new_document := ""
	err := repo.db.QueryRow(`
		UPDATE payment 
		SET document=CONCAT(document, CAST(id AS VARCHAR(20)))
		WHERE id=$1
		RETURNING document
	`, id).Scan(&new_document)
	if err != nil {
		log.Printf("update for set document error = %v", err)
		return nil, errors.New("não foi possível retornar os pagamentos salvos")
	}
	return &new_document, nil
}

func (repo *SqlRepository) Add(add_entity *entities.AddPaymentEntity) (*entities.PaymentEntity, error) {
	to_db := add_entity.GetToDb()
	new_id, new_status, new_method, new_account := -1, -1, -1, -1
	err := repo.db.QueryRow(`
		INSERT INTO payment(description, cost_center, bar_code, document, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, status, method, account
		`,
		to_db.Description,
		to_db.Cost_center,
		to_db.Bar_code,
		to_db.Document,
		to_db.Updated_at,
		to_db.Created_at,
	).Scan(&new_id, &new_status, &new_method, &new_account)

	if err != nil {
		log.Printf("add payments error = %v | values = %+v | now = %v", err, to_db, to_db.Updated_at.Format(time.RFC3339))
		return nil, errors.New("não foi possível adicionar um novo pagamento com os dados informados")
	}

	new_document, err := repo.setDocument(new_id)
	if err != nil {
		return nil, errors.New("não foi possível retornar um nome de documento para o novo pagamento adicionado")
	}

	factory := &factories.PaymentFactory{}
	retrieve_entity := factory.GetFromDb(
		int64(new_id),
		to_db.Description,
		int(to_db.Cost_center),
		new_status,
		to_db.Bar_code,
		*new_document,
		"",
		int(new_method),
		int(new_account),
		nil,
		to_db.Updated_at,
		to_db.Created_at,
	)
	return &retrieve_entity, nil
}

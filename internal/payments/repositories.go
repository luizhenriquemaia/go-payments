package payments

import (
	"database/sql"
	"errors"
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

func (repo *SqlRepository) Get() (*[]PaymentEntity, error) {
	rows, err := repo.db.Query(`
		SELECT * FROM payment ORDER BY id DESC
	`)
	if err != nil {
		log.Printf("get payments error = %v", err)
		return nil, errors.New("não foi possível retornar os pagamentos salvos")
	}
	defer rows.Close()

	var payments_db []PaymentModel
	for rows.Next() {
		var payment PaymentModel
		if err := rows.Scan(
			&payment.Id,
			&payment.Description,
			&payment.Cost_center,
			&payment.Status,
			&payment.Bar_code,
			&payment.Document,
			&payment.Receipt,
			&payment.Paid_at,
			&payment.Updated_at,
			&payment.Created_at,
		); err != nil {
			log.Printf("parsing payment to entity in get payments error = %v", err)
			return nil, errors.New("não foi possível retornar os pagamentos salvos")
		}
		payments_db = append(payments_db, payment)
	}

	factory := PaymentFactory{}
	entities := make([]PaymentEntity, len(payments_db))
	for i, payment := range payments_db {
		entities[i] = factory.Get_from_db(
			payment.Id,
			payment.Description,
			payment.Cost_center,
			payment.Status,
			payment.Bar_code,
			payment.Document,
			payment.Receipt,
			payment.Paid_at,
			payment.Updated_at,
			payment.Created_at,
		)
	}
	return &entities, nil
}

func (repo *SqlRepository) get_filtered_rows(
	status Payment_status,
	cost_center Cost_center,
) (*sql.Rows, error) {
	if status == StatusNotInformed && cost_center == CcNotInformed {
		return repo.db.Query(`
			SELECT * FROM payment ORDER BY id DESC
		`)
	} else if status == StatusNotInformed && cost_center != CcNotInformed {
		return repo.db.Query(`
			SELECT * FROM payment WHERE cost_center=$1 ORDER BY id DESC
			`,
			cost_center,
		)
	} else if status != StatusNotInformed && cost_center == CcNotInformed {
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

func (repo *SqlRepository) Fetch_by_status_cost_center(
	status Payment_status,
	cost_center Cost_center,
) (*[]PaymentEntity, error) {
	rows, err := repo.get_filtered_rows(status, cost_center)
	if err != nil {
		log.Printf("get filtered payments error = %v", err)
		return nil, errors.New("não foi possível retornar os pagamentos filtrados")
	}
	defer rows.Close()

	var payments_db []PaymentModel
	for rows.Next() {
		var payment PaymentModel
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
		); err != nil {
			log.Printf("parsing filtered payment to entity in get payments error = %v", err)
			return nil, errors.New("não foi possível retornar os pagamentos filtrados")
		}
		payments_db = append(payments_db, payment)
	}

	factory := PaymentFactory{}
	entities := make([]PaymentEntity, len(payments_db))
	for i, payment := range payments_db {
		entities[i] = factory.Get_from_db(
			payment.Id,
			payment.Description,
			payment.Cost_center,
			payment.Status,
			payment.Bar_code,
			payment.Document,
			payment.Receipt,
			payment.Paid_at,
			payment.Updated_at,
			payment.Created_at,
		)
	}
	return &entities, nil
}

func (repo *SqlRepository) Add(add_entity *AddPaymentEntity) (*PaymentEntity, error) {
	to_db := add_entity.Get_to_db()
	new_id, new_status := -1, -1
	err := repo.db.QueryRow(`
		INSERT INTO payment(description, cost_center, bar_code, document, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, status
		`,
		to_db.description,
		to_db.cost_center,
		to_db.bar_code,
		to_db.document,
		to_db.updated_at,
		to_db.created_at,
	).Scan(&new_id, &new_status)

	if err != nil {
		log.Printf("add payments error = %v | values = %+v | now = %v", err, to_db, to_db.updated_at.Format(time.RFC3339))
		return nil, errors.New("não foi possível adicionar um novo pagamento com os dados informados")
	}
	factory := &PaymentFactory{}
	retrieve_entity := factory.Get_from_db(
		int64(new_id),
		to_db.description,
		int(to_db.cost_center),
		new_status,
		to_db.bar_code,
		to_db.document,
		"",
		nil,
		to_db.updated_at,
		to_db.created_at,
	)
	return &retrieve_entity, nil
}

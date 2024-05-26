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

func (repo *SqlRepository) Add(add_entity *AddPaymentEntity) (*PaymentEntity, error) {
	to_db := add_entity.Get_to_db()
	new_id, new_status := -1, -1
	err := repo.db.QueryRow(`
		INSERT INTO payment(description, cost_center, bar_code, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, status
		`,
		to_db.description,
		to_db.cost_center,
		to_db.bar_code,
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
		to_db.cost_center,
		new_status,
		to_db.bar_code,
		to_db.updated_at,
		to_db.created_at,
	)
	return &retrieve_entity, nil
}

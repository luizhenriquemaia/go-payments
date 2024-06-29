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

func ExpensesRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{db}
}

func (repo *SqlRepository) GetFilteredRows(
	status enums.ExpenseStatus,
	cost_center enums.CostCenter,
) (*sql.Rows, error) {
	if status == enums.StatusNotInformed && cost_center == enums.CcNotInformed {
		return repo.db.Query(`
			SELECT * FROM expense ORDER BY id DESC
		`)
	} else if status == enums.StatusNotInformed && cost_center != enums.CcNotInformed {
		return repo.db.Query(`
			SELECT * FROM expense WHERE cost_center=$1 ORDER BY id DESC
			`,
			cost_center,
		)
	} else if status != enums.StatusNotInformed && cost_center == enums.CcNotInformed {
		return repo.db.Query(`
		SELECT * FROM expense WHERE status=$1 ORDER BY id DESC
			`,
			status,
		)
	} else {
		return repo.db.Query(`
		SELECT * FROM expense WHERE cost_center=$1 AND status=$2 ORDER BY id DESC
			`,
			cost_center,
			status,
		)
	}
}

func (repo *SqlRepository) FetchByStatusCC(
	status enums.ExpenseStatus,
	cost_center enums.CostCenter,
) (*[]entities.ExpenseEntity, error) {
	rows, err := repo.GetFilteredRows(status, cost_center)
	if err != nil {
		log.Printf("get filtered expenses error = %v", err)
		return nil, errors.New("não foi possível retornar as despesas filtradas")
	}
	defer rows.Close()

	var expenses_db []models.ExpenseModel
	for rows.Next() {
		var expense models.ExpenseModel
		if err := rows.Scan(
			&expense.Id,
			&expense.Description,
			&expense.Cost_center,
			&expense.Status,
			&expense.Bar_code,
			&expense.Updated_at,
			&expense.Created_at,
			&expense.Document,
		); err != nil {
			log.Printf("parsing filtered expense to entity in get expenses error = %v", err)
			return nil, errors.New("não foi possível retornar as despesas filtradas")
		}
		expenses_db = append(expenses_db, expense)
	}

	factory := factories.PaymentFactory{}
	entities := make([]entities.ExpenseEntity, len(expenses_db))
	for i, expense := range expenses_db {
		entities[i] = factory.GetFromDb(
			expense.Id,
			expense.Description,
			expense.Cost_center,
			expense.Status,
			expense.Bar_code,
			expense.Document,
			expense.Updated_at,
			expense.Created_at,
		)
	}
	return &entities, nil
}

func (repo *SqlRepository) setDocument(id int) (*string, error) {
	new_document := ""
	err := repo.db.QueryRow(`
		UPDATE expense 
		SET document=CONCAT(document, CAST(id AS VARCHAR(20)))
		WHERE id=$1
		RETURNING document
	`, id).Scan(&new_document)
	if err != nil {
		log.Printf("update for set document error = %v", err)
		return nil, errors.New("não foi possível retornar as despesas salvas")
	}
	return &new_document, nil
}

func (repo *SqlRepository) Add(add_entity *entities.AddExpenseEntity) (*entities.ExpenseEntity, error) {
	to_db := add_entity.GetToDb()
	new_id, new_status := -1, -1
	err := repo.db.QueryRow(`
		INSERT INTO expense(description, cost_center, bar_code, document, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, status
		`,
		to_db.Description,
		to_db.Cost_center,
		to_db.Bar_code,
		to_db.Document,
		to_db.Updated_at,
		to_db.Created_at,
	).Scan(&new_id, &new_status)

	if err != nil {
		log.Printf("add expenses error = %v | values = %+v | now = %v", err, to_db, to_db.Updated_at.Format(time.RFC3339))
		return nil, errors.New("não foi possível adicionar uma nova despesa com os dados informados")
	}

	new_document, err := repo.setDocument(new_id)
	if err != nil {
		return nil, errors.New("não foi possível retornar um nome de documento para a nova despesa adicionada")
	}

	factory := &factories.PaymentFactory{}
	retrieve_entity := factory.GetFromDb(
		int64(new_id),
		to_db.Description,
		int(to_db.Cost_center),
		new_status,
		to_db.Bar_code,
		*new_document,
		to_db.Updated_at,
		to_db.Created_at,
	)
	return &retrieve_entity, nil
}

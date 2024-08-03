package tests

import (
	"fmt"
	"go-payments/internal/expenses/entities"
	"go-payments/internal/expenses/enums"
	"go-payments/internal/expenses/services"
	"log"
	"time"
)

func AddTestExpenses(quantity int) []*entities.ExpenseEntityResponse {
	now := time.Now().UTC()
	str_dt := fmt.Sprintf("%v", now.Year()) + "-"
	month := int(now.Month())
	if month < 10 {
		str_dt += fmt.Sprintf("0%v", month) + "-"
	} else {
		str_dt += fmt.Sprintf("%v", month) + "-"
	}
	day := now.Day()
	if day < 10 {
		str_dt += fmt.Sprintf("0%v", day)
	} else {
		str_dt += fmt.Sprintf("%v", day)
	}

	result := make([]*entities.ExpenseEntityResponse, quantity)
	for i := 0; i < quantity; i++ {
		new_exp := &entities.AddExpenseEntity{
			Description: "Test expense",
			Cost_center: enums.CcEnergy,
			Bar_code:    "149182410294981210238129418284912",
			Due_date:    str_dt,
		}
		added_exp, err := services.AddExpensesService(new_exp)
		if err != nil {
			log.Fatalf("Couldn't added test expenses: %+v | %v", new_exp, err)
		}
		result = append(result, added_exp)
	}
	return result
}

package enums

import (
	"errors"
)

type CostCenter int64
type ExpenseStatus int64

const (
	CcNotInformed CostCenter = -1
	CcEnergy      CostCenter = 0
	CcWater       CostCenter = 1
	CcEducation   CostCenter = 2
	CcCondominium CostCenter = 3
)

var CostCenters = map[CostCenter]string{
	CcNotInformed: "not_informed",
	CcEnergy:      "energy",
	CcWater:       "water",
	CcEducation:   "education",
	CcCondominium: "condominium",
}

const (
	StatusNotInformed  ExpenseStatus = -1
	StatusPending      ExpenseStatus = 0
	StatusPaid         ExpenseStatus = 1
	StatusPaymentError ExpenseStatus = 2
	StatusOverdue      ExpenseStatus = 3
)

var ExpenseStatuses = map[ExpenseStatus]string{
	StatusNotInformed:  "not_informed",
	StatusPending:      "pending",
	StatusPaid:         "paid",
	StatusPaymentError: "payment_error",
	StatusOverdue:      "overdue",
}

func (cc CostCenter) IsValid() bool {
	if cc >= 0 && cc <= 3 {
		return true
	}
	return false
}

func (cc CostCenter) String() string {
	if cc < -1 || cc > 3 {
		return "unknown"
	}
	return CostCenters[cc]
}

func (status ExpenseStatus) String() string {
	if status < -1 || status > 3 {
		return "unknown"
	}
	return ExpenseStatuses[status]
}

func (status ExpenseStatus) CheckCanPay() bool {
	return ExpenseStatuses[status] != "paid"
}

func GetCostCenterByValue(cc int) (CostCenter, error) {
	if cc < -1 || cc > 3 {
		return CcNotInformed, errors.New("centro de custo inválido")
	}
	return CostCenter(cc), nil
}

func GetExpenseStatusByValue(status int) (ExpenseStatus, error) {
	if status < -1 || status > 3 {
		return StatusNotInformed, errors.New("centro de custo inválido")
	}
	return ExpenseStatus(status), nil
}

func GetCostCenterByName(cc_string string) (CostCenter, error) {
	switch cc_string {
	case "energy":
		return CcEnergy, nil
	case "water":
		return CcWater, nil
	case "education":
		return CcEducation, nil
	case "condominium":
		return CcCondominium, nil
	}
	return CcNotInformed, errors.New("centro de custo inválido")
}

func GetExpenseStatusByName(status string) (ExpenseStatus, error) {
	switch status {
	case "pending":
		return StatusPending, nil
	case "paid":
		return StatusPaid, nil
	case "payment_error":
		return StatusPaymentError, nil
	case "overdue":
		return StatusOverdue, nil
	}
	return StatusNotInformed, errors.New("status de pagamento inválido")
}

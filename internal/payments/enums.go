package payments

import (
	"errors"
)

type Cost_center int64
type Payment_status int64

const (
	CcNotInformed Cost_center = -1
	CcEnergy      Cost_center = 0
	CcWater       Cost_center = 1
	CcEducation   Cost_center = 2
	CcCondominium Cost_center = 3
)

var CostCenters = map[Cost_center]string{
	CcNotInformed: "not_informed",
	CcEnergy:      "energy",
	CcWater:       "water",
	CcEducation:   "education",
	CcCondominium: "condominium",
}

const (
	StatusNotInformed  Payment_status = -1
	StatusPending      Payment_status = 0
	StatusPaid         Payment_status = 1
	StatusPaymentError Payment_status = 2
	StatusOverdue      Payment_status = 3
)

var PaymentStatuses = map[Payment_status]string{
	StatusNotInformed:  "not_informed",
	StatusPending:      "pending",
	StatusPaid:         "paid",
	StatusPaymentError: "payment_error",
	StatusOverdue:      "overdue",
}

func (cc Cost_center) IsValid() bool {
	if cc >= 0 && cc <= 3 {
		return true
	}
	return false
}

func (cc Cost_center) String() string {
	if cc < -1 || cc > 3 {
		return "unknown"
	}
	return CostCenters[cc]
}

func Get_cost_center_by_value(cc int) (Cost_center, error) {
	if cc < -1 || cc > 3 {
		return CcNotInformed, errors.New("centro de custo inválido")
	}
	return Cost_center(cc), nil
}

func Get_cost_center_by_name(cc_string string) (Cost_center, error) {
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

func (status Payment_status) String() string {
	if status < -1 || status > 3 {
		return "unknown"
	}
	return PaymentStatuses[status]
}

func Get_payment_status_by_value(status int) (Payment_status, error) {
	if status < -1 || status > 3 {
		return StatusNotInformed, errors.New("centro de custo inválido")
	}
	return Payment_status(status), nil
}

func Get_payment_status_by_name(status string) (Payment_status, error) {
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

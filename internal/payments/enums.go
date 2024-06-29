package payments

import (
	"errors"
)

type CostCenter int64
type PaymentStatus int64

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
	StatusNotInformed  PaymentStatus = -1
	StatusPending      PaymentStatus = 0
	StatusPaid         PaymentStatus = 1
	StatusPaymentError PaymentStatus = 2
	StatusOverdue      PaymentStatus = 3
)

var PaymentStatuses = map[PaymentStatus]string{
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

func (cc CostCenter) string() string {
	if cc < -1 || cc > 3 {
		return "unknown"
	}
	return CostCenters[cc]
}

func getCostCenterByValue(cc int) (CostCenter, error) {
	if cc < -1 || cc > 3 {
		return CcNotInformed, errors.New("centro de custo inválido")
	}
	return CostCenter(cc), nil
}

func getCostCenterByName(cc_string string) (CostCenter, error) {
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

func (status PaymentStatus) string() string {
	if status < -1 || status > 3 {
		return "unknown"
	}
	return PaymentStatuses[status]
}

func getPaymentStatusByValue(status int) (PaymentStatus, error) {
	if status < -1 || status > 3 {
		return StatusNotInformed, errors.New("centro de custo inválido")
	}
	return PaymentStatus(status), nil
}

func getPaymentStatusByName(status string) (PaymentStatus, error) {
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

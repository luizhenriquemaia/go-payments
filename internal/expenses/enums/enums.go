package enums

import (
	"errors"
)

type CostCenter int64
type PaymentStatus int64
type PaymentMethod int64
type PaymentAccount int64

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

const (
	PaymentMethodNotInformed PaymentMethod = -1
	PaymentMethodCreditCard  PaymentMethod = 0
	PaymentMethodPix         PaymentMethod = 1
	PaymentMethodBoleto      PaymentMethod = 2
)

var PaymentMethods = map[PaymentMethod]string{
	PaymentMethodNotInformed: "not_informed",
	PaymentMethodCreditCard:  "credit_card",
	PaymentMethodPix:         "pix",
	PaymentMethodBoleto:      "boleto",
}

const (
	PaymentAccountNotInformed PaymentAccount = -1
	PaymentAccountNuBankPF    PaymentAccount = 0
	PaymentAccountInterPF     PaymentAccount = 1
	PaymentAccountInterPJ     PaymentAccount = 2
)

var PaymentAccounts = map[PaymentAccount]string{
	PaymentAccountNotInformed: "not_informed",
	PaymentAccountNuBankPF:    "nu_bank_pf",
	PaymentAccountInterPF:     "inter_pf",
	PaymentAccountInterPJ:     "inter_pj",
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

func (status PaymentStatus) String() string {
	if status < -1 || status > 3 {
		return "unknown"
	}
	return PaymentStatuses[status]
}

func (method PaymentMethod) String() string {
	if method < -1 || method > 2 {
		return "unknown"
	}
	return PaymentMethods[method]
}

func (account PaymentAccount) String() string {
	if account < -1 || account > 2 {
		return "unknown"
	}
	return PaymentAccounts[account]
}

func GetCostCenterByValue(cc int) (CostCenter, error) {
	if cc < -1 || cc > 3 {
		return CcNotInformed, errors.New("centro de custo inválido")
	}
	return CostCenter(cc), nil
}

func GetPaymentStatusByValue(status int) (PaymentStatus, error) {
	if status < -1 || status > 3 {
		return StatusNotInformed, errors.New("centro de custo inválido")
	}
	return PaymentStatus(status), nil
}

func GetPaymentMethodByValue(method int) (PaymentMethod, error) {
	if method < -1 || method > 2 {
		return PaymentMethodNotInformed, errors.New("método de pagamento inválido")
	}
	return PaymentMethod(method), nil
}

func GetPaymentAccountByValue(account int) (PaymentAccount, error) {
	if account < -1 || account > 2 {
		return PaymentAccountNotInformed, errors.New("conta de pagamento inválido")
	}
	return PaymentAccount(account), nil
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

func GetPaymentStatusByName(status string) (PaymentStatus, error) {
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

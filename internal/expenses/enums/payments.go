package enums

import (
	"errors"
)

type PaymentMethod int64
type PaymentAccount int64

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

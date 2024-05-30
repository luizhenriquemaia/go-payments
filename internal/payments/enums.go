package payments

type Cost_center int64
type Payment_status int64

const (
	Energy Cost_center = iota
	Water
	Education
	Condominium
)

const (
	Pending Payment_status = iota
	Paid
	Payment_error
	Overdue
)

func (cc Cost_center) IsValid() bool {
	switch cc {
	case Energy, Water, Education, Condominium:
		return true
	}
	return false
}

func (cc Cost_center) String() string {
	switch cc {
	case Energy:
		return "energy"
	case Water:
		return "water"
	case Education:
		return "education"
	case Condominium:
		return "condominium"
	}
	return "unknown"
}

func (status Payment_status) String() string {
	switch status {
	case Pending:
		return "pending"
	case Paid:
		return "paid"
	case Payment_error:
		return "payment_error"
	case Overdue:
		return "overdue"
	}
	return "unknown"
}

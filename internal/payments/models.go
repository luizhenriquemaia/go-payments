package payments

import "time"

type Cost_center int64

const (
	Energy Cost_center = iota
	Water
	Education
	Condominium
)

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

type PaymentModel struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Cost_center int       `json:"cost_center"`
	Status      int       `json:"status"`
	Bar_code    string    `json:"bar_code"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}

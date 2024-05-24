package payments

const (
	Energy = iota
	Water
	Education
)

type PaymentModel struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Cost_center int    `json:"cost_center"`
}

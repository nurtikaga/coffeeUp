package coffeeup

type Promo struct {
	ID       uint    `json:"id"`
	Code     string  `json:"code"`
	Discount float64 `json:"discount"`
}

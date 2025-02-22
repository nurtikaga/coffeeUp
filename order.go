package coffeeup

type Order struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"user_id"`
	Products []int   `json:"products"`
	Total    float64 `json:"total"`
}

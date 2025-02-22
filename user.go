package coffeeup

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Number   int    `json:"number"`
	Password string `json:"password_hash" binding:"required"`
	Role     int    `json:"role"`
}

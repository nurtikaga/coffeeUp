package coffeeup

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Password string `json:"password"`
}

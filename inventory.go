package coffeeup

type Inventory struct {
	RemainsMilk  int     `json:"remainsmilk"`
	RemainsWater int     `json:"remainswater"`
	RemainsSeed  int     `json:"remainsseed"`
	RemainsCup   int     `json:"remainscup"`
	RemainsSugar float32 `json:"remainssugar"`
}

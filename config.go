package coffeeup

type Config struct {
	Addr string
	Host string
}

func NewConfig() *Config {
	return &Config{
		Addr: ":8081",
	}
}

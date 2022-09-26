package config

var Conf Config

type Config struct {
	Secret secret
}

type secret struct {
	URL      string `env:"URL" envDefault:"https://example.com"`
	UserID   string `env:"UserID" envDefault:"1111111"`
	Email    string `env:"Email" envDefault:"example@email.com"`
	Password string `env:"Password" envDefault:"your_nice_password"`
}

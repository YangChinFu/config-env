package config

// AppConfig .
type AppConfig struct {
	Env   string
	Debug bool
}

// NewAppConfig return a new AppConfig struct
func NewAppConfig() AppConfig {
	return AppConfig{
		Env:   getString("APP_ENV", "debug"),
		Debug: getBool("APP_DEBUG", true),
	}
}

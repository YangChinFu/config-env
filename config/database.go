package config

// MySQLConfig .
type MySQLConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// NewMySQLConfig return a new MySQLConfig struct
func NewMySQLConfig() MySQLConfig {
	return MySQLConfig{
		Host:     getString("DB_HOST", "127.0.0.1"),
		Port:     getInt("DB_PORT", 3306),
		Database: getString("DB_DATABASE", "default"),
		Username: getString("DB_USERNAME", "default"),
		Password: getString("DB_PASSWORD", "secret"),
	}
}

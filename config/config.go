package config

import (
	"strconv"

	"github.com/YangChinFu/config-env/util"
)

// Config .
type Config struct {
	AppConfig
	MySQLConfig
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		AppConfig:   NewAppConfig(),
		MySQLConfig: NewMySQLConfig(),
	}
}

func getBool(key string, defaultVal ...bool) bool {
	valStr := util.GetEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal[0]
}

func getInt(key string, defaultVal ...int) int {
	valStr := util.GetEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}

	return defaultVal[0]
}

func getFloat(key string, defaultVal ...float64) float64 {
	valStr := util.GetEnv(key, "")
	if val, err := strconv.ParseFloat(valStr, 64); err == nil {
		return val
	}

	return defaultVal[0]
}

func getString(key string, defaultVal ...string) string {
	return util.GetEnv(key, defaultVal...)
}

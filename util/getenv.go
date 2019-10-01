package util

import "os"

// GetEnv can return the value of key in the env file,
// and the fallback will be returned if the key doesn't exist in env file.
func GetEnv(key string, fallback ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback[0]
}

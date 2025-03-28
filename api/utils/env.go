package utils

import "os"

// GetEnv returns the value of the environment variable or the fallback.
func GetEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

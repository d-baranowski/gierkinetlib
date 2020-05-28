package environment

import "os"

// GetEnv with a fallback value if not set
func GetStringVarDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

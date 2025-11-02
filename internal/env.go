package internal

import (
	"fmt"
	"os"
	"strconv"
)

// GetEnvInt retrieves an integer value from the environment variables
// based on the provided key. If the key is not found or the value
// cannot be converted to an integer, an error is returned.
func GetEnvInt(key string) (int, error) {
	value := os.Getenv(key)

	if value == "" {
		return 0, fmt.Errorf("missing environment setting for %s", key)
	}

	val, err := strconv.Atoi(value)

	if err != nil {
		return 0, err
	}

	return val, nil
}

// GetString retrieves a string value from the environment variables
// based on the provided key. If the key is not found, an error is returned.
func GetString(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("missing environment setting for %s", key)
	}

	return value, nil
}

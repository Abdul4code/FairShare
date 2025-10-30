package internal

import (
	"fmt"
	"os"
	"strconv"
)

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

func GetString(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("missing environment setting for %s", key)
	}

	return value, nil
}

package util

import (
	"os"

	"github.com/jinzhu/copier"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func DeepCopy[T any](source *T) (*T, error) {
	var target T

	if err := copier.Copy(&target, source); err != nil {
		return nil, err
	}

	return &target, nil
}

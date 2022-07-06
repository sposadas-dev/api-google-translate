package helper

import (
	"errors"
	"fmt"
	"os"
)

func GetValueFromEnvironmentVariable(variable string) (string, error) {
	if variable == "" {
		return "", errors.New("Variable not specified")
	}

	variableObtained := os.Getenv(variable)
	if variableObtained == "" {
		return "", errors.New(fmt.Sprintf("Please set %s environment variable", variable))
	}

	return variableObtained, nil
}

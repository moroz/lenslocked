package utils

import (
	"encoding/hex"
	"fmt"
	"os"
)

func RequireEnvVar(name string) string {
	value := os.Getenv(name)

	if value == "" {
		err := fmt.Sprintf("Environment variable %s is not set!", name)
		panic(err)
	}

	return value
}

func RequireHexEnvVar(name string) []byte {
	value := RequireEnvVar(name)
	decoded, err := hex.DecodeString(value)
	if err != nil {
		msg := fmt.Sprintf("Could not decode environment variable %s from hex", name)
		panic(msg)
	}
	return decoded
}

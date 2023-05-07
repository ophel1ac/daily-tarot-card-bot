package utils

import (
	"os"
	"strings"
)

func GetBotKey() string {
	key, err := os.ReadFile("secret.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(key), "\n")
}

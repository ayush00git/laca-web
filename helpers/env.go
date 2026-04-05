package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(envVar string) (string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("No .env file found")
	}

	Var := os.Getenv(envVar)
	return Var
}

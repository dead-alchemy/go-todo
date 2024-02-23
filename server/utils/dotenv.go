package util

import (
	"os"

	"github.com/joho/godotenv"
)

// function to get env variable from .env file or OS.
func GoDotEnv(key string) string {
	// setting channel for data so other routines can use it.
	env := make(chan string, 1)

	// load GO_ENV from os if exists or isn't production
	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load("../.env")

		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}

	// return the env variable.
	return <-env
}

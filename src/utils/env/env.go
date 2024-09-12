package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	envVarsFilename = "env/vars.env"
)

func SetEnvVars() {
	err := godotenv.Load(envVarsFilename)
	if err != nil {
		log.Fatal("error loading config")
	}
}

func GetEnvVar(name string) string {
	return os.Getenv(name)
}

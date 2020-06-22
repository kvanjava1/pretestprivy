package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dotEnv struct {
	key   string
	value string
}

func GetConfigValues() map[string]string {
	config := make(map[string]string)
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		config["ENV"] = os.Getenv("ENVIRONMENT")
		config["PORT"] = os.Getenv("DEVELOPMENT_PORT")
		config["APP_NAME"] = os.Getenv("DEVELOPMENT_APP_NAME")
		config["DEBUG_MODE"] = os.Getenv("DEVELOPMENT_DEBUG_MODE")

	} else if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		config["ENV"] = os.Getenv("ENVIRONMENT")
		config["PORT"] = os.Getenv("PRODUCTION_PORT")
		config["APP_NAME"] = os.Getenv("PRODUCTION_APP_NAME")
		config["DEBUG_MODE"] = os.Getenv("PRODUCTION_DEBUG_MODE")

	}

	return config
}

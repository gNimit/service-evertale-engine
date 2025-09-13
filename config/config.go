package config

import (
	"github.com/gNimit/service-evertale-engine/pkg/types"
	"github.com/gNimit/service-evertale-engine/pkg/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	EnvAppEnv string // Application environment (local, development, staging, production, testing)
	EnvPort   string // Port on which the application runs
)

// init initializes the configuration package.
func init() {
	if os.Getenv("APP_ENV") != "test" {
		rootPath, _ := utils.FindRootDir()
		if rootPath == "" {
			err := godotenv.Load()
			if err != nil {
				log.Fatalln("Error loading .env file:", err)
			}
		} else {
			err := godotenv.Load(rootPath + `/.env`)
			if err != nil {
				log.Fatalln("Error loading .env file:", err)
			}
		}
	}

	EnvAppEnv = os.Getenv("APP_ENV")
	EnvPort = os.Getenv("PORT")
}

// validate validates the configuration values.
func validate() {
	if EnvAppEnv == "" {
		// Default to local if not set
		EnvAppEnv = string(types.AppEnvLocal)
	}

	if EnvPort == "" {
		// Default to port 8080 if not set
		EnvPort = "8080"
	}
}

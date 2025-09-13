package types

// AppEnvironment defines the environment in which the application is running.
type AppEnvironment string

const (
	AppEnvLocal       AppEnvironment = "local"
	AppEnvDevelopment AppEnvironment = "dev"
	AppEnvStaging     AppEnvironment = "staging"
	AppEnvProduction  AppEnvironment = "prod"
	AppEnvTesting     AppEnvironment = "testing"
)

// LogLevel defines the severity level for logging.
type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
)

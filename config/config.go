package config

type Environment string

const (
	EnvLocal   Environment = "local"
	EnvDev     Environment = "dev"
	EnvTest    Environment = "test"
	EnvStaging Environment = "staging"
	EnvProd    Environment = "prod"
)

type Config struct {
	Env  Environment `json:"env"`
	Port string      `json:"port"`
}


func Load() (*Config, error)
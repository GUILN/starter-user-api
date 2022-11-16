package config

import (
	"os"
)

type AppConfig struct {
	paramaterGetter func(string) string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (ac *AppConfig) GetParameter(name string) string {
	return ac.paramaterGetter(name)
}

func (ac *AppConfig) WithEnvVarGetter() *AppConfig {
	ac.paramaterGetter = envVarConfigGetter
	return ac
}

func envVarConfigGetter(paramaterName string) string {
	return os.Getenv(paramaterName)
}

package rest

import (
	"github.com/bygui86/go-testing/rest-examples/http-server/logging"
	"github.com/bygui86/go-testing/rest-examples/http-server/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostDefault = "0.0.0.0"
	restPortDefault = 8080
)

func LoadConfig() *Config {
	logging.Log.Debug("Load REST configurations")
	return &Config{
		RestHost: utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort: utils.GetIntEnv(restPortEnvVar, restPortDefault),
	}
}

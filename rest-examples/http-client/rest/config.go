package rest

import (
	"github.com/bygui86/go-testing/rest-examples/http-client/logging"
	"github.com/bygui86/go-testing/rest-examples/http-client/utils"
)

const (
	restServerHostEnvVar = "REST_SERVER_HOST"
	restServerPortEnvVar = "REST_SERVER_PORT"
	restHostEnvVar       = "REST_HOST"
	restPortEnvVar       = "REST_PORT"

	restServerHostDefault = "localhost"
	restServerPortDefault = 8080
	restHostDefault       = "0.0.0.0"
	restPortDefault       = 8080
)

func LoadConfig() *Config {
	logging.Log.Debug("Load REST configurations")

	return &Config{
		RestServerHost: utils.GetStringEnv(restServerHostEnvVar, restServerHostDefault),
		RestServerPort: utils.GetIntEnv(restServerPortEnvVar, restServerPortDefault),
		RestHost:       utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort:       utils.GetIntEnv(restPortEnvVar, restPortDefault),
	}
}

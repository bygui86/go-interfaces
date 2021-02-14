package logging

import (
	"fmt"

	"github.com/bygui86/go-testing/logging-example/utils"
)

const (
	encodingEnvVar   = "LOG_ENCODING"        // available values: console (default), json
	levelEnvVar      = "LOG_LEVEL"           //  available values: trace, debug, info (default), warn, error, fatal
	outPathEnvVar    = "LOG_OUTPUT_PATH"     //  available values: stdout
	errOutPathEnvVar = "LOG_ERR_OUTPUT_PATH" //  available values: stderr

	encodingDefault   = "console"
	levelDefault      = "info"
	outPathDefault    = "stdout"
	errOutPathDefault = "stderr"
)

func LoadConfig() *Config {
	fmt.Println("[INFO] Load Logging configurations")

	return &Config{
		encoding:      utils.GetStringEnv(encodingEnvVar, encodingDefault),
		level:         utils.GetStringEnv(levelEnvVar, levelDefault),
		outputPath:    utils.GetStringEnv(outPathEnvVar, outPathDefault),
		errOutputPath: utils.GetStringEnv(errOutPathEnvVar, errOutPathDefault),
	}
}

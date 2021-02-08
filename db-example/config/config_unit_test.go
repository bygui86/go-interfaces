package config_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/db-example/config"
	"github.com/bygui86/go-testing/db-example/logging"
)

const (
	monitorKey   = "ENABLE_MONITORING"
	monitorValue = false

	traceKey   = "ENABLE_TRACING"
	traceValue = false

	techKey   = "TRACING_TECH"
	techValue = "zipkin"
)

func TestLoadConfig(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.Nil(t, logErr)

	monitorErr := os.Setenv(monitorKey, strconv.FormatBool(monitorValue))
	require.Nil(t, monitorErr)
	traceErr := os.Setenv(traceKey, strconv.FormatBool(traceValue))
	require.Nil(t, traceErr)
	techErr := os.Setenv(techKey, techValue)
	require.Nil(t, techErr)

	cfg := config.LoadConfig()

	assert.Equal(t, monitorValue, cfg.GetEnableMonitoring())
	assert.Equal(t, traceValue, cfg.GetEnableTracing())
	assert.Equal(t, techValue, cfg.GetTracingTech())

	err := os.Unsetenv(monitorKey)
	require.Nil(t, err)
	err = os.Unsetenv(traceKey)
	require.Nil(t, err)
	err = os.Unsetenv(techKey)
	require.Nil(t, err)
}

func TestLoadConfig_Defaults(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.Nil(t, logErr)

	cfg := config.LoadConfig()

	assert.Equal(t, true, cfg.GetEnableMonitoring())
	assert.Equal(t, true, cfg.GetEnableTracing())
	assert.Equal(t, config.TracingTechJaeger, cfg.GetTracingTech())
}

func TestLoadConfig_TracingTechNotSupported(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.Nil(t, logErr)

	techErr := os.Setenv(techKey, "not-supported")
	require.Nil(t, techErr)

	cfg := config.LoadConfig()

	assert.Equal(t, true, cfg.GetEnableMonitoring())
	assert.Equal(t, true, cfg.GetEnableTracing())
	assert.Equal(t, config.TracingTechJaeger, cfg.GetTracingTech())

	err := os.Unsetenv(techKey)
	require.Nil(t, err)
}

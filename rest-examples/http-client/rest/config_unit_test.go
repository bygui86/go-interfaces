// +build unit !integration

package rest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/rest-examples/http-client/logging"
	"github.com/bygui86/go-testing/rest-examples/http-client/rest"
)

const (
	serverHostKey   = "REST_SERVER_HOST"
	serverHostValue = "server"

	serverPortKey   = "REST_SERVER_PORT"
	serverPortValue = 9876

	hostKey   = "REST_HOST"
	hostValue = "remote"

	portKey   = "REST_PORT"
	portValue = 9876
)

func TestLoadConfig(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	serverHostErr := os.Setenv(serverHostKey, serverHostValue)
	require.NoError(t, serverHostErr)
	serverPortErr := os.Setenv(serverPortKey, strconv.Itoa(serverPortValue))
	require.NoError(t, serverPortErr)
	hostErr := os.Setenv(hostKey, hostValue)
	require.NoError(t, hostErr)
	portErr := os.Setenv(portKey, strconv.Itoa(portValue))
	require.NoError(t, portErr)

	cfg := rest.LoadConfig()
	assert.Equal(t, serverHostValue, cfg.RestServerHost)
	assert.Equal(t, serverPortValue, cfg.RestServerPort)
	assert.Equal(t, hostValue, cfg.RestHost)
	assert.Equal(t, portValue, cfg.RestPort)

	err := os.Unsetenv(hostKey)
	require.NoError(t, err)
	err = os.Unsetenv(portKey)
	require.NoError(t, err)
}

func TestLoadConfig_Defaults(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := rest.LoadConfig()
	assert.Equal(t, rest.RestServerHostDefault(), cfg.RestServerHost)
	assert.Equal(t, rest.RestServerPortDefault(), cfg.RestServerPort)
	assert.Equal(t, rest.RestHostDefault(), cfg.RestHost)
	assert.Equal(t, rest.RestPortDefault(), cfg.RestPort)
}

// +build unit !integration

package rest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/rest-examples/http-server/logging"
	"github.com/bygui86/go-testing/rest-examples/http-server/rest"
)

const (
	hostKey   = "REST_HOST"
	hostValue = "remote"

	portKey   = "REST_PORT"
	portValue = 9876
)

func TestLoadConfig(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	hostErr := os.Setenv(hostKey, hostValue)
	require.NoError(t, hostErr)
	portErr := os.Setenv(portKey, strconv.Itoa(portValue))
	require.NoError(t, portErr)

	cfg := rest.LoadConfig()

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

	assert.Equal(t, rest.RestHostDefault, cfg.RestHost)
	assert.Equal(t, rest.RestPortDefault, cfg.RestPort)
}

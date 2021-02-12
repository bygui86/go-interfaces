// +build unit !integration

package database

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/db-example/logging"
)

const (
	hostKey   = "DB_HOST"
	hostValue = "postgresql.remote.test"

	portKey   = "DB_PORT"
	portValue = 5433

	userKey   = "DB_USERNAME"
	userValue = "john"

	pwKey   = "DB_PASSWORD"
	pwValue = "myPass"

	nameKey   = "DB_NAME"
	nameValue = "test"

	sslKey   = "DB_SSL_MODE"
	sslValue = "enable"
)

func TestLoadConfig(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	hostErr := os.Setenv(hostKey, hostValue)
	require.NoError(t, hostErr)
	portErr := os.Setenv(portKey, strconv.Itoa(portValue))
	require.NoError(t, portErr)
	userErr := os.Setenv(userKey, userValue)
	require.NoError(t, userErr)
	pwErr := os.Setenv(pwKey, pwValue)
	require.NoError(t, pwErr)
	nameErr := os.Setenv(nameKey, nameValue)
	require.NoError(t, nameErr)
	sslErr := os.Setenv(sslKey, sslValue)
	require.NoError(t, sslErr)

	cfg := loadConfig()

	assert.Equal(t, hostValue, cfg.dbHost)
	assert.Equal(t, portValue, cfg.dbPort)
	assert.Equal(t, userValue, cfg.dbUsername)
	assert.Equal(t, pwValue, cfg.dbPassword)
	assert.Equal(t, nameValue, cfg.dbName)
	assert.Equal(t, sslValue, cfg.dbSslMode)

	err := os.Unsetenv(hostKey)
	require.NoError(t, err)
	err = os.Unsetenv(portKey)
	require.NoError(t, err)
	err = os.Unsetenv(userKey)
	require.NoError(t, err)
	err = os.Unsetenv(pwKey)
	require.NoError(t, err)
	err = os.Unsetenv(nameKey)
	require.NoError(t, err)
	err = os.Unsetenv(sslKey)
	require.NoError(t, err)
}

func TestLoadConfig_Defaults(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := loadConfig()

	assert.Equal(t, dbHostDefault, cfg.dbHost)
	assert.Equal(t, dbPortDefault, cfg.dbPort)
	assert.Equal(t, dbUsernameDefault, cfg.dbUsername)
	assert.Equal(t, dbPasswordDefault, cfg.dbPassword)
	assert.Equal(t, dbNameDefault, cfg.dbName)
	assert.Equal(t, dbSslModeDefault, cfg.dbSslMode)
}

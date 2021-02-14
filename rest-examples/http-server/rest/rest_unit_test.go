// +build unit !integration

package rest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/rest-examples/http-server/database"
	"github.com/bygui86/go-testing/rest-examples/http-server/logging"
	"github.com/bygui86/go-testing/rest-examples/http-server/rest"
)

func TestNew_Unit(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := &rest.Config{
		RestHost: "localhost",
		RestPort: 8080,
	}
	db := database.New() // TODO replace with a mock
	server := rest.New(cfg, db)
	assert.NotNil(t, server)
}

func TestStart_Unit(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := &rest.Config{
		RestHost: "localhost",
		RestPort: 8080,
	}
	db := database.New() // TODO replace with a mock
	server := rest.New(cfg, db)
	require.NotNil(t, server)

	startErr := server.Start()
	assert.NoError(t, startErr)
	assert.True(t, server.Running())

	server.Shutdown(1)
}

func TestShutdown_Unit(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := &rest.Config{
		RestHost: "localhost",
		RestPort: 8080,
	}

	db := database.New() // TODO replace with a mock

	server := rest.New(cfg, db)
	require.NotNil(t, server)

	startErr := server.Start()
	require.NoError(t, startErr)
	require.True(t, server.Running())

	server.Shutdown(1)
	assert.False(t, server.Running())
}

func TestShutdown_Unit_NoStart(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := &rest.Config{
		RestHost: "localhost",
		RestPort: 8080,
	}

	db := database.New() // TODO replace with a mock

	server := rest.New(cfg, db)
	require.NotNil(t, server)

	server.Shutdown(1)
	assert.False(t, server.Running())
}

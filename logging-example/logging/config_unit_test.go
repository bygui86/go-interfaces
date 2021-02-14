// +build unit !integration

package logging_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-testing/logging-example/logging"
)

const (
	encodingKey   = "LOG_ENCODING"
	encodingValue = "json"

	levelKey   = "LOG_LEVEL"
	levelValue = "debug"

	outPathKey   = "LOG_OUTPUT_PATH"
	outPathValue = "outbuffer"

	errOutPathKey   = "LOG_ERR_OUTPUT_PATH"
	errOutPathValue = "errbuffer"
)

func TestLoadConfig(t *testing.T) {
	encodingErr := os.Setenv(encodingKey, encodingValue)
	require.NoError(t, encodingErr)
	levelErr := os.Setenv(levelKey, levelValue)
	require.NoError(t, levelErr)
	outErr := os.Setenv(outPathKey, outPathValue)
	require.NoError(t, outErr)
	erroutErr := os.Setenv(errOutPathKey, errOutPathValue)
	require.NoError(t, erroutErr)

	cfg := logging.LoadConfig()
	assert.Equal(t, encodingValue, cfg.Encoding())
	assert.Equal(t, levelValue, cfg.Level())
	assert.Equal(t, outPathValue, cfg.OutputPath())
	assert.Equal(t, errOutPathValue, cfg.ErrOutputPath())

	err := os.Unsetenv(encodingKey)
	require.NoError(t, err)
	err = os.Unsetenv(levelKey)
	require.NoError(t, err)
	err = os.Unsetenv(outPathKey)
	require.NoError(t, err)
	err = os.Unsetenv(errOutPathKey)
	require.NoError(t, err)
}

func TestLoadConfig_Defaults(t *testing.T) {
	cfg := logging.LoadConfig()
	assert.Equal(t, logging.EncodingDefault(), cfg.Encoding())
	assert.Equal(t, logging.LevelDefault(), cfg.Level())
	assert.Equal(t, logging.OutputPathDefault(), cfg.OutputPath())
	assert.Equal(t, logging.ErrOutputPathDefault(), cfg.ErrOutputPath())
}

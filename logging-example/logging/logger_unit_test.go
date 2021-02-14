// +build unit !integration

package logging_test

import (
	"bytes"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	"github.com/stretchr/testify/assert"

	"github.com/bygui86/go-testing/logging-example/logging"
)

func TestInitGlobalLoggerFromEnvVar(t *testing.T) {
	cfg := logging.LoadConfig()
	err := logging.InitGlobalLoggerFromEnvVar(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, &logging.ZapConfig)
	assert.NotNil(t, &logging.Log)
	assert.NotNil(t, &logging.SugaredLog)
	assert.Equal(t, cfg.Encoding(), logging.ZapConfig.Encoding)
	assert.Equal(t, cfg.Level(), logging.ZapConfig.Level.String())
	assert.Equal(t, cfg.OutputPath(), logging.ZapConfig.OutputPaths[0])
	assert.Equal(t, cfg.ErrOutputPath(), logging.ZapConfig.ErrorOutputPaths[0])

	// INFO: this line is part of the test, it should not generate panic
	logging.Log.Info("no panic!")
}

func TestInitGlobalLogger_Stdout(t *testing.T) {
	cfg := logging.LoadConfig()
	zapCfg, cfgErr := logging.BuildLoggerConfigFromEnvVar(cfg)
	require.NoError(t, cfgErr)

	err := logging.InitGlobalLogger(zapCfg)
	assert.NoError(t, err)
	assert.NotNil(t, &logging.ZapConfig)
	assert.NotNil(t, &logging.Log)
	assert.NotNil(t, &logging.SugaredLog)
	assert.Equal(t, cfg.Encoding(), logging.ZapConfig.Encoding)
	assert.Equal(t, cfg.Level(), logging.ZapConfig.Level.String())
	assert.Equal(t, cfg.OutputPath(), logging.ZapConfig.OutputPaths[0])
	assert.Equal(t, cfg.ErrOutputPath(), logging.ZapConfig.ErrorOutputPaths[0])

	// INFO: this line is part of the test, it should not generate panic
	logging.Log.Info("no panic!")
}

func TestBuildLoggerConfig(t *testing.T) {
	encoding := "console"
	level := "info"
	outPath := "stdout"
	errPath := "stderr"
	cfg, err := logging.BuildLoggerConfig(encoding, level, outPath, errPath)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, encoding, cfg.Encoding)
	assert.Equal(t, level, cfg.Level.String())
	assert.Equal(t, outPath, cfg.OutputPaths[0])
	assert.Equal(t, errPath, cfg.ErrorOutputPaths[0])
}

// from https://stackoverflow.com/questions/52734529/testing-zap-logging-for-a-logger-built-from-a-custom-config/52737940

// MemorySink implements zap.Sink by writing all messages to a buffer.
type MemorySink struct {
	*bytes.Buffer
}

// Implement Close and Sync as no-ops to satisfy the interface.
// The Write method is provided by the embedded buffer.
func (s *MemorySink) Close() error { return nil }
func (s *MemorySink) Sync() error  { return nil }

func TestInitGlobalLogger_MemorySink(t *testing.T) {
	encoding := "console"
	level := "info"
	outSinkPath := "memory"
	outPath := "memory://"
	errOutPath := "stderr"

	// Create a sink instance
	sink := &MemorySink{
		new(bytes.Buffer),
	}

	// Register new sink with zap for the "memory" protocol
	recErr := zap.RegisterSink(
		outSinkPath,
		func(*url.URL) (zap.Sink, error) {
			return sink, nil
		},
	)
	require.NoError(t, recErr)

	// Redirect all messages to the MemorySink (memory protocol)
	cfg, cfgErr := logging.BuildLoggerConfig(encoding, level, outPath, errOutPath)
	require.NoError(t, cfgErr)

	initErr := logging.InitGlobalLogger(cfg)
	require.NoError(t, initErr)
	require.NotNil(t, &logging.ZapConfig)
	require.NotNil(t, &logging.Log)
	require.NotNil(t, &logging.SugaredLog)
	require.Equal(t, encoding, logging.ZapConfig.Encoding)
	require.Equal(t, level, logging.ZapConfig.Level.String())
	require.Equal(t, outPath, logging.ZapConfig.OutputPaths[0])
	require.Equal(t, errOutPath, logging.ZapConfig.ErrorOutputPaths[0])

	logging.Log.Info(
		"failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	output := sink.String()
	// t.Logf("MemorySink output:   *** %s ***", output)
	assert.True(t, strings.Contains(output, `"url": "http://example.com"`))
}

// from https://gianarb.it/blog/golang-mockmania-zap-logger

func TestZapGeneric(t *testing.T) {
	logger := zaptest.NewLogger(t)
	assert.NotNil(t, logger)

	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}

func TestZapHook(t *testing.T) {
	logger := zaptest.NewLogger(
		t,
		zaptest.WrapOptions(
			zap.Hooks(
				func(e zapcore.Entry) error {
					if e.Level == zap.FatalLevel {
						t.Fatal("Error should never happen!")
					}
					return nil
				},
			),
		),
	)
	assert.NotNil(t, logger)

	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}

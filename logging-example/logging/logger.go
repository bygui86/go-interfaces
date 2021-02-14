package logging

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ZapConfig  *zap.Config
	Log        *zap.Logger
	SugaredLog *zap.SugaredLogger
)

func InitGlobalLoggerFromEnvVar(cfg *Config) error {
	fmt.Println("[INFO] Initialize global logger")

	var cfgErr error
	ZapConfig, cfgErr = BuildLoggerConfigFromEnvVar(cfg)
	if cfgErr != nil {
		return cfgErr
	}

	var logErr error
	Log, logErr = ZapConfig.Build()
	if logErr != nil {
		return logErr
	}

	SugaredLog = Log.Sugar()
	return nil
}

func BuildLoggerConfigFromEnvVar(cfg *Config) (*zap.Config, error) {
	zapLevel, err := getZapLevel(cfg.level)
	if err != nil {
		return nil, err
	}

	return &zap.Config{
		Encoding:         cfg.encoding,
		Level:            zap.NewAtomicLevelAt(zapLevel),
		OutputPaths:      []string{cfg.outputPath},
		ErrorOutputPaths: []string{cfg.errOutputPath},
		EncoderConfig:    buildEncoderConfig(zapLevel),
	}, nil
}

func InitGlobalLogger(cfg *zap.Config) error {
	fmt.Println("[INFO] Initialize global logger")

	ZapConfig = cfg

	var logErr error
	Log, logErr = ZapConfig.Build()
	if logErr != nil {
		return logErr
	}

	SugaredLog = Log.Sugar()
	return nil
}

func BuildLoggerConfig(encoding, level, outputPath, errOutputPath string) (*zap.Config, error) {
	zapLevel, err := getZapLevel(level)
	if err != nil {
		return nil, err
	}

	return &zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(zapLevel),
		OutputPaths:      []string{outputPath},
		ErrorOutputPaths: []string{errOutputPath},
		EncoderConfig:    buildEncoderConfig(zapLevel),
	}, nil
}

func getZapLevel(levelString string) (zapcore.Level, error) {
	level := zapcore.InfoLevel
	err := level.Set(levelString)
	if err != nil {
		return zapcore.InfoLevel, err
	}
	return level, nil
}

func buildEncoderConfig(level zapcore.Level) zapcore.EncoderConfig {
	if level == zapcore.DebugLevel {
		return zapcore.EncoderConfig{
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			MessageKey:   "message",
		}
	} else {
		return zapcore.EncoderConfig{
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			MessageKey:  "message",
		}
	}
}

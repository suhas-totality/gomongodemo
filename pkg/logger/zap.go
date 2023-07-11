package logger

import (
	"log"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Log *zap.SugaredLogger
}

func NewLogger(level int) *Logger {
	logLevel := parseLevel(level)
	path := getCWD()
	createDir(path)
	writerSync := getLogWriter(path)
	encoder := getEncoder()
	return &Logger{
		Log: zap.New(
			zapcore.NewCore(
				encoder,
				writerSync,
				logLevel,
			),
			zap.AddCaller(),
		).Sugar(),
	}
}

func parseLevel(level int) zapcore.Level {
	logLevel := []zapcore.Level{
		zapcore.DebugLevel,
		zapcore.InfoLevel,
		zapcore.WarnLevel,
		zapcore.ErrorLevel,
	}
	return logLevel[level]
}

func getCWD() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("unable to get cwd; exiting program")
	}
	return path
}

func createDir(path string) {
	if _, err := os.Stat(filepath.Join(path, LOG_DIRECTORY)); os.IsNotExist(err) {
		log.Println("creating logs directory", LOG_DIRECTORY)
		errCreatingDir := os.Mkdir(LOG_DIRECTORY, os.ModePerm)
		if errCreatingDir != nil {
			log.Fatalln("unable to create directory", LOG_DIRECTORY)
		}
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(path, LOG_DIRECTORY, LOG_FILE),
		MaxSize:    LOG_FILE_MAX_SIZE,
		MaxBackups: LOG_FILE_MAX_NUMBER,
		MaxAge:     LOG_FILE_MAX_AGE,
		Compress:   LOG_FILE_COMPRESSION_ENABLE,
	}
	return zapcore.AddSync(lumberJackLogger)
}

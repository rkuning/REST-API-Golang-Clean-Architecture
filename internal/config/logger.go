package config

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(level int) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.Level(level))

	return logger
}

func ConfigureLogFileOutput(logger *logrus.Logger) {
	stdoutWriter := os.Stdout
	fileWriter, err := os.OpenFile("logger.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.WithError(err).Error("failed to open log file")
		return
	}

	multiWriter := io.MultiWriter(stdoutWriter, fileWriter)
	logger.SetOutput(multiWriter)
}

package config

import "github.com/sirupsen/logrus"

func NewLogger(level int) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.Level(level))

	return logger
}
